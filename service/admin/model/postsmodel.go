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
	postsRowsExpectAutoSet   = strings.Join(stringx.Remove(postsFieldNames, "`create_time`", "`update_time`"), ",")
	postsRowsWithPlaceHolder = strings.Join(stringx.Remove(postsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	PostsModel interface {
		Insert(data Posts) (sql.Result, error)
		FindOne(id int64) (*Posts, error)
		Update(data Posts) error
		Delete(id int64) error
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
		Blur        int32          `db:"blur" json:"blur"`
	}
)

func NewPostsModel(conn sqlx.SqlConn) PostsModel {
	return &defaultPostsModel{
		conn:  conn,
		table: "`posts`",
	}
}

func (m *defaultPostsModel) Insert(data Posts) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, postsRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Id, data.CreatedAt, data.UpdatedAt, data.Title, data.Description, data.Pre, data.Next, data.Url, data.Path, data.Tags, data.Blur)
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
	_, err := m.conn.Exec(query, data.CreatedAt, data.UpdatedAt, data.Title, data.Description, data.Pre, data.Next, data.Url, data.Path, data.Tags, data.Blur, data.Id)
	return err
}

func (m *defaultPostsModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
