package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TRoleModel = (*customTRoleModel)(nil)

type (
	// TRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTRoleModel.
	TRoleModel interface {
		tRoleModel
	}

	customTRoleModel struct {
		*defaultTRoleModel
	}
)

// NewTRoleModel returns a model for the database table.
func NewTRoleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TRoleModel {
	return &customTRoleModel{
		defaultTRoleModel: newTRoleModel(conn, c, opts...),
	}
}
