package model

import (
	"context"
	"fmt"

	"github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TArticleModel = (*customTArticleModel)(nil)

type (
	// TArticleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTArticleModel.
	TArticleModel interface {
		tArticleModel
		FindById(ctx context.Context, id int64) (*TArticle, error)
		FindByIds(ctx context.Context, ids []int64) ([]*TArticle, error)
		CountGroupByCategoryIDs(ctx context.Context, ids []int64) (map[int64]int64, error)
		CountGroupByTagNames(ctx context.Context, names []string) (map[string]int64, error)
		GetDailyStatistics(ctx context.Context) (map[string]int64, error)
	}

	customTArticleModel struct {
		*defaultTArticleModel
	}
)

// NewTArticleModel returns a model for the database table.
func NewTArticleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TArticleModel {
	return &customTArticleModel{
		defaultTArticleModel: newTArticleModel(conn, c, opts...),
	}
}

func (m *customTArticleModel) FindById(ctx context.Context, id int64) (*TArticle, error) {
	return m.FindOne(ctx, id)
}

func (m *customTArticleModel) FindByIds(ctx context.Context, ids []int64) ([]*TArticle, error) {
	if len(ids) == 0 {
		return []*TArticle{}, nil
	}

	query := fmt.Sprintf("select %s from %s where id = any($1)", tArticleRows, m.table)
	var list []*TArticle
	if err := m.QueryRowsNoCacheCtx(ctx, &list, query, pq.Array(ids)); err != nil {
		return nil, err
	}

	return list, nil
}

func (m *customTArticleModel) CountGroupByCategoryIDs(ctx context.Context, ids []int64) (map[int64]int64, error) {
	if len(ids) == 0 {
		return map[int64]int64{}, nil
	}

	var rows []struct {
		CategoryID   int64 `db:"category_id"`
		ArticleCount int64 `db:"article_count"`
	}
	query := fmt.Sprintf("select category_id, count(*) as article_count from %s where category_id = any($1) group by category_id order by category_id", m.table)
	if err := m.QueryRowsNoCacheCtx(ctx, &rows, query, pq.Array(ids)); err != nil {
		return nil, err
	}

	result := make(map[int64]int64, len(rows))
	for _, row := range rows {
		result[row.CategoryID] = row.ArticleCount
	}
	return result, nil
}

func (m *customTArticleModel) CountGroupByTagNames(ctx context.Context, names []string) (map[string]int64, error) {
	if len(names) == 0 {
		return map[string]int64{}, nil
	}

	var rows []struct {
		TagName      string `db:"tag_name"`
		ArticleCount int64  `db:"article_count"`
	}
	query := fmt.Sprintf(`select tag_name, count(*) as article_count
from (select unnest(tags) as tag_name from %s) t
where tag_name = any($1)
group by tag_name
order by tag_name`, m.table)
	if err := m.QueryRowsNoCacheCtx(ctx, &rows, query, pq.Array(names)); err != nil {
		return nil, err
	}

	result := make(map[string]int64, len(rows))
	for _, row := range rows {
		result[row.TagName] = row.ArticleCount
	}
	return result, nil
}

func (m *customTArticleModel) GetDailyStatistics(ctx context.Context) (map[string]int64, error) {
	var rows []struct {
		Date         string `db:"date"`
		ArticleCount int64  `db:"article_count"`
	}
	query := fmt.Sprintf(`select to_char(created_at, 'YYYY-MM-DD') as date, count(*) as article_count
from %s
group by date
order by date desc`, m.table)
	if err := m.QueryRowsNoCacheCtx(ctx, &rows, query); err != nil {
		return nil, err
	}

	result := make(map[string]int64, len(rows))
	for _, row := range rows {
		result[row.Date] = row.ArticleCount
	}
	return result, nil
}
