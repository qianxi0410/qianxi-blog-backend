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
	commentsFieldNames          = builderx.RawFieldNames(&Comments{})
	commentsRows                = strings.Join(commentsFieldNames, ",")
	commentsRowsExpectAutoSet   = strings.Join(stringx.Remove(commentsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	commentsRowsWithPlaceHolder = strings.Join(stringx.Remove(commentsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	CommentsModel interface {
		Insert(data Comments) (sql.Result, error)
		FindOne(id int64) (*Comments, error)
		Update(data Comments) error
		Delete(id int64) error
	}

	defaultCommentsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Comments struct {
		Id         int64     `db:"id"`
		CreatedAt  time.Time `db:"created_at"`
		UpdateedAt time.Time `db:"updateed_at"`
		Content    string    `db:"content"`
		Login      string    `db:"login"`
		Name       string    `db:"name"`
		Avatar     string    `db:"avatar"`
		PostId     int64     `db:"post_id"`
	}
)

func NewCommentsModel(conn sqlx.SqlConn) CommentsModel {
	return &defaultCommentsModel{
		conn:  conn,
		table: "`comments`",
	}
}

func (m *defaultCommentsModel) Insert(data Comments) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, commentsRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.CreatedAt, data.UpdateedAt, data.Content, data.Login, data.Name, data.Avatar, data.PostId)
	return ret, err
}

func (m *defaultCommentsModel) FindOne(id int64) (*Comments, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", commentsRows, m.table)
	var resp Comments
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

func (m *defaultCommentsModel) Update(data Comments) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, commentsRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.CreatedAt, data.UpdateedAt, data.Content, data.Login, data.Name, data.Avatar, data.PostId, data.Id)
	return err
}

func (m *defaultCommentsModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
