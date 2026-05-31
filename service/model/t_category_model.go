package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TCategoryModel = (*customTCategoryModel)(nil)

type (
	// TCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTCategoryModel.
	TCategoryModel interface {
		tCategoryModel
	}

	customTCategoryModel struct {
		*defaultTCategoryModel
	}
)

// NewTCategoryModel returns a model for the database table.
func NewTCategoryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TCategoryModel {
	return &customTCategoryModel{
		defaultTCategoryModel: newTCategoryModel(conn, c, opts...),
	}
}
