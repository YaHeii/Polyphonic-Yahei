package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TPageModel = (*customTPageModel)(nil)

type (
	// TPageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTPageModel.
	TPageModel interface {
		tPageModel
		FindById(ctx context.Context, id int64) (*TPage, error)
		FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TPage, int64, error)
		Save(ctx context.Context, data *TPage) (int64, error)
		Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error)
	}

	customTPageModel struct {
		*defaultTPageModel
	}
)

// NewTPageModel returns a model for the database table.
func NewTPageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TPageModel {
	return &customTPageModel{
		defaultTPageModel: newTPageModel(conn, c, opts...),
	}
}

func (m *customTPageModel) FindById(ctx context.Context, id int64) (*TPage, error) {
	return m.FindOne(ctx, id)
}

func (m *customTPageModel) FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TPage, int64, error) {
	whereClause, bindArgs := buildResourceWhereClause(conditions, args...)

	countQuery := fmt.Sprintf("select count(*) from %s", m.table)
	if whereClause != "" {
		countQuery += " where " + whereClause
	}
	var total int64
	if err := m.QueryRowNoCacheCtx(ctx, &total, countQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf("select %s from %s", tPageRows, m.table)
	if whereClause != "" {
		listQuery += " where " + whereClause
	}
	if sorts != "" {
		listQuery += " order by " + sorts
	}
	if page > 0 && size > 0 {
		offset := (page - 1) * size
		listQuery += fmt.Sprintf(" limit $%d offset $%d", len(bindArgs)+1, len(bindArgs)+2)
		bindArgs = append(bindArgs, size, offset)
	}

	var list []*TPage
	if err := m.QueryRowsNoCacheCtx(ctx, &list, listQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (m *customTPageModel) Save(ctx context.Context, data *TPage) (int64, error) {
	if data.Id > 0 {
		query := fmt.Sprintf(`insert into %s (id, page_name, page_label, page_cover, is_carousel, carousel_covers)
values ($1, $2, $3, $4, $5, $6)
on conflict (id) do update set
	page_name = excluded.page_name,
	page_label = excluded.page_label,
	page_cover = excluded.page_cover,
	is_carousel = excluded.is_carousel,
	carousel_covers = excluded.carousel_covers,
	updated_at = now()
returning id`, m.table)
		if err := m.QueryRowNoCacheCtx(ctx, &data.Id, query, data.Id, data.PageName, data.PageLabel, data.PageCover, data.IsCarousel, data.CarouselCovers); err != nil {
			return 0, err
		}
		return data.Id, nil
	}

	query := fmt.Sprintf(`insert into %s (page_name, page_label, page_cover, is_carousel, carousel_covers)
values ($1, $2, $3, $4, $5)
returning id`, m.table)
	if err := m.QueryRowNoCacheCtx(ctx, &data.Id, query, data.PageName, data.PageLabel, data.PageCover, data.IsCarousel, data.CarouselCovers); err != nil {
		return 0, err
	}

	return data.Id, nil
}

func (m *customTPageModel) Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error) {
	whereClause, bindArgs := buildResourceWhereClause(conditions, args...)

	query := fmt.Sprintf("delete from %s", m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}

	return rowsAffected(m.ExecNoCacheCtx(ctx, query, bindArgs...))
}
