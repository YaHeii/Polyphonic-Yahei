package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TOperationLogModel = (*customTOperationLogModel)(nil)

type (
	// TOperationLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTOperationLogModel.
	TOperationLogModel interface {
		tOperationLogModel
		withSession(session sqlx.Session) TOperationLogModel
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
