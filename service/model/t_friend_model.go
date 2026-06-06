package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TFriendModel = (*customTFriendModel)(nil)

type (
	// TFriendModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTFriendModel.
	TFriendModel interface {
		tFriendModel
		FindById(ctx context.Context, id int64) (*TFriend, error)
		FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TFriend, int64, error)
		Save(ctx context.Context, data *TFriend) (int64, error)
		Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error)
	}

	customTFriendModel struct {
		*defaultTFriendModel
	}
)

// NewTFriendModel returns a model for the database table.
func NewTFriendModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TFriendModel {
	return &customTFriendModel{
		defaultTFriendModel: newTFriendModel(conn, c, opts...),
	}
}

func (m *customTFriendModel) FindById(ctx context.Context, id int64) (*TFriend, error) {
	return m.FindOne(ctx, id)
}

func (m *customTFriendModel) FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TFriend, int64, error) {
	whereClause, bindArgs := buildSocialWhereClause(conditions, args...)

	countQuery := fmt.Sprintf("select count(*) from %s", m.table)
	if whereClause != "" {
		countQuery += " where " + whereClause
	}
	var total int64
	if err := m.QueryRowNoCacheCtx(ctx, &total, countQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf("select %s from %s", tFriendRows, m.table)
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

	var list []*TFriend
	if err := m.QueryRowsNoCacheCtx(ctx, &list, listQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (m *customTFriendModel) Save(ctx context.Context, data *TFriend) (int64, error) {
	if data.Id > 0 {
		query := fmt.Sprintf(`insert into %s (id, link_name, link_avatar, link_address, link_intro)
values ($1, $2, $3, $4, $5)
on conflict (id) do update set
	link_name = excluded.link_name,
	link_avatar = excluded.link_avatar,
	link_address = excluded.link_address,
	link_intro = excluded.link_intro,
	updated_at = now()
returning id`, m.table)
		if err := m.QueryRowNoCacheCtx(ctx, &data.Id, query, data.Id, data.LinkName, data.LinkAvatar, data.LinkAddress, data.LinkIntro); err != nil {
			return 0, err
		}
		return data.Id, nil
	}

	query := fmt.Sprintf(`insert into %s (link_name, link_avatar, link_address, link_intro)
values ($1, $2, $3, $4)
on conflict (link_name) do update set
	link_avatar = excluded.link_avatar,
	link_address = excluded.link_address,
	link_intro = excluded.link_intro,
	updated_at = now()
returning id`, m.table)
	if err := m.QueryRowNoCacheCtx(ctx, &data.Id, query, data.LinkName, data.LinkAvatar, data.LinkAddress, data.LinkIntro); err != nil {
		return 0, err
	}

	return data.Id, nil
}

func (m *customTFriendModel) Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error) {
	whereClause, bindArgs := buildSocialWhereClause(conditions, args...)

	query := fmt.Sprintf("delete from %s", m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}

	return rowsAffected(m.ExecNoCacheCtx(ctx, query, bindArgs...))
}
