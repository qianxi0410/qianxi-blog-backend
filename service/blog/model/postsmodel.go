package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	postsFieldNames          = builderx.RawFieldNames(&Posts{})
	postsRows                = strings.Join(postsFieldNames, ",")
	postsRowsExpectAutoSet   = strings.Join(stringx.Remove(postsFieldNames), ",")
	postsRowsWithPlaceHolder = strings.Join(stringx.Remove(postsFieldNames, "`id`", "`created_at`"), "=?,") + "=?"

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
		PostsAll() ([]Posts, error)
		MaxId() (int64, error)
	}

	defaultPostsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Posts struct {
		Id          int64          `db:"id" json:"id"`
		CreatedAt   time.Time      `db:"created_at" json:"created_at"`
		UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`
		Title       string         `db:"title" json:"title"`
		Description sql.NullString `db:"description" json:"description"`
		Pre         int64          `db:"pre" json:"pre"`
		Next        int64          `db:"next" json:"next"`
		Url         string         `db:"url" json:"url"`
		Path        string         `db:"path" json:"path"`
		Tags        sql.NullString `db:"tags" json:"tags"`
	}
)

func NewPostsModel(conn sqlx.SqlConn) PostsModel {
	return &defaultPostsModel{
		conn:  conn,
		table: "`posts`",
	}
}

func (m *defaultPostsModel) MaxId() (int64, error) {
	var ret int64
	query := fmt.Sprintf("select max(id) from %s", m.table)
	err := m.conn.QueryRow(&ret, query)

	if err != nil {
		return -1, err
	}

	return ret, nil
}

func (m *defaultPostsModel) Title(id int64) (string, error) {
	var ret string

	query := fmt.Sprintf("select title from %s where id = %d", m.table, id)
	err := m.conn.QueryRow(&ret, query)

	if err != nil {
		return "", err
	}

	return ret, nil
}

func (m *defaultPostsModel) PostsWithTag(offset, size int64, tag string) ([]Posts, error) {
	var ret []Posts
	query := fmt.Sprintf("select * from %s where tags like '%%%s%%' order by created_at desc limit %d offset %d", m.table, tag, size, offset)
	err := m.conn.QueryRows(&ret, query)

	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (m *defaultPostsModel) PostsAll() ([]Posts, error) {
	var ret []Posts
	query := fmt.Sprintf("select * from %s order by created_at desc", m.table)
	err := m.conn.QueryRows(&ret, query)

	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (m *defaultPostsModel) Posts(offset, size int64) ([]Posts, error) {
	var ret []Posts
	query := fmt.Sprintf("select * from %s order by created_at desc limit %d offset %d", m.table, size, offset)
	err := m.conn.QueryRows(&ret, query)

	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (m *defaultPostsModel) CountWtihTag(tag string) (int64, error) {
	var ret int64
	query := fmt.Sprintf("select count(1) from %s where tags like '%%%s%%'", m.table, tag)
	err := m.conn.QueryRow(&ret, query)

	if err != nil {
		return -1, err
	}

	return ret, nil
}

func (m *defaultPostsModel) Count() (int64, error) {
	var ret int64
	query := fmt.Sprintf("select count(1) from %s", m.table)
	err := m.conn.QueryRow(&ret, query)

	if err != nil {
		return -1, err
	}

	return ret, nil
}

func (m *defaultPostsModel) Insert(data Posts) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, postsRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Id, data.CreatedAt, data.UpdatedAt, data.Title, data.Description, data.Pre, data.Next, data.Url, data.Path, data.Tags)
	fmt.Println(query)
	return ret, err
}

func (m *defaultPostsModel) FindOne(id int64) (*Posts, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", postsRows, m.table)
	var resp Posts
	err := m.conn.QueryRow(&resp, query, id)
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
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, postsRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.CreatedAt, data.UpdatedAt, data.Title, data.Description, data.Pre, data.Next, data.Url, data.Path, data.Tags, data.Id)
	return err
}

func (m *defaultPostsModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultPostsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBlogPostsIdPrefix, primary)
}

func (m *defaultPostsModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", postsRows, m.table)
	return conn.QueryRow(v, query, primary)
}
