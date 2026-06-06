package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TRoleMenuModel = (*customTRoleMenuModel)(nil)

type (
	// TRoleMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTRoleMenuModel.
	TRoleMenuModel interface {
		tRoleMenuModel
		FindMenuIDsByRoleID(ctx context.Context, roleID int64) ([]int64, error)
		ReplaceByRoleID(ctx context.Context, roleID int64, menuIDs []int64) (int64, error)
		DeleteByRoleIDs(ctx context.Context, roleIDs []int64) (int64, error)
		DeleteByMenuIDs(ctx context.Context, menuIDs []int64) (int64, error)
		Clean(ctx context.Context) (int64, error)
		withSession(session sqlx.Session) TRoleMenuModel
	}

	customTRoleMenuModel struct {
		*defaultTRoleMenuModel
	}
)

// NewTRoleMenuModel returns a model for the database table.
func NewTRoleMenuModel(conn sqlx.SqlConn) TRoleMenuModel {
	return &customTRoleMenuModel{
		defaultTRoleMenuModel: newTRoleMenuModel(conn),
	}
}

func (m *customTRoleMenuModel) withSession(session sqlx.Session) TRoleMenuModel {
	return NewTRoleMenuModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customTRoleMenuModel) FindMenuIDsByRoleID(ctx context.Context, roleID int64) ([]int64, error) {
	query := fmt.Sprintf("select menu_id from %s where role_id = $1 order by menu_id asc", m.table)

	var records []struct {
		MenuId int64 `db:"menu_id"`
	}
	if err := m.conn.QueryRowsCtx(ctx, &records, query, roleID); err != nil {
		return nil, err
	}

	ids := make([]int64, 0, len(records))
	for _, record := range records {
		ids = append(ids, record.MenuId)
	}

	return ids, nil
}

func (m *customTRoleMenuModel) ReplaceByRoleID(ctx context.Context, roleID int64, menuIDs []int64) (int64, error) {
	if _, err := m.conn.ExecCtx(ctx, fmt.Sprintf("delete from %s where role_id = $1", m.table), roleID); err != nil {
		return 0, err
	}

	var rows int64
	for _, menuID := range menuIDs {
		if menuID == 0 {
			continue
		}
		result, err := m.conn.ExecCtx(ctx, fmt.Sprintf("insert into %s (role_id, menu_id) values ($1, $2)", m.table), roleID, menuID)
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

func (m *customTRoleMenuModel) DeleteByRoleIDs(ctx context.Context, roleIDs []int64) (int64, error) {
	if len(roleIDs) == 0 {
		return 0, nil
	}

	query := fmt.Sprintf("delete from %s where role_id = any($1)", m.table)
	return rowsAffected(m.conn.ExecCtx(ctx, query, buildPermissionIDArray(roleIDs)))
}

func (m *customTRoleMenuModel) DeleteByMenuIDs(ctx context.Context, menuIDs []int64) (int64, error) {
	if len(menuIDs) == 0 {
		return 0, nil
	}

	query := fmt.Sprintf("delete from %s where menu_id = any($1)", m.table)
	return rowsAffected(m.conn.ExecCtx(ctx, query, buildPermissionIDArray(menuIDs)))
}

func (m *customTRoleMenuModel) Clean(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("delete from %s", m.table)
	return rowsAffected(m.conn.ExecCtx(ctx, query))
}
