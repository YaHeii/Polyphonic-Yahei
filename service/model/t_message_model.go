package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TMessageModel = (*customTMessageModel)(nil)

type (
	// TMessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTMessageModel.
	TMessageModel interface {
		tMessageModel
		FindById(ctx context.Context, id int64) (*TMessage, error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (int64, error)
		FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*TMessage, total int64, err error)
		Deletes(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		Updates(ctx context.Context, columns map[string]interface{}, conditions string, args ...interface{}) (rows int64, err error)
		Save(ctx context.Context, data *TMessage) (rows int64, err error)
	}

	customTMessageModel struct {
		*defaultTMessageModel
	}
)

// NewTMessageModel returns a model for the database table.
func NewTMessageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TMessageModel {
	return &customTMessageModel{
		defaultTMessageModel: newTMessageModel(conn, c, opts...),
	}
}

func (m *customTMessageModel) FindById(ctx context.Context, id int64) (*TMessage, error) {
	return m.FindOne(ctx, id)
}

func (m *customTMessageModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (int64, error) {
	whereClause, bindArgs := buildMessageWhereClause(conditions, args...)

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

func (m *customTMessageModel) FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*TMessage, total int64, err error) {
	whereClause, bindArgs := buildMessageWhereClause(conditions, args...)

	countQuery := fmt.Sprintf("select count(*) from %s", m.table)
	if whereClause != "" {
		countQuery += " where " + whereClause
	}
	if err = m.QueryRowNoCacheCtx(ctx, &total, countQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf("select %s from %s", tMessageRows, m.table)
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

func (m *customTMessageModel) Deletes(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	whereClause, bindArgs := buildMessageWhereClause(conditions, args...)

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

func (m *customTMessageModel) Updates(ctx context.Context, columns map[string]interface{}, conditions string, args ...interface{}) (rows int64, err error) {
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

	whereClause, whereArgs := buildMessageWhereClauseWithStartIndex(conditions, index, args...)
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

func (m *customTMessageModel) Save(ctx context.Context, data *TMessage) (rows int64, err error) {
	if err := m.Update(ctx, data); err != nil {
		return 0, err
	}
	return 1, nil
}

func buildMessageWhereClause(conditions string, args ...interface{}) (string, []interface{}) {
	return buildMessageWhereClauseWithStartIndex(conditions, 1, args...)
}

func buildMessageWhereClauseWithStartIndex(conditions string, start int, args ...interface{}) (string, []interface{}) {
	if conditions == "" {
		return "", args
	}

	replaced := strings.ReplaceAll(conditions, "id in (?)", "id = any(?)")
	var normalizedArgs []interface{}
	for _, arg := range args {
		switch v := arg.(type) {
		case []int64:
			normalizedArgs = append(normalizedArgs, v)
		default:
			normalizedArgs = append(normalizedArgs, arg)
		}
	}

	var builder strings.Builder
	index := start
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
