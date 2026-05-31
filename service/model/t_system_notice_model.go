package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TSystemNoticeModel = (*customTSystemNoticeModel)(nil)

type (
	// TSystemNoticeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTSystemNoticeModel.
	TSystemNoticeModel interface {
		tSystemNoticeModel
	}

	customTSystemNoticeModel struct {
		*defaultTSystemNoticeModel
	}
)

// NewTSystemNoticeModel returns a model for the database table.
func NewTSystemNoticeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TSystemNoticeModel {
	return &customTSystemNoticeModel{
		defaultTSystemNoticeModel: newTSystemNoticeModel(conn, c, opts...),
	}
}
