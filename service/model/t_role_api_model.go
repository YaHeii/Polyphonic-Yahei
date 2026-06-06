package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TRoleApiModel = (*customTRoleApiModel)(nil)

type (
	// TRoleApiModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTRoleApiModel.
	TRoleApiModel interface {
		tRoleApiModel
		FindApiIDsByRoleID(ctx context.Context, roleID int64) ([]int64, error)
		ReplaceByRoleID(ctx context.Context, roleID int64, apiIDs []int64) (int64, error)
		DeleteByRoleIDs(ctx context.Context, roleIDs []int64) (int64, error)
		DeleteByApiIDs(ctx context.Context, apiIDs []int64) (int64, error)
		Clean(ctx context.Context) (int64, error)
		withSession(session sqlx.Session) TRoleApiModel
	}

	customTRoleApiModel struct {
		*defaultTRoleApiModel
	}
)

// NewTRoleApiModel returns a model for the database table.
func NewTRoleApiModel(conn sqlx.SqlConn) TRoleApiModel {
	return &customTRoleApiModel{
		defaultTRoleApiModel: newTRoleApiModel(conn),
	}
}

func (m *customTRoleApiModel) withSession(session sqlx.Session) TRoleApiModel {
	return NewTRoleApiModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customTRoleApiModel) FindApiIDsByRoleID(ctx context.Context, roleID int64) ([]int64, error) {
	query := fmt.Sprintf("select api_id from %s where role_id = $1 order by api_id asc", m.table)

	var records []struct {
		ApiId int64 `db:"api_id"`
	}
	if err := m.conn.QueryRowsCtx(ctx, &records, query, roleID); err != nil {
		return nil, err
	}

	ids := make([]int64, 0, len(records))
	for _, record := range records {
		ids = append(ids, record.ApiId)
	}

	return ids, nil
}

func (m *customTRoleApiModel) ReplaceByRoleID(ctx context.Context, roleID int64, apiIDs []int64) (int64, error) {
	if _, err := m.conn.ExecCtx(ctx, fmt.Sprintf("delete from %s where role_id = $1", m.table), roleID); err != nil {
		return 0, err
	}

	var rows int64
	for _, apiID := range apiIDs {
		if apiID == 0 {
			continue
		}
		result, err := m.conn.ExecCtx(ctx, fmt.Sprintf("insert into %s (role_id, api_id) values ($1, $2)", m.table), roleID, apiID)
		if err != nil {
			return rows, err
		}
		affected, err := result.RowsAffected()
		if err != nil {
			return rows, err
		}
		rows += affected
	}

	return rows, nil
}

func (m *customTRoleApiModel) DeleteByRoleIDs(ctx context.Context, roleIDs []int64) (int64, error) {
	if len(roleIDs) == 0 {
		return 0, nil
	}

	query := fmt.Sprintf("delete from %s where role_id = any($1)", m.table)
	return rowsAffected(m.conn.ExecCtx(ctx, query, buildPermissionIDArray(roleIDs)))
}

func (m *customTRoleApiModel) DeleteByApiIDs(ctx context.Context, apiIDs []int64) (int64, error) {
	if len(apiIDs) == 0 {
		return 0, nil
	}

	query := fmt.Sprintf("delete from %s where api_id = any($1)", m.table)
	return rowsAffected(m.conn.ExecCtx(ctx, query, buildPermissionIDArray(apiIDs)))
}

func (m *customTRoleApiModel) Clean(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("delete from %s", m.table)
	return rowsAffected(m.conn.ExecCtx(ctx, query))
}
