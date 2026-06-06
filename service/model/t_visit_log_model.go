package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TVisitLogModel = (*customTVisitLogModel)(nil)

type (
	// TVisitLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTVisitLogModel.
	TVisitLogModel interface {
		tVisitLogModel
		withSession(session sqlx.Session) TVisitLogModel
		FindById(ctx context.Context, id int64) (*TVisitLog, error)
		FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TVisitLog, int64, error)
		Save(ctx context.Context, data *TVisitLog) (int64, error)
		Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error)
	}

	customTVisitLogModel struct {
		*defaultTVisitLogModel
	}
)

// NewTVisitLogModel returns a model for the database table.
func NewTVisitLogModel(conn sqlx.SqlConn) TVisitLogModel {
	return &customTVisitLogModel{
		defaultTVisitLogModel: newTVisitLogModel(conn),
	}
}

func (m *customTVisitLogModel) withSession(session sqlx.Session) TVisitLogModel {
	return NewTVisitLogModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customTVisitLogModel) FindById(ctx context.Context, id int64) (*TVisitLog, error) {
	return m.FindOne(ctx, id)
}

func (m *customTVisitLogModel) FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TVisitLog, int64, error) {
	whereClause, bindArgs := buildSyslogWhereClause(conditions, args...)

	countQuery := fmt.Sprintf("select count(*) from %s", m.table)
	if whereClause != "" {
		countQuery += " where " + whereClause
	}
	var total int64
	if err := m.conn.QueryRowCtx(ctx, &total, countQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf("select %s from %s", tVisitLogRows, m.table)
	if whereClause != "" {
		listQuery += " where " + whereClause
	}
	if sorts != "" {
		listQuery += " order by " + sorts
	}
	if page > 0 && size > 0 {
		offset := (page - 1) * size
		listQuery += fmt.Sprintf(" limit $%d offset $%d", len(bindArgs)+1, len(bindArgs)+2)
		bindArgs = append(bindArgs, size, offset)
	}

	var list []*TVisitLog
	if err := m.conn.QueryRowsCtx(ctx, &list, listQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (m *customTVisitLogModel) Save(ctx context.Context, data *TVisitLog) (int64, error) {
	if data.Id > 0 {
		query := fmt.Sprintf("insert into %s (id, user_id, terminal_id, page_name) values ($1, $2, $3, $4) returning id", m.table)
		if err := m.conn.QueryRowCtx(ctx, &data.Id, query, data.Id, data.UserId, data.TerminalId, data.PageName); err != nil {
			return 0, err
		}
		return data.Id, nil
	}

	query := fmt.Sprintf("insert into %s (user_id, terminal_id, page_name) values ($1, $2, $3) returning id", m.table)
	if err := m.conn.QueryRowCtx(ctx, &data.Id, query, data.UserId, data.TerminalId, data.PageName); err != nil {
		return 0, err
	}
	return data.Id, nil
}

func (m *customTVisitLogModel) Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error) {
	whereClause, bindArgs := buildSyslogWhereClause(conditions, args...)

	query := fmt.Sprintf("delete from %s", m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}

	return rowsAffected(m.conn.ExecCtx(ctx, query, bindArgs...))
}
