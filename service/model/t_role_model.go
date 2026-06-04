package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TRoleModel = (*customTRoleModel)(nil)

type (
	// TRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTRoleModel.
	TRoleModel interface {
		tRoleModel
		FindDefaultRoles(ctx context.Context) ([]*TRole, error)
		FindRolesByUserID(ctx context.Context, userID string) ([]*TRole, error)
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

func (m *customTRoleModel) FindDefaultRoles(ctx context.Context) ([]*TRole, error) {
	query := "select " + tRoleRows + " from " + m.table + " where is_default = true and status = 0 order by id asc"

	var roles []*TRole
	if err := m.QueryRowsNoCacheCtx(ctx, &roles, query); err != nil {
		return nil, err
	}

	return roles, nil
}

func (m *customTRoleModel) FindRolesByUserID(ctx context.Context, userID string) ([]*TRole, error) {
	query := `
select r.id, r.parent_id, r.role_key, r.role_label, r.role_comment, r.is_default, r.status, r.created_at, r.updated_at
from "public"."t_user_role" ur
join ` + m.table + ` r on r.id = ur.role_id
where ur.user_id = $1
order by r.id asc`

	var roles []*TRole
	if err := m.QueryRowsNoCacheCtx(ctx, &roles, query, userID); err != nil {
		return nil, err
	}

	return roles, nil
}
