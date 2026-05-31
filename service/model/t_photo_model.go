package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TPhotoModel = (*customTPhotoModel)(nil)

type (
	// TPhotoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTPhotoModel.
	TPhotoModel interface {
		tPhotoModel
	}

	customTPhotoModel struct {
		*defaultTPhotoModel
	}
)

// NewTPhotoModel returns a model for the database table.
func NewTPhotoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TPhotoModel {
	return &customTPhotoModel{
		defaultTPhotoModel: newTPhotoModel(conn, c, opts...),
	}
}
