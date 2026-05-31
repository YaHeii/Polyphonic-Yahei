package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TTalkModel = (*customTTalkModel)(nil)

type (
	// TTalkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTTalkModel.
	TTalkModel interface {
		tTalkModel
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
