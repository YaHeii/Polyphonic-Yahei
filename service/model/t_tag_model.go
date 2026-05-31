package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TTagModel = (*customTTagModel)(nil)

type (
	// TTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTTagModel.
	TTagModel interface {
		tTagModel
	}

	customTTagModel struct {
		*defaultTTagModel
	}
)

// NewTTagModel returns a model for the database table.
func NewTTagModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TTagModel {
	return &customTTagModel{
		defaultTTagModel: newTTagModel(conn, c, opts...),
	}
}
