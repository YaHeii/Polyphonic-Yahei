package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TSystemNoticeModel = (*customTSystemNoticeModel)(nil)

type (
	// TSystemNoticeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTSystemNoticeModel.
	TSystemNoticeModel interface {
		tSystemNoticeModel
		FindById(ctx context.Context, id int64) (*TSystemNotice, error)
		FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*TSystemNotice, total int64, err error)
		Deletes(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
	}

	customTSystemNoticeModel struct {
		*defaultTSystemNoticeModel
	}
)

// NewTSystemNoticeModel returns a model for the database table.
func NewTSystemNoticeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TSystemNoticeModel {
	return &customTSystemNoticeModel{
		defaultTSystemNoticeModel: newTSystemNoticeModel(conn, c, opts...),
	}
}

func (m *customTSystemNoticeModel) FindById(ctx context.Context, id int64) (*TSystemNotice, error) {
	return m.FindOne(ctx, id)
}

func (m *customTSystemNoticeModel) FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*TSystemNotice, total int64, err error) {
	whereClause, bindArgs := buildSystemNoticeWhereClause(conditions, args...)

	countQuery := fmt.Sprintf("select count(*) from %s", m.table)
	if whereClause != "" {
		countQuery += " where " + whereClause
	}
	if err = m.QueryRowNoCacheCtx(ctx, &total, countQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf("select %s from %s", tSystemNoticeRows, m.table)
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

func (m *customTSystemNoticeModel) Deletes(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	whereClause, bindArgs := buildSystemNoticeWhereClause(conditions, args...)

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

func buildSystemNoticeWhereClause(conditions string, args ...interface{}) (string, []interface{}) {
	if conditions == "" {
		return "", args
	}

	replaced := strings.ReplaceAll(conditions, "id in (?)", "id = any(?)")
	normalizedArgs := make([]interface{}, 0, len(args))
	for _, arg := range args {
		switch v := arg.(type) {
		case []int64:
			normalizedArgs = append(normalizedArgs, pq.Array(v))
		default:
			normalizedArgs = append(normalizedArgs, arg)
		}
	}

	var builder strings.Builder
	index := 1
	for _, ch := range replaced {
		if ch == '?' {
			builder.WriteString(fmt.Sprintf("$%d", index))
			index++
			continue
		}
		builder.WriteRune(ch)
	}

	return builder.String(), normalizedArgs
}
