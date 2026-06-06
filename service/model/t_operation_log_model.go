package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TOperationLogModel = (*customTOperationLogModel)(nil)

type (
	// TOperationLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTOperationLogModel.
	TOperationLogModel interface {
		tOperationLogModel
		withSession(session sqlx.Session) TOperationLogModel
		FindById(ctx context.Context, id int64) (*TOperationLog, error)
		FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TOperationLog, int64, error)
		Save(ctx context.Context, data *TOperationLog) (int64, error)
		Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error)
	}

	customTOperationLogModel struct {
		*defaultTOperationLogModel
	}
)

// NewTOperationLogModel returns a model for the database table.
func NewTOperationLogModel(conn sqlx.SqlConn) TOperationLogModel {
	return &customTOperationLogModel{
		defaultTOperationLogModel: newTOperationLogModel(conn),
	}
}

func (m *customTOperationLogModel) withSession(session sqlx.Session) TOperationLogModel {
	return NewTOperationLogModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customTOperationLogModel) FindById(ctx context.Context, id int64) (*TOperationLog, error) {
	return m.FindOne(ctx, id)
}

func (m *customTOperationLogModel) FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TOperationLog, int64, error) {
	whereClause, bindArgs := buildSyslogWhereClause(conditions, args...)

	countQuery := fmt.Sprintf("select count(*) from %s", m.table)
	if whereClause != "" {
		countQuery += " where " + whereClause
	}
	var total int64
	if err := m.conn.QueryRowCtx(ctx, &total, countQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf("select %s from %s", tOperationLogRows, m.table)
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

	var list []*TOperationLog
	if err := m.conn.QueryRowsCtx(ctx, &list, listQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (m *customTOperationLogModel) Save(ctx context.Context, data *TOperationLog) (int64, error) {
	if data.Id > 0 {
		query := fmt.Sprintf("insert into %s (id, user_id, terminal_id, opt_module, opt_desc, request_uri, request_method, request_data, response_data, response_status, cost) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) returning id", m.table)
		if err := m.conn.QueryRowCtx(ctx, &data.Id, query, data.Id, data.UserId, data.TerminalId, data.OptModule, data.OptDesc, data.RequestUri, data.RequestMethod, data.RequestData, data.ResponseData, data.ResponseStatus, data.Cost); err != nil {
			return 0, err
		}
		return data.Id, nil
	}

	query := fmt.Sprintf("insert into %s (user_id, terminal_id, opt_module, opt_desc, request_uri, request_method, request_data, response_data, response_status, cost) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning id", m.table)
	if err := m.conn.QueryRowCtx(ctx, &data.Id, query, data.UserId, data.TerminalId, data.OptModule, data.OptDesc, data.RequestUri, data.RequestMethod, data.RequestData, data.ResponseData, data.ResponseStatus, data.Cost); err != nil {
		return 0, err
	}
	return data.Id, nil
}

func (m *customTOperationLogModel) Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error) {
	whereClause, bindArgs := buildSyslogWhereClause(conditions, args...)

	query := fmt.Sprintf("delete from %s", m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}

	return rowsAffected(m.conn.ExecCtx(ctx, query, bindArgs...))
}
