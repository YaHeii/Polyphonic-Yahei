package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TMenuModel = (*customTMenuModel)(nil)

type (
	// TMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTMenuModel.
	TMenuModel interface {
		tMenuModel
		FindById(ctx context.Context, id int64) (*TMenu, error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) ([]*TMenu, error)
		FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TMenu, int64, error)
		FindByRoleID(ctx context.Context, roleID int64) ([]*TMenu, error)
		FindByUserID(ctx context.Context, userID string) ([]*TMenu, error)
		Save(ctx context.Context, data *TMenu) (int64, error)
		Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error)
		Clean(ctx context.Context) (int64, error)
	}

	customTMenuModel struct {
		*defaultTMenuModel
	}
)

// NewTMenuModel returns a model for the database table.
func NewTMenuModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TMenuModel {
	return &customTMenuModel{
		defaultTMenuModel: newTMenuModel(conn, c, opts...),
	}
}

func (m *customTMenuModel) FindById(ctx context.Context, id int64) (*TMenu, error) {
	return m.FindOne(ctx, id)
}

func (m *customTMenuModel) FindALL(ctx context.Context, conditions string, args ...interface{}) ([]*TMenu, error) {
	whereClause, bindArgs := buildPermissionWhereClause(conditions, args...)

	query := fmt.Sprintf("select %s from %s", tMenuRows, m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}
	query += " order by rank asc, id asc"

	var list []*TMenu
	if err := m.QueryRowsNoCacheCtx(ctx, &list, query, bindArgs...); err != nil {
		return nil, err
	}

	return list, nil
}

func (m *customTMenuModel) FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TMenu, int64, error) {
	whereClause, bindArgs := buildPermissionWhereClause(conditions, args...)

	countQuery := fmt.Sprintf("select count(*) from %s", m.table)
	if whereClause != "" {
		countQuery += " where " + whereClause
	}
	var total int64
	if err := m.QueryRowNoCacheCtx(ctx, &total, countQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf("select %s from %s", tMenuRows, m.table)
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

	var list []*TMenu
	if err := m.QueryRowsNoCacheCtx(ctx, &list, listQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (m *customTMenuModel) FindByRoleID(ctx context.Context, roleID int64) ([]*TMenu, error) {
	query := fmt.Sprintf(`select m.%s
from %s m
join "public"."t_role_menu" rm on rm.menu_id = m.id
where rm.role_id = $1
order by m.rank asc, m.id asc`, tMenuRows, m.table)

	var list []*TMenu
	if err := m.QueryRowsNoCacheCtx(ctx, &list, query, roleID); err != nil {
		return nil, err
	}

	return list, nil
}

func (m *customTMenuModel) FindByUserID(ctx context.Context, userID string) ([]*TMenu, error) {
	query := fmt.Sprintf(`select %s
from %s m
where m.id in (
	select distinct rm.menu_id
	from "public"."t_role_menu" rm
	join "public"."t_user_role" ur on ur.role_id = rm.role_id
	where ur.user_id = $1
)
order by m.rank asc, m.id asc`, tMenuRows, m.table)

	var list []*TMenu
	if err := m.QueryRowsNoCacheCtx(ctx, &list, query, userID); err != nil {
		return nil, err
	}

	return list, nil
}

func (m *customTMenuModel) Save(ctx context.Context, data *TMenu) (int64, error) {
	if data.Extra == "" {
		data.Extra = "{}"
	}

	if data.Id > 0 {
		query := fmt.Sprintf(`insert into %s (id, parent_id, path, name, component, redirect, type, title, icon, rank, perm, params, keep_alive, always_show, visible, status, extra)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
on conflict (id) do update set
	parent_id = excluded.parent_id,
	path = excluded.path,
	name = excluded.name,
	component = excluded.component,
	redirect = excluded.redirect,
	type = excluded.type,
	title = excluded.title,
	icon = excluded.icon,
	rank = excluded.rank,
	perm = excluded.perm,
	params = excluded.params,
	keep_alive = excluded.keep_alive,
	always_show = excluded.always_show,
	visible = excluded.visible,
	status = excluded.status,
	extra = excluded.extra,
	updated_at = now()
returning id`, m.table)
		if err := m.QueryRowNoCacheCtx(ctx, &data.Id, query, data.Id, data.ParentId, data.Path, data.Name, data.Component, data.Redirect, data.Type, data.Title, data.Icon, data.Rank, data.Perm, data.Params, data.KeepAlive, data.AlwaysShow, data.Visible, data.Status, data.Extra); err != nil {
			return 0, err
		}
		return data.Id, nil
	}

	query := fmt.Sprintf(`insert into %s (parent_id, path, name, component, redirect, type, title, icon, rank, perm, params, keep_alive, always_show, visible, status, extra)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
on conflict (path, perm) do update set
	parent_id = excluded.parent_id,
	name = excluded.name,
	component = excluded.component,
	redirect = excluded.redirect,
	type = excluded.type,
	title = excluded.title,
	icon = excluded.icon,
	rank = excluded.rank,
	params = excluded.params,
	keep_alive = excluded.keep_alive,
	always_show = excluded.always_show,
	visible = excluded.visible,
	status = excluded.status,
	extra = excluded.extra,
	updated_at = now()
returning id`, m.table)
	if err := m.QueryRowNoCacheCtx(ctx, &data.Id, query, data.ParentId, data.Path, data.Name, data.Component, data.Redirect, data.Type, data.Title, data.Icon, data.Rank, data.Perm, data.Params, data.KeepAlive, data.AlwaysShow, data.Visible, data.Status, data.Extra); err != nil {
		return 0, err
	}

	return data.Id, nil
}

func (m *customTMenuModel) Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error) {
	whereClause, bindArgs := buildPermissionWhereClause(conditions, args...)

	query := fmt.Sprintf("delete from %s", m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}

	return rowsAffected(m.ExecNoCacheCtx(ctx, query, bindArgs...))
}

func (m *customTMenuModel) Clean(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("delete from %s", m.table)
	return rowsAffected(m.ExecNoCacheCtx(ctx, query))
}
