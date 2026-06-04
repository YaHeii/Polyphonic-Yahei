package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TCategoryModel = (*customTCategoryModel)(nil)

type (
	// TCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTCategoryModel.
	TCategoryModel interface {
		tCategoryModel
		FindById(ctx context.Context, id int64) (*TCategory, error)
		FindByIds(ctx context.Context, ids []int64) ([]*TCategory, error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) ([]*TCategory, error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (int64, error)
		FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*TCategory, total int64, err error)
		Deletes(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
	}

	customTCategoryModel struct {
		*defaultTCategoryModel
	}
)

// NewTCategoryModel returns a model for the database table.
func NewTCategoryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TCategoryModel {
	return &customTCategoryModel{
		defaultTCategoryModel: newTCategoryModel(conn, c, opts...),
	}
}

func (m *customTCategoryModel) FindById(ctx context.Context, id int64) (*TCategory, error) {
	return m.FindOne(ctx, id)
}

func (m *customTCategoryModel) FindByIds(ctx context.Context, ids []int64) ([]*TCategory, error) {
	if len(ids) == 0 {
		return []*TCategory{}, nil
	}

	query := fmt.Sprintf("select %s from %s where id = any($1)", tCategoryRows, m.table)
	var list []*TCategory
	if err := m.QueryRowsNoCacheCtx(ctx, &list, query, pq.Array(ids)); err != nil {
		return nil, err
	}

	return list, nil
}

func (m *customTCategoryModel) FindALL(ctx context.Context, conditions string, args ...interface{}) ([]*TCategory, error) {
	whereClause, bindArgs := buildCategoryWhereClause(conditions, args...)

	query := fmt.Sprintf("select %s from %s", tCategoryRows, m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}

	var list []*TCategory
	if err := m.QueryRowsNoCacheCtx(ctx, &list, query, bindArgs...); err != nil {
		return nil, err
	}

	return list, nil
}

func (m *customTCategoryModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (int64, error) {
	whereClause, bindArgs := buildCategoryWhereClause(conditions, args...)

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

func (m *defaultTCategoryModel) FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*TCategory, total int64, err error) {
	whereClause, bindArgs := buildCategoryWhereClause(conditions, args...)

	countQuery := fmt.Sprintf("select count(*) from %s", m.table)
	if whereClause != "" {
		countQuery += " where " + whereClause
	}
	if err = m.QueryRowNoCacheCtx(ctx, &total, countQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf("select %s from %s", tCategoryRows, m.table)
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

func (m *customTCategoryModel) Deletes(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	whereClause, bindArgs := buildCategoryWhereClause(conditions, args...)

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

func buildCategoryWhereClause(conditions string, args ...interface{}) (string, []interface{}) {
	if conditions == "" {
		return "", args
	}

	var builder strings.Builder
	index := 1
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
