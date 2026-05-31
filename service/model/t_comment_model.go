package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TCommentModel = (*customTCommentModel)(nil)

type (
	// TCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTCommentModel.
	TCommentModel interface {
		tCommentModel
	}

	customTCommentModel struct {
		*defaultTCommentModel
	}
)

// NewTCommentModel returns a model for the database table.
func NewTCommentModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TCommentModel {
	return &customTCommentModel{
		defaultTCommentModel: newTCommentModel(conn, c, opts...),
	}
}
