package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TUserRoleModel = (*customTUserRoleModel)(nil)

type (
	// TUserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTUserRoleModel.
	TUserRoleModel interface {
		tUserRoleModel
		FindByUserID(ctx context.Context, userID string) ([]*TUserRole, error)
		ReplaceByUserID(ctx context.Context, userID string, roleIDs []int64) (int64, error)
		DeleteByRoleIDs(ctx context.Context, roleIDs []int64) (int64, error)
		withSession(session sqlx.Session) TUserRoleModel
	}

	customTUserRoleModel struct {
		*defaultTUserRoleModel
	}
)

// NewTUserRoleModel returns a model for the database table.
func NewTUserRoleModel(conn sqlx.SqlConn) TUserRoleModel {
	return &customTUserRoleModel{
		defaultTUserRoleModel: newTUserRoleModel(conn),
	}
}

func (m *customTUserRoleModel) withSession(session sqlx.Session) TUserRoleModel {
	return NewTUserRoleModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customTUserRoleModel) FindByUserID(ctx context.Context, userID string) ([]*TUserRole, error) {
	query := "select " + tUserRoleRows + " from " + m.tableName() + " where user_id = $1 order by id asc"

	var list []*TUserRole
	if err := m.conn.QueryRowsCtx(ctx, &list, query, userID); err != nil {
		return nil, err
	}

	return list, nil
}

func (m *customTUserRoleModel) ReplaceByUserID(ctx context.Context, userID string, roleIDs []int64) (int64, error) {
	if _, err := m.conn.ExecCtx(ctx, fmt.Sprintf("delete from %s where user_id = $1", m.tableName()), userID); err != nil {
		return 0, err
	}

	var rows int64
	for _, roleID := range roleIDs {
		if roleID == 0 {
			continue
		}
		result, err := m.conn.ExecCtx(ctx, fmt.Sprintf("insert into %s (user_id, role_id) values ($1, $2)", m.tableName()), userID, roleID)
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

func (m *customTUserRoleModel) DeleteByRoleIDs(ctx context.Context, roleIDs []int64) (int64, error) {
	if len(roleIDs) == 0 {
		return 0, nil
	}

	query := fmt.Sprintf("delete from %s where role_id = any($1)", m.tableName())
	return rowsAffected(m.conn.ExecCtx(ctx, query, buildPermissionIDArray(roleIDs)))
}
