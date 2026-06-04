package model

import (
	"context"
	"fmt"

	"github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TTagModel = (*customTTagModel)(nil)

type (
	// TTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTTagModel.
	TTagModel interface {
		tTagModel
		FindById(ctx context.Context, id int64) (*TTag, error)
		FindByNames(ctx context.Context, names []string) ([]*TTag, error)
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

func (m *customTTagModel) FindById(ctx context.Context, id int64) (*TTag, error) {
	return m.FindOne(ctx, id)
}

func (m *customTTagModel) FindByNames(ctx context.Context, names []string) ([]*TTag, error) {
	if len(names) == 0 {
		return []*TTag{}, nil
	}

	query := fmt.Sprintf("select %s from %s where tag_name = any($1)", tTagRows, m.table)
	var list []*TTag
	if err := m.QueryRowsNoCacheCtx(ctx, &list, query, pq.Array(names)); err != nil {
		return nil, err
	}

	return list, nil
}
