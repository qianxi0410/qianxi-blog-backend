package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	systemFieldNames          = builderx.RawFieldNames(&System{})
	systemRows                = strings.Join(systemFieldNames, ",")
	systemRowsExpectAutoSet   = strings.Join(stringx.Remove(systemFieldNames, "`create_time`", "`update_time`"), ",")
	systemRowsWithPlaceHolder = strings.Join(stringx.Remove(systemFieldNames, "`key`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SystemModel interface {
		Insert(data System) (sql.Result, error)
		FindOne(key string) (*System, error)
		Update(data System) error
		Delete(key string) error
		All() ([]System, error)
	}

	defaultSystemModel struct {
		conn  sqlx.SqlConn
		table string
	}

	System struct {
		Key   string `db:"key"`
		Value string `db:"value"`
	}
)

func NewSystemModel(conn sqlx.SqlConn) SystemModel {
	return &defaultSystemModel{
		conn:  conn,
		table: "`system`",
	}
}

func (m *defaultSystemModel) All() ([]System, error) {
	var ret []System
	query := fmt.Sprintf("select %s from %s", systemRows, m.table)
	err := m.conn.QueryRows(&ret, query)
	switch err {
	case nil:
		return ret, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSystemModel) Insert(data System) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, systemRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Key, data.Value)
	return ret, err
}

func (m *defaultSystemModel) FindOne(key string) (*System, error) {
	query := fmt.Sprintf("select %s from %s where `key` = ? limit 1", systemRows, m.table)
	var resp System
	err := m.conn.QueryRow(&resp, query, key)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSystemModel) Update(data System) error {
	query := fmt.Sprintf("update %s set %s where `key` = ?", m.table, systemRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Value, data.Key)
	return err
}

func (m *defaultSystemModel) Delete(key string) error {
	query := fmt.Sprintf("delete from %s where `key` = ?", m.table)
	_, err := m.conn.Exec(query, key)
	return err
}
