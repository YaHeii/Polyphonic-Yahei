package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TRoleApiModel = (*customTRoleApiModel)(nil)

type (
	// TRoleApiModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTRoleApiModel.
	TRoleApiModel interface {
		tRoleApiModel
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
