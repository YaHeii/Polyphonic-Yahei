package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TVisitLogModel = (*customTVisitLogModel)(nil)

type (
	// TVisitLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTVisitLogModel.
	TVisitLogModel interface {
		tVisitLogModel
		withSession(session sqlx.Session) TVisitLogModel
	}

	customTVisitLogModel struct {
		*defaultTVisitLogModel
	}
)

// NewTVisitLogModel returns a model for the database table.
func NewTVisitLogModel(conn sqlx.SqlConn) TVisitLogModel {
	return &customTVisitLogModel{
		defaultTVisitLogModel: newTVisitLogModel(conn),
	}
}

func (m *customTVisitLogModel) withSession(session sqlx.Session) TVisitLogModel {
	return NewTVisitLogModel(sqlx.NewSqlConnFromSession(session))
}
