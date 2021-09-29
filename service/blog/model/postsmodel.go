package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	postsFieldNames          = builderx.RawFieldNames(&Posts{})
	postsRows                = strings.Join(postsFieldNames, ",")
	postsRowsExpectAutoSet   = strings.Join(stringx.Remove(postsFieldNames, "`create_time`", "`update_time`"), ",")
	postsRowsWithPlaceHolder = strings.Join(stringx.Remove(postsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheBlogPostsIdPrefix = "cache:blog:posts:id:"
)

type (
	PostsModel interface {
		Insert(data Posts) (sql.Result, error)
		FindOne(id int64) (*Posts, error)
		Update(data Posts) error
		Delete(id int64) error
		Count() (int64, error)
		CountWtihTag(tag string) (int64, error)
		Posts(offset, size int64) ([]Posts, error)
		PostsWithTag(offset, size int64, tag string) ([]Posts, error)
		Title(id int64) (string, error)
	}

	defaultPostsModel struct {
		sqlc.CachedConn
		table string
	}

	Posts struct {
		Id          int64          `db:"id"`
		CreatedAt   time.Time      `db:"created_at"`
		UpdatedAt   time.Time      `db:"updated_at"`
		DeletedAt   sql.NullTime   `db:"deleted_at"`
		Title       string         `db:"title"`
		Description sql.NullString `db:"description"`
		Pre         int64          `db:"pre"`
		Next        int64          `db:"next"`
		Url         string         `db:"url"`
		Path        string         `db:"path"`
		Tags        sql.NullString `db:"tags"`
	}
)

func NewPostsModel(conn sqlx.SqlConn, c cache.CacheConf) PostsModel {
	return &defaultPostsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`posts`",
	}
}

func (m *defaultPostsModel) Title(id int64) (string, error) {
	var ret string

	query := fmt.Sprintf("select title from %s where id = %d", m.table, id)
	err := m.QueryRowNoCache(&ret, query)

	if err != nil {
		return "", err
	}

	return ret, nil
}

func (m *defaultPostsModel) PostsWithTag(offset, size int64, tag string) ([]Posts, error) {
	var ret []Posts
	query := fmt.Sprintf("select * from %s where tags like '%%%s%%' limit %d offset %d", m.table, tag, size, offset)
	err := m.QueryRowsNoCache(&ret, query)

	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (m *defaultPostsModel) Posts(offset, size int64) ([]Posts, error) {
	var ret []Posts
	query := fmt.Sprintf("select * from %s limit %d offset %d", m.table, size, offset)
	err := m.QueryRowsNoCache(&ret, query)

	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (m *defaultPostsModel) CountWtihTag(tag string) (int64, error) {
	var ret int64
	query := fmt.Sprintf("select count(1) from %s where tags like '%%%s%%'", m.table, tag)
	err := m.QueryRowNoCache(&ret, query)

	if err != nil {
		return -1, err
	}

	return ret, nil
}

func (m *defaultPostsModel) Count() (int64, error) {
	var ret int64
	query := fmt.Sprintf("select count(1) from %s", m.table)
	err := m.QueryRowNoCache(&ret, query)

	if err != nil {
		return -1, err
	}

	return ret, nil
}

func (m *defaultPostsModel) Insert(data Posts) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, postsRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Id, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Title, data.Description, data.Pre, data.Next, data.Url, data.Path, data.Tags)

	return ret, err
}

func (m *defaultPostsModel) FindOne(id int64) (*Posts, error) {
	blogPostsIdKey := fmt.Sprintf("%s%v", cacheBlogPostsIdPrefix, id)
	var resp Posts
	err := m.QueryRow(&resp, blogPostsIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", postsRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPostsModel) Update(data Posts) error {
	blogPostsIdKey := fmt.Sprintf("%s%v", cacheBlogPostsIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, postsRowsWithPlaceHolder)
		return conn.Exec(query, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Title, data.Description, data.Pre, data.Next, data.Url, data.Path, data.Tags, data.Id)
	}, blogPostsIdKey)
	return err
}

func (m *defaultPostsModel) Delete(id int64) error {

	blogPostsIdKey := fmt.Sprintf("%s%v", cacheBlogPostsIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, blogPostsIdKey)
	return err
}

func (m *defaultPostsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBlogPostsIdPrefix, primary)
}

func (m *defaultPostsModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", postsRows, m.table)
	return conn.QueryRow(v, query, primary)
}
