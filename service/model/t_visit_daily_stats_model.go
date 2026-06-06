package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TVisitDailyStatsModel = (*customTVisitDailyStatsModel)(nil)

type (
	// TVisitDailyStatsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTVisitDailyStatsModel.
	TVisitDailyStatsModel interface {
		tVisitDailyStatsModel
		withSession(session sqlx.Session) TVisitDailyStatsModel
		Increment(ctx context.Context, date string, visitType int64, delta int64) (int64, error)
		FindCount(ctx context.Context, date string, visitType int64) (int64, error)
		SumByVisitType(ctx context.Context, visitType int64) (int64, error)
		FindByDateRange(ctx context.Context, startDate string, endDate string, visitType int64) ([]*TVisitDailyStats, error)
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

func (m *customTVisitDailyStatsModel) Increment(ctx context.Context, date string, visitType int64, delta int64) (int64, error) {
	query := fmt.Sprintf(`insert into %s as stats (date, view_count, visit_type)
values ($1, $2, $3)
on conflict (date, visit_type) do update set
	view_count = stats.view_count + excluded.view_count,
	updated_at = now()
returning view_count`, m.table)

	var count int64
	if err := m.conn.QueryRowCtx(ctx, &count, query, date, delta, visitType); err != nil {
		return 0, err
	}
	return count, nil
}

func (m *customTVisitDailyStatsModel) FindCount(ctx context.Context, date string, visitType int64) (int64, error) {
	query := fmt.Sprintf("select coalesce(sum(view_count), 0) from %s where date = $1 and visit_type = $2", m.table)
	var count int64
	if err := m.conn.QueryRowCtx(ctx, &count, query, date, visitType); err != nil {
		return 0, err
	}
	return count, nil
}

func (m *customTVisitDailyStatsModel) SumByVisitType(ctx context.Context, visitType int64) (int64, error) {
	query := fmt.Sprintf("select coalesce(sum(view_count), 0) from %s where visit_type = $1", m.table)
	var count int64
	if err := m.conn.QueryRowCtx(ctx, &count, query, visitType); err != nil {
		return 0, err
	}
	return count, nil
}

func (m *customTVisitDailyStatsModel) FindByDateRange(ctx context.Context, startDate string, endDate string, visitType int64) ([]*TVisitDailyStats, error) {
	query := fmt.Sprintf("select %s from %s where date >= $1 and date <= $2 and visit_type = $3 order by date asc", tVisitDailyStatsRows, m.table)

	var list []*TVisitDailyStats
	if err := m.conn.QueryRowsCtx(ctx, &list, query, startDate, endDate, visitType); err != nil {
		return nil, err
	}
	return list, nil
}
