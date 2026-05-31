package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TVisitDailyStatsModel = (*customTVisitDailyStatsModel)(nil)

type (
	// TVisitDailyStatsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTVisitDailyStatsModel.
	TVisitDailyStatsModel interface {
		tVisitDailyStatsModel
		withSession(session sqlx.Session) TVisitDailyStatsModel
	}

	customTVisitDailyStatsModel struct {
		*defaultTVisitDailyStatsModel
	}
)

// NewTVisitDailyStatsModel returns a model for the database table.
func NewTVisitDailyStatsModel(conn sqlx.SqlConn) TVisitDailyStatsModel {
	return &customTVisitDailyStatsModel{
		defaultTVisitDailyStatsModel: newTVisitDailyStatsModel(conn),
	}
}

func (m *customTVisitDailyStatsModel) withSession(session sqlx.Session) TVisitDailyStatsModel {
	return NewTVisitDailyStatsModel(sqlx.NewSqlConnFromSession(session))
}
