package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TRoleMenuModel = (*customTRoleMenuModel)(nil)

type (
	// TRoleMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTRoleMenuModel.
	TRoleMenuModel interface {
		tRoleMenuModel
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
