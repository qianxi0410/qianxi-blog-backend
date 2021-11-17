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
	visitFieldNames          = builderx.RawFieldNames(&Visit{})
	visitRows                = strings.Join(visitFieldNames, ",")
	visitRowsExpectAutoSet   = strings.Join(stringx.Remove(visitFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	visitRowsWithPlaceHolder = strings.Join(stringx.Remove(visitFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	VisitModel interface {
		Insert(data Visit) (sql.Result, error)
		FindOne(id int64) (*Visit, error)
		Update(data Visit) error
		Delete(id int64) error
		Count() (int64, error)
		PeopleCount() (int64, error)
		WeekCount() ([]WeekCount, error)
	}

	defaultVisitModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Visit struct {
		Id        int64     `db:"id"`
		VisitTime time.Time `db:"visit_time"`
		Ip        string    `db:"ip"`
		Uri       string    `db:"uri"`
	}
)

func NewVisitModel(conn sqlx.SqlConn) VisitModel {
	return &defaultVisitModel{
		conn:  conn,
		table: "`visit`",
	}
}

func (m *defaultVisitModel) WeekCount() ([]WeekCount, error) {
	var weekCounts []WeekCount

	query := "select day(visit_time) day,COUNT(*) count from visit where date_sub(curdate(), interval 7 day ) <= date(visit_time) group by day(visit_time)"
	err := m.conn.QueryRows(&weekCounts, query)
	if err != nil {
		return nil, err
	}

	return weekCounts, nil
}

func (m *defaultVisitModel) PeopleCount() (int64, error) {
	var ret int64
	query := fmt.Sprintf("select COUNT(1) from %s where uri = ?", m.table)
	err := m.conn.QueryRow(&ret, query, "/")
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func (m *defaultVisitModel) Count() (int64, error) {
	var ret int64
	query := fmt.Sprintf("select COUNT(1) from %s", m.table)
	err := m.conn.QueryRow(&ret, query)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func (m *defaultVisitModel) Insert(data Visit) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, visitRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.VisitTime, data.Ip, data.Uri)
	return ret, err
}

func (m *defaultVisitModel) FindOne(id int64) (*Visit, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", visitRows, m.table)
	var resp Visit
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

func (m *defaultVisitModel) Update(data Visit) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, visitRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.VisitTime, data.Ip, data.Uri, data.Id)
	return err
}

func (m *defaultVisitModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
