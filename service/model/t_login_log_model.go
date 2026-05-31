package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TLoginLogModel = (*customTLoginLogModel)(nil)

type (
	// TLoginLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTLoginLogModel.
	TLoginLogModel interface {
		tLoginLogModel
		withSession(session sqlx.Session) TLoginLogModel
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
