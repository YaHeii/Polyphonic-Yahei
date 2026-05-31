package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TFriendModel = (*customTFriendModel)(nil)

type (
	// TFriendModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTFriendModel.
	TFriendModel interface {
		tFriendModel
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
