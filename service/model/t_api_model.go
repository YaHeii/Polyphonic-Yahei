package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TApiModel = (*customTApiModel)(nil)

type (
	// TApiModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTApiModel.
	TApiModel interface {
		tApiModel
		FindById(ctx context.Context, id int64) (*TApi, error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) ([]*TApi, error)
		FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TApi, int64, error)
		FindByRoleID(ctx context.Context, roleID int64) ([]*TApi, error)
		FindByUserID(ctx context.Context, userID string) ([]*TApi, error)
		Save(ctx context.Context, data *TApi) (int64, error)
		Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error)
		Clean(ctx context.Context) (int64, error)
	}

	customTApiModel struct {
		*defaultTApiModel
	}
)

// NewTApiModel returns a model for the database table.
func NewTApiModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TApiModel {
	return &customTApiModel{
		defaultTApiModel: newTApiModel(conn, c, opts...),
	}
}

func (m *customTApiModel) FindById(ctx context.Context, id int64) (*TApi, error) {
	return m.FindOne(ctx, id)
}

func (m *customTApiModel) FindALL(ctx context.Context, conditions string, args ...interface{}) ([]*TApi, error) {
	whereClause, bindArgs := buildPermissionWhereClause(conditions, args...)

	query := fmt.Sprintf("select %s from %s", tApiRows, m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}
	query += " order by id asc"

	var list []*TApi
	if err := m.QueryRowsNoCacheCtx(ctx, &list, query, bindArgs...); err != nil {
		return nil, err
	}

	return list, nil
}

func (m *customTApiModel) FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TApi, int64, error) {
	whereClause, bindArgs := buildPermissionWhereClause(conditions, args...)

	countQuery := fmt.Sprintf("select count(*) from %s", m.table)
	if whereClause != "" {
		countQuery += " where " + whereClause
	}
	var total int64
	if err := m.QueryRowNoCacheCtx(ctx, &total, countQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf("select %s from %s", tApiRows, m.table)
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

	var list []*TApi
	if err := m.QueryRowsNoCacheCtx(ctx, &list, listQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (m *customTApiModel) FindByRoleID(ctx context.Context, roleID int64) ([]*TApi, error) {
	query := fmt.Sprintf(`select a.%s
from %s a
join "public"."t_role_api" ra on ra.api_id = a.id
where ra.role_id = $1
order by a.id asc`, tApiRows, m.table)

	var list []*TApi
	if err := m.QueryRowsNoCacheCtx(ctx, &list, query, roleID); err != nil {
		return nil, err
	}

	return list, nil
}

func (m *customTApiModel) FindByUserID(ctx context.Context, userID string) ([]*TApi, error) {
	query := fmt.Sprintf(`select %s
from %s a
where a.id in (
	select distinct ra.api_id
	from "public"."t_role_api" ra
	join "public"."t_user_role" ur on ur.role_id = ra.role_id
	where ur.user_id = $1
)
order by a.id asc`, tApiRows, m.table)

	var list []*TApi
	if err := m.QueryRowsNoCacheCtx(ctx, &list, query, userID); err != nil {
		return nil, err
	}

	return list, nil
}

func (m *customTApiModel) Save(ctx context.Context, data *TApi) (int64, error) {
	if data.Id > 0 {
		query := fmt.Sprintf(`insert into %s (id, parent_id, name, path, method, traceable, status)
values ($1, $2, $3, $4, $5, $6, $7)
on conflict (id) do update set
	parent_id = excluded.parent_id,
	name = excluded.name,
	path = excluded.path,
	method = excluded.method,
	traceable = excluded.traceable,
	status = excluded.status,
	updated_at = now()
returning id`, m.table)
		if err := m.QueryRowNoCacheCtx(ctx, &data.Id, query, data.Id, data.ParentId, data.Name, data.Path, data.Method, data.Traceable, data.Status); err != nil {
			return 0, err
		}
		return data.Id, nil
	}

	query := fmt.Sprintf(`insert into %s (parent_id, name, path, method, traceable, status)
values ($1, $2, $3, $4, $5, $6)
on conflict (path, method, name) do update set
	parent_id = excluded.parent_id,
	traceable = excluded.traceable,
	status = excluded.status,
	updated_at = now()
returning id`, m.table)
	if err := m.QueryRowNoCacheCtx(ctx, &data.Id, query, data.ParentId, data.Name, data.Path, data.Method, data.Traceable, data.Status); err != nil {
		return 0, err
	}

	return data.Id, nil
}

func (m *customTApiModel) Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error) {
	whereClause, bindArgs := buildPermissionWhereClause(conditions, args...)

	query := fmt.Sprintf("delete from %s", m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}

	return rowsAffected(m.ExecNoCacheCtx(ctx, query, bindArgs...))
}

func (m *customTApiModel) Clean(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("delete from %s", m.table)
	return rowsAffected(m.ExecNoCacheCtx(ctx, query))
}
