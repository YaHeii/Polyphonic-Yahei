package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TMenuModel = (*customTMenuModel)(nil)

type (
	// TMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTMenuModel.
	TMenuModel interface {
		tMenuModel
	}

	customTMenuModel struct {
		*defaultTMenuModel
	}
)

// NewTMenuModel returns a model for the database table.
func NewTMenuModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TMenuModel {
	return &customTMenuModel{
		defaultTMenuModel: newTMenuModel(conn, c, opts...),
	}
}
