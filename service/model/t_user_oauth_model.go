package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TUserOauthModel = (*customTUserOauthModel)(nil)

type (
	// TUserOauthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTUserOauthModel.
	TUserOauthModel interface {
		tUserOauthModel
		FindByUserID(ctx context.Context, userID string) ([]*TUserOauth, error)
	}

	customTUserOauthModel struct {
		*defaultTUserOauthModel
	}
)

// NewTUserOauthModel returns a model for the database table.
func NewTUserOauthModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TUserOauthModel {
	return &customTUserOauthModel{
		defaultTUserOauthModel: newTUserOauthModel(conn, c, opts...),
	}
}

func (m *customTUserOauthModel) FindByUserID(ctx context.Context, userID string) ([]*TUserOauth, error) {
	query := "select " + tUserOauthRows + " from " + m.table + " where user_id = $1 order by id asc"

	var list []*TUserOauth
	if err := m.QueryRowsNoCacheCtx(ctx, &list, query, userID); err != nil {
		return nil, err
	}

	return list, nil
}
