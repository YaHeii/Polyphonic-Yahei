package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TTagModel = (*customTTagModel)(nil)

type (
	// TTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTTagModel.
	TTagModel interface {
		tTagModel
		FindById(ctx context.Context, id int64) (*TTag, error)
		FindByNames(ctx context.Context, names []string) ([]*TTag, error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) ([]*TTag, error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (int64, error)
		FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*TTag, total int64, err error)
		Deletes(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
	}

	customTTagModel struct {
		*defaultTTagModel
	}
)

// NewTTagModel returns a model for the database table.
func NewTTagModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TTagModel {
	return &customTTagModel{
		defaultTTagModel: newTTagModel(conn, c, opts...),
	}
}

func (m *customTTagModel) FindById(ctx context.Context, id int64) (*TTag, error) {
	return m.FindOne(ctx, id)
}

func (m *customTTagModel) FindByNames(ctx context.Context, names []string) ([]*TTag, error) {
	if len(names) == 0 {
		return []*TTag{}, nil
	}

	query := fmt.Sprintf("select %s from %s where tag_name = any($1)", tTagRows, m.table)
	var list []*TTag
	if err := m.QueryRowsNoCacheCtx(ctx, &list, query, names); err != nil {
		return nil, err
	}

	return list, nil
}

func (m *customTTagModel) FindALL(ctx context.Context, conditions string, args ...interface{}) ([]*TTag, error) {
	whereClause, bindArgs := buildPostgresWhereClause(conditions, args...)

	query := fmt.Sprintf("select %s from %s", tTagRows, m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}

	var list []*TTag
	if err := m.QueryRowsNoCacheCtx(ctx, &list, query, bindArgs...); err != nil {
		return nil, err
	}

	return list, nil
}

func (m *customTTagModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (int64, error) {
	whereClause, bindArgs := buildPostgresWhereClause(conditions, args...)

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

// 分页查询记录
func (m *defaultTTagModel) FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*TTag, total int64, err error) {
	whereClause, bindArgs := buildPostgresWhereClause(conditions, args...)

	countQuery := fmt.Sprintf("select count(*) from %s", m.table)
	if whereClause != "" {
		countQuery += " where " + whereClause
	}
	if err = m.QueryRowNoCacheCtx(ctx, &total, countQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf("select %s from %s", tTagRows, m.table)
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

func buildPostgresWhereClause(conditions string, args ...interface{}) (string, []interface{}) {
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

// 删除记录（批量操作）
func (m *customTTagModel) Deletes(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	whereClause, bindArgs := buildPostgresWhereClause(conditions, args...)

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
