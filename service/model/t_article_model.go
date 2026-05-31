package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TArticleModel = (*customTArticleModel)(nil)

type (
	// TArticleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTArticleModel.
	TArticleModel interface {
		tArticleModel
	}

	customTArticleModel struct {
		*defaultTArticleModel
	}
)

// NewTArticleModel returns a model for the database table.
func NewTArticleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TArticleModel {
	return &customTArticleModel{
		defaultTArticleModel: newTArticleModel(conn, c, opts...),
	}
}
