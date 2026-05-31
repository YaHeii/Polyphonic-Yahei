package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TUserOauthModel = (*customTUserOauthModel)(nil)

type (
	// TUserOauthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTUserOauthModel.
	TUserOauthModel interface {
		tUserOauthModel
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
