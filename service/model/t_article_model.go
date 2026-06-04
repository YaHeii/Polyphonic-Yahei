package model

import (
	"context"
	"fmt"
	"strings"

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
		FindALL(ctx context.Context, conditions string, args ...interface{}) ([]*TArticle, error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (int64, error)
		FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*TArticle, total int64, err error)
		Deletes(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		Updates(ctx context.Context, columns map[string]interface{}, conditions string, args ...interface{}) (rows int64, err error)
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

func (m *customTArticleModel) FindALL(ctx context.Context, conditions string, args ...interface{}) ([]*TArticle, error) {
	whereClause, bindArgs := buildArticleWhereClause(conditions, args...)

	query := fmt.Sprintf("select %s from %s", tArticleRows, m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}

	var list []*TArticle
	if err := m.QueryRowsNoCacheCtx(ctx, &list, query, bindArgs...); err != nil {
		return nil, err
	}

	return list, nil
}

func (m *customTArticleModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (int64, error) {
	whereClause, bindArgs := buildArticleWhereClause(conditions, args...)

	query := fmt.Sprintf("select count(*) from %s", m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}

	var total int64
	if err := m.QueryRowNoCacheCtx(ctx, &total, query, bindArgs...); err != nil {
		return 0, err
	}

	return total, nil
}

func (m *customTArticleModel) FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*TArticle, total int64, err error) {
	whereClause, bindArgs := buildArticleWhereClause(conditions, args...)

	countQuery := fmt.Sprintf("select count(*) from %s", m.table)
	if whereClause != "" {
		countQuery += " where " + whereClause
	}
	if err = m.QueryRowNoCacheCtx(ctx, &total, countQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf("select %s from %s", tArticleRows, m.table)
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
	if err = m.QueryRowsNoCacheCtx(ctx, &list, listQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (m *customTArticleModel) Deletes(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	whereClause, bindArgs := buildArticleWhereClause(conditions, args...)

	query := fmt.Sprintf("delete from %s", m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}

	result, err := m.ExecNoCacheCtx(ctx, query, bindArgs...)
	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affected, nil
}

func (m *customTArticleModel) Updates(ctx context.Context, columns map[string]interface{}, conditions string, args ...interface{}) (rows int64, err error) {
	if len(columns) == 0 {
		return 0, nil
	}

	setClauses := make([]string, 0, len(columns))
	bindArgs := make([]interface{}, 0, len(columns)+len(args))
	index := 1
	for column, value := range columns {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", column, index))
		bindArgs = append(bindArgs, value)
		index++
	}

	whereClause, whereArgs := buildArticleWhereClauseWithStartIndex(conditions, index, args...)
	bindArgs = append(bindArgs, whereArgs...)

	query := fmt.Sprintf("update %s set %s", m.table, strings.Join(setClauses, ", "))
	if whereClause != "" {
		query += " where " + whereClause
	}

	result, err := m.ExecNoCacheCtx(ctx, query, bindArgs...)
	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affected, nil
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

func buildArticleWhereClause(conditions string, args ...interface{}) (string, []interface{}) {
	return buildArticleWhereClauseWithStartIndex(conditions, 1, args...)
}

func buildArticleWhereClauseWithStartIndex(conditions string, start int, args ...interface{}) (string, []interface{}) {
	if conditions == "" {
		return "", args
	}

	var builder strings.Builder
	index := start
	for _, ch := range conditions {
		if ch == '?' {
			builder.WriteString(fmt.Sprintf("$%d", index))
			index++
			continue
		}
		builder.WriteRune(ch)
	}

	return builder.String(), args
}
