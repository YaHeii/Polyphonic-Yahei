package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TRoleModel = (*customTRoleModel)(nil)

type (
	// TRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTRoleModel.
	TRoleModel interface {
		tRoleModel
		FindById(ctx context.Context, id int64) (*TRole, error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) ([]*TRole, error)
		FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TRole, int64, error)
		FindDefaultRoles(ctx context.Context) ([]*TRole, error)
		FindRolesByUserID(ctx context.Context, userID string) ([]*TRole, error)
		Save(ctx context.Context, data *TRole) (int64, error)
		Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error)
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

func (m *customTRoleModel) FindById(ctx context.Context, id int64) (*TRole, error) {
	return m.FindOne(ctx, id)
}

func (m *customTRoleModel) FindALL(ctx context.Context, conditions string, args ...interface{}) ([]*TRole, error) {
	whereClause, bindArgs := buildPermissionWhereClause(conditions, args...)

	query := fmt.Sprintf("select %s from %s", tRoleRows, m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}
	query += " order by id asc"

	var list []*TRole
	if err := m.QueryRowsNoCacheCtx(ctx, &list, query, bindArgs...); err != nil {
		return nil, err
	}

	return list, nil
}

func (m *customTRoleModel) FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TRole, int64, error) {
	whereClause, bindArgs := buildPermissionWhereClause(conditions, args...)

	countQuery := fmt.Sprintf("select count(*) from %s", m.table)
	if whereClause != "" {
		countQuery += " where " + whereClause
	}
	var total int64
	if err := m.QueryRowNoCacheCtx(ctx, &total, countQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf("select %s from %s", tRoleRows, m.table)
	if whereClause != "" {
		listQuery += " where " + whereClause
	}
	if sorts != "" {
		listQuery += " order by " + sorts
	}
	if page > 0 && size > 0 {
		offset := (page - 1) * size
		listQuery += fmt.Sprintf(" limit $%d offset $%d", len(bindArgs)+1, len(bindArgs)+2)
		bindArgs = append(bindArgs, size, offset)
	}

	var list []*TRole
	if err := m.QueryRowsNoCacheCtx(ctx, &list, listQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (m *customTRoleModel) FindDefaultRoles(ctx context.Context) ([]*TRole, error) {
	query := "select " + tRoleRows + " from " + m.table + " where is_default = true and status = 0 order by id asc"

	var roles []*TRole
	if err := m.QueryRowsNoCacheCtx(ctx, &roles, query); err != nil {
		return nil, err
	}

	return roles, nil
}

func (m *customTRoleModel) Save(ctx context.Context, data *TRole) (int64, error) {
	if data.Id > 0 {
		query := fmt.Sprintf(`insert into %s (id, parent_id, role_key, role_label, role_comment, is_default, status)
values ($1, $2, $3, $4, $5, $6, $7)
on conflict (id) do update set
	parent_id = excluded.parent_id,
	role_key = excluded.role_key,
	role_label = excluded.role_label,
	role_comment = excluded.role_comment,
	is_default = excluded.is_default,
	status = excluded.status,
	updated_at = now()
returning id`, m.table)
		if err := m.QueryRowNoCacheCtx(ctx, &data.Id, query, data.Id, data.ParentId, data.RoleKey, data.RoleLabel, data.RoleComment, data.IsDefault, data.Status); err != nil {
			return 0, err
		}
		return data.Id, nil
	}

	query := fmt.Sprintf(`insert into %s (parent_id, role_key, role_label, role_comment, is_default, status)
values ($1, $2, $3, $4, $5, $6)
returning id`, m.table)
	if err := m.QueryRowNoCacheCtx(ctx, &data.Id, query, data.ParentId, data.RoleKey, data.RoleLabel, data.RoleComment, data.IsDefault, data.Status); err != nil {
		return 0, err
	}

	return data.Id, nil
}

func (m *customTRoleModel) Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error) {
	whereClause, bindArgs := buildPermissionWhereClause(conditions, args...)

	query := fmt.Sprintf("delete from %s", m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}

	return rowsAffected(m.ExecNoCacheCtx(ctx, query, bindArgs...))
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
