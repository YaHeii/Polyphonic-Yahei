package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TPageModel = (*customTPageModel)(nil)

type (
	// TPageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTPageModel.
	TPageModel interface {
		tPageModel
	}

	customTPageModel struct {
		*defaultTPageModel
	}
)

// NewTPageModel returns a model for the database table.
func NewTPageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TPageModel {
	return &customTPageModel{
		defaultTPageModel: newTPageModel(conn, c, opts...),
	}
}
