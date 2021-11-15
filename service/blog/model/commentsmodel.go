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
	commentsRowsExpectAutoSet   = strings.Join(stringx.Remove(commentsFieldNames, "`id`"), ",")
	commentsRowsWithPlaceHolder = strings.Join(stringx.Remove(commentsFieldNames, "`id`", "`created_at`", "`updated_at`"), "=?,") + "=?"

	cacheBlogCommentsIdPrefix = "cache:blog:comments:id:"
)

type (
	CommentsModel interface {
		Insert(data Comments) (sql.Result, error)
		FindOne(id int64) (*Comments, error)
		Update(data Comments) error
		Delete(id int64) error
		CommentsWithPostId(id int64) ([]Comments, error)
		Count() (int64, error)
		DeleteByPostId(postId int64) error
		Comments(size int64, offset int64) ([]Comments, error)
		CommentsAll() ([]Comments, error)
	}

	defaultCommentsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Comments struct {
		Id         int64     `db:"id" json:"id"`
		CreatedAt  time.Time `db:"created_at" json:"created_at"`
		UpdateedAt time.Time `db:"updateed_at" json:"updateed_at"`
		Content    string    `db:"content" json:"content"`
		Login      string    `db:"login" json:"login"`
		Name       string    `db:"name" json:"name"`
		Avatar     string    `db:"avatar" json:"avatar"`
		PostId     int64     `db:"post_id" json:"post_id"`
	}
)

func NewCommentsModel(conn sqlx.SqlConn) CommentsModel {
	return &defaultCommentsModel{
		conn:  conn,
		table: "`comments`",
	}
}
func (m *defaultCommentsModel) CommentsAll() ([]Comments, error) {
	var ret []Comments

	query := fmt.Sprintf("select * from %s", m.table)
	err := m.conn.QueryRows(&ret, query)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (m *defaultCommentsModel) Comments(size int64, offset int64) ([]Comments, error) {
	var ret []Comments

	query := fmt.Sprintf("select * from %s limit %d offset %d", m.table, size, offset)
	err := m.conn.QueryRows(&ret, query)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (m *defaultCommentsModel) DeleteByPostId(postId int64) error {
	query := fmt.Sprintf("delete from %s where `post_id` = ?", m.table)
	_, err := m.conn.Exec(query, postId)
	return err
}

func (m *defaultCommentsModel) Count() (int64, error) {
	var ret int64
	query := fmt.Sprintf("select count(1) from %s", m.table)
	err := m.conn.QueryRow(&ret, query)

	if err != nil {
		return -1, err
	}

	return ret, nil
}

func (m *defaultCommentsModel) CommentsWithPostId(postId int64) ([]Comments, error) {
	var ret []Comments

	query := fmt.Sprintf("select * from %s where post_id = %d", m.table, postId)
	err := m.conn.QueryRows(&ret, query)
	if err != nil {
		return nil, err
	}

	return ret, nil
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

func (m *defaultCommentsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBlogCommentsIdPrefix, primary)
}

func (m *defaultCommentsModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", commentsRows, m.table)
	return conn.QueryRow(v, query, primary)
}
