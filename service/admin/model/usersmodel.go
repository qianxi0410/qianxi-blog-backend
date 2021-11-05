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
	usersFieldNames          = builderx.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	UsersModel interface {
		Insert(data Users) (sql.Result, error)
		FindOne(id int64) (*Users, error)
		FindOneByName(name string) (*Users, error)
		FindOneByPassword(password string) (*Users, error)
		Update(data Users) error
		Delete(id int64) error
		FindWithNameAndPass(username string, password string) (*Users, error)
	}

	defaultUsersModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Users struct {
		Id       int64  `db:"id"`
		Name     string `db:"name"`
		Password string `db:"password"`
	}
)

func NewUsersModel(conn sqlx.SqlConn) UsersModel {
	return &defaultUsersModel{
		conn:  conn,
		table: "`users`",
	}
}

func (m *defaultUsersModel) FindWithNameAndPass(username string, password string) (*Users, error) {
	query := fmt.Sprintf("select %s from %s where `name` = ? and `password` = ? limit 1", usersRows, m.table)
	var resp Users
	err := m.conn.QueryRow(&resp, query, username, password)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Insert(data Users) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, usersRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.Password)
	return ret, err
}

func (m *defaultUsersModel) FindOne(id int64) (*Users, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
	var resp Users
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

func (m *defaultUsersModel) FindOneByName(name string) (*Users, error) {
	var resp Users
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", usersRows, m.table)
	err := m.conn.QueryRow(&resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByPassword(password string) (*Users, error) {
	var resp Users
	query := fmt.Sprintf("select %s from %s where `password` = ? limit 1", usersRows, m.table)
	err := m.conn.QueryRow(&resp, query, password)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Update(data Users) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.Password, data.Id)
	return err
}

func (m *defaultUsersModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
