package model

import (
	"context"
	"fmt"

	"github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TCategoryModel = (*customTCategoryModel)(nil)

type (
	// TCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTCategoryModel.
	TCategoryModel interface {
		tCategoryModel
		FindById(ctx context.Context, id int64) (*TCategory, error)
		FindByIds(ctx context.Context, ids []int64) ([]*TCategory, error)
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

func (m *customTCategoryModel) FindById(ctx context.Context, id int64) (*TCategory, error) {
	return m.FindOne(ctx, id)
}

func (m *customTCategoryModel) FindByIds(ctx context.Context, ids []int64) ([]*TCategory, error) {
	if len(ids) == 0 {
		return []*TCategory{}, nil
	}

	query := fmt.Sprintf("select %s from %s where id = any($1)", tCategoryRows, m.table)
	var list []*TCategory
	if err := m.QueryRowsNoCacheCtx(ctx, &list, query, pq.Array(ids)); err != nil {
		return nil, err
	}

	return list, nil
}
