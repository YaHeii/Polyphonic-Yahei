package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TFileLogModel = (*customTFileLogModel)(nil)

type (
	// TFileLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTFileLogModel.
	TFileLogModel interface {
		tFileLogModel
		withSession(session sqlx.Session) TFileLogModel
	}

	customTFileLogModel struct {
		*defaultTFileLogModel
	}
)

// NewTFileLogModel returns a model for the database table.
func NewTFileLogModel(conn sqlx.SqlConn) TFileLogModel {
	return &customTFileLogModel{
		defaultTFileLogModel: newTFileLogModel(conn),
	}
}

func (m *customTFileLogModel) withSession(session sqlx.Session) TFileLogModel {
	return NewTFileLogModel(sqlx.NewSqlConnFromSession(session))
}
