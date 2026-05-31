package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TWebsiteConfigModel = (*customTWebsiteConfigModel)(nil)

type (
	// TWebsiteConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTWebsiteConfigModel.
	TWebsiteConfigModel interface {
		tWebsiteConfigModel
	}

	customTWebsiteConfigModel struct {
		*defaultTWebsiteConfigModel
	}
)

// NewTWebsiteConfigModel returns a model for the database table.
func NewTWebsiteConfigModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TWebsiteConfigModel {
	return &customTWebsiteConfigModel{
		defaultTWebsiteConfigModel: newTWebsiteConfigModel(conn, c, opts...),
	}
}
