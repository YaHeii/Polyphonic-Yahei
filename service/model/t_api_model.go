package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TApiModel = (*customTApiModel)(nil)

type (
	// TApiModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTApiModel.
	TApiModel interface {
		tApiModel
	}

	customTApiModel struct {
		*defaultTApiModel
	}
)

// NewTApiModel returns a model for the database table.
func NewTApiModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TApiModel {
	return &customTApiModel{
		defaultTApiModel: newTApiModel(conn, c, opts...),
	}
}
