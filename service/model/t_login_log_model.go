package model

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TLoginLogModel = (*customTLoginLogModel)(nil)

type (
	// TLoginLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTLoginLogModel.
	TLoginLogModel interface {
		tLoginLogModel
		withSession(session sqlx.Session) TLoginLogModel
		FindById(ctx context.Context, id int64) (*TLoginLog, error)
		FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TLoginLog, int64, error)
		Save(ctx context.Context, data *TLoginLog) (int64, error)
		Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error)
		UpdateLatestLogout(ctx context.Context, userID string, logoutAt time.Time) (int64, error)
	}

	customTLoginLogModel struct {
		*defaultTLoginLogModel
	}
)

// NewTLoginLogModel returns a model for the database table.
func NewTLoginLogModel(conn sqlx.SqlConn) TLoginLogModel {
	return &customTLoginLogModel{
		defaultTLoginLogModel: newTLoginLogModel(conn),
	}
}

func (m *customTLoginLogModel) withSession(session sqlx.Session) TLoginLogModel {
	return NewTLoginLogModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customTLoginLogModel) FindById(ctx context.Context, id int64) (*TLoginLog, error) {
	return m.FindOne(ctx, id)
}

func (m *customTLoginLogModel) FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TLoginLog, int64, error) {
	whereClause, bindArgs := buildSyslogWhereClause(conditions, args...)

	countQuery := fmt.Sprintf("select count(*) from %s", m.table)
	if whereClause != "" {
		countQuery += " where " + whereClause
	}
	var total int64
	if err := m.conn.QueryRowCtx(ctx, &total, countQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf("select %s from %s", tLoginLogRows, m.table)
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

	var list []*TLoginLog
	if err := m.conn.QueryRowsCtx(ctx, &list, listQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (m *customTLoginLogModel) Save(ctx context.Context, data *TLoginLog) (int64, error) {
	if data.Id > 0 {
		query := fmt.Sprintf("insert into %s (id, user_id, terminal_id, login_type, app_name, login_at, logout_at) values ($1, $2, $3, $4, $5, $6, $7) returning id", m.table)
		if err := m.conn.QueryRowCtx(ctx, &data.Id, query, data.Id, data.UserId, data.TerminalId, data.LoginType, data.AppName, data.LoginAt, data.LogoutAt); err != nil {
			return 0, err
		}
		return data.Id, nil
	}

	query := fmt.Sprintf("insert into %s (user_id, terminal_id, login_type, app_name, login_at, logout_at) values ($1, $2, $3, $4, $5, $6) returning id", m.table)
	if err := m.conn.QueryRowCtx(ctx, &data.Id, query, data.UserId, data.TerminalId, data.LoginType, data.AppName, data.LoginAt, data.LogoutAt); err != nil {
		return 0, err
	}
	return data.Id, nil
}

func (m *customTLoginLogModel) Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error) {
	whereClause, bindArgs := buildSyslogWhereClause(conditions, args...)

	query := fmt.Sprintf("delete from %s", m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}

	return rowsAffected(m.conn.ExecCtx(ctx, query, bindArgs...))
}

func (m *customTLoginLogModel) UpdateLatestLogout(ctx context.Context, userID string, logoutAt time.Time) (int64, error) {
	query := fmt.Sprintf(`update %s
set logout_at = $1, updated_at = now()
where id = (
	select id from %s
	where user_id = $2 and logout_at is null
	order by login_at desc, id desc
	limit 1
)`, m.table, m.table)

	return rowsAffected(m.conn.ExecCtx(ctx, query, logoutAt, userID))
}
