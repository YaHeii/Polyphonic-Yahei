package model

import "context"

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TUserRoleModel = (*customTUserRoleModel)(nil)

type (
	// TUserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTUserRoleModel.
	TUserRoleModel interface {
		tUserRoleModel
		FindByUserID(ctx context.Context, userID string) ([]*TUserRole, error)
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
