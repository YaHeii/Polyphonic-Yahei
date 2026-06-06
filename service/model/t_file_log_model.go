package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TFileLogModel = (*customTFileLogModel)(nil)

type (
	// TFileLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTFileLogModel.
	TFileLogModel interface {
		tFileLogModel
		withSession(session sqlx.Session) TFileLogModel
		FindById(ctx context.Context, id int64) (*TFileLog, error)
		FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TFileLog, int64, error)
		Save(ctx context.Context, data *TFileLog) (int64, error)
		Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error)
	}

	customTFileLogModel struct {
		*defaultTFileLogModel
	}
)

// NewTFileLogModel returns a model for the database table.
func NewTFileLogModel(conn sqlx.SqlConn) TFileLogModel {
	return &customTFileLogModel{
		defaultTFileLogModel: newTFileLogModel(conn),
	}
}

func (m *customTFileLogModel) withSession(session sqlx.Session) TFileLogModel {
	return NewTFileLogModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customTFileLogModel) FindById(ctx context.Context, id int64) (*TFileLog, error) {
	return m.FindOne(ctx, id)
}

func (m *customTFileLogModel) FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TFileLog, int64, error) {
	whereClause, bindArgs := buildSyslogWhereClause(conditions, args...)

	countQuery := fmt.Sprintf("select count(*) from %s", m.table)
	if whereClause != "" {
		countQuery += " where " + whereClause
	}
	var total int64
	if err := m.conn.QueryRowCtx(ctx, &total, countQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf("select %s from %s", tFileLogRows, m.table)
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

	var list []*TFileLog
	if err := m.conn.QueryRowsCtx(ctx, &list, listQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (m *customTFileLogModel) Save(ctx context.Context, data *TFileLog) (int64, error) {
	if data.Id > 0 {
		query := fmt.Sprintf("insert into %s (id, user_id, terminal_id, file_path, file_name, file_type, file_size, file_md5, file_url) values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id", m.table)
		if err := m.conn.QueryRowCtx(ctx, &data.Id, query, data.Id, data.UserId, data.TerminalId, data.FilePath, data.FileName, data.FileType, data.FileSize, data.FileMd5, data.FileUrl); err != nil {
			return 0, err
		}
		return data.Id, nil
	}

	query := fmt.Sprintf("insert into %s (user_id, terminal_id, file_path, file_name, file_type, file_size, file_md5, file_url) values ($1, $2, $3, $4, $5, $6, $7, $8) returning id", m.table)
	if err := m.conn.QueryRowCtx(ctx, &data.Id, query, data.UserId, data.TerminalId, data.FilePath, data.FileName, data.FileType, data.FileSize, data.FileMd5, data.FileUrl); err != nil {
		return 0, err
	}
	return data.Id, nil
}

func (m *customTFileLogModel) Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error) {
	whereClause, bindArgs := buildSyslogWhereClause(conditions, args...)

	query := fmt.Sprintf("delete from %s", m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}

	return rowsAffected(m.conn.ExecCtx(ctx, query, bindArgs...))
}
