package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TVisitorModel = (*customTVisitorModel)(nil)

type (
	// TVisitorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTVisitorModel.
	TVisitorModel interface {
		tVisitorModel
		withSession(session sqlx.Session) TVisitorModel
	}

	customTVisitorModel struct {
		*defaultTVisitorModel
	}
)

// NewTVisitorModel returns a model for the database table.
func NewTVisitorModel(conn sqlx.SqlConn) TVisitorModel {
	return &customTVisitorModel{
		defaultTVisitorModel: newTVisitorModel(conn),
	}
}

func (m *customTVisitorModel) withSession(session sqlx.Session) TVisitorModel {
	return NewTVisitorModel(sqlx.NewSqlConnFromSession(session))
}
