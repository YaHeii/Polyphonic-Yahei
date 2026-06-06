package model

import (
	"context"
	"fmt"

	"github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TTalkModel = (*customTTalkModel)(nil)

type (
	// TTalkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTTalkModel.
	TTalkModel interface {
		tTalkModel
		FindById(ctx context.Context, id int64) (*TTalk, error)
		FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TTalk, int64, error)
		Save(ctx context.Context, data *TTalk) (int64, error)
		Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error)
	}

	customTTalkModel struct {
		*defaultTTalkModel
	}
)

// NewTTalkModel returns a model for the database table.
func NewTTalkModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TTalkModel {
	return &customTTalkModel{
		defaultTTalkModel: newTTalkModel(conn, c, opts...),
	}
}

func (m *customTTalkModel) FindById(ctx context.Context, id int64) (*TTalk, error) {
	return m.FindOne(ctx, id)
}

func (m *customTTalkModel) FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TTalk, int64, error) {
	whereClause, bindArgs := buildSocialWhereClause(conditions, args...)

	countQuery := fmt.Sprintf("select count(*) from %s", m.table)
	if whereClause != "" {
		countQuery += " where " + whereClause
	}
	var total int64
	if err := m.QueryRowNoCacheCtx(ctx, &total, countQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf("select %s from %s", tTalkRows, m.table)
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

	var list []*TTalk
	if err := m.QueryRowsNoCacheCtx(ctx, &list, listQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (m *customTTalkModel) Save(ctx context.Context, data *TTalk) (int64, error) {
	if data.Images == nil {
		data.Images = pq.StringArray{}
	}

	if data.Id > 0 {
		query := fmt.Sprintf(`insert into %s (id, user_id, content, images, is_top, status, like_count)
values ($1, $2, $3, $4, $5, $6, $7)
on conflict (id) do update set
	user_id = excluded.user_id,
	content = excluded.content,
	images = excluded.images,
	is_top = excluded.is_top,
	status = excluded.status,
	like_count = excluded.like_count,
	updated_at = now()
returning id`, m.table)
		if err := m.QueryRowNoCacheCtx(ctx, &data.Id, query, data.Id, data.UserId, data.Content, data.Images, data.IsTop, data.Status, data.LikeCount); err != nil {
			return 0, err
		}
		return data.Id, nil
	}

	query := fmt.Sprintf(`insert into %s (user_id, content, images, is_top, status, like_count)
values ($1, $2, $3, $4, $5, $6)
returning id`, m.table)
	if err := m.QueryRowNoCacheCtx(ctx, &data.Id, query, data.UserId, data.Content, data.Images, data.IsTop, data.Status, data.LikeCount); err != nil {
		return 0, err
	}

	return data.Id, nil
}

func (m *customTTalkModel) Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error) {
	whereClause, bindArgs := buildSocialWhereClause(conditions, args...)

	query := fmt.Sprintf("delete from %s", m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}

	return rowsAffected(m.ExecNoCacheCtx(ctx, query, bindArgs...))
}
