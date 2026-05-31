package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TMessageModel = (*customTMessageModel)(nil)

type (
	// TMessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTMessageModel.
	TMessageModel interface {
		tMessageModel
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
