package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TAlbumModel = (*customTAlbumModel)(nil)

type (
	// TAlbumModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTAlbumModel.
	TAlbumModel interface {
		tAlbumModel
		FindById(ctx context.Context, id int64) (*TAlbum, error)
		FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TAlbum, int64, error)
		Save(ctx context.Context, data *TAlbum) (int64, error)
		Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error)
		Updates(ctx context.Context, columns map[string]interface{}, conditions string, args ...interface{}) (int64, error)
	}

	customTAlbumModel struct {
		*defaultTAlbumModel
	}
)

// NewTAlbumModel returns a model for the database table.
func NewTAlbumModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TAlbumModel {
	return &customTAlbumModel{
		defaultTAlbumModel: newTAlbumModel(conn, c, opts...),
	}
}

func (m *customTAlbumModel) FindById(ctx context.Context, id int64) (*TAlbum, error) {
	return m.FindOne(ctx, id)
}

func (m *customTAlbumModel) FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TAlbum, int64, error) {
	whereClause, bindArgs := buildResourceWhereClause(conditions, args...)

	countQuery := fmt.Sprintf("select count(*) from %s", m.table)
	if whereClause != "" {
		countQuery += " where " + whereClause
	}
	var total int64
	if err := m.QueryRowNoCacheCtx(ctx, &total, countQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf("select %s from %s", tAlbumRows, m.table)
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

	var list []*TAlbum
	if err := m.QueryRowsNoCacheCtx(ctx, &list, listQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (m *customTAlbumModel) Save(ctx context.Context, data *TAlbum) (int64, error) {
	if data.Id > 0 {
		query := fmt.Sprintf(`insert into %s (id, album_name, album_desc, album_cover, is_delete, status)
values ($1, $2, $3, $4, $5, $6)
on conflict (id) do update set
	album_name = excluded.album_name,
	album_desc = excluded.album_desc,
	album_cover = excluded.album_cover,
	is_delete = excluded.is_delete,
	status = excluded.status,
	updated_at = now()
returning id`, m.table)
		if err := m.QueryRowNoCacheCtx(ctx, &data.Id, query, data.Id, data.AlbumName, data.AlbumDesc, data.AlbumCover, data.IsDelete, data.Status); err != nil {
			return 0, err
		}
		return data.Id, nil
	}

	query := fmt.Sprintf(`insert into %s (album_name, album_desc, album_cover, is_delete, status)
values ($1, $2, $3, $4, $5)
returning id`, m.table)
	if err := m.QueryRowNoCacheCtx(ctx, &data.Id, query, data.AlbumName, data.AlbumDesc, data.AlbumCover, data.IsDelete, data.Status); err != nil {
		return 0, err
	}

	return data.Id, nil
}

func (m *customTAlbumModel) Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error) {
	whereClause, bindArgs := buildResourceWhereClause(conditions, args...)

	query := fmt.Sprintf("delete from %s", m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}

	return rowsAffected(m.ExecNoCacheCtx(ctx, query, bindArgs...))
}

func (m *customTAlbumModel) Updates(ctx context.Context, columns map[string]interface{}, conditions string, args ...interface{}) (int64, error) {
	if len(columns) == 0 {
		return 0, nil
	}

	setClauses := make([]string, 0, len(columns))
	bindArgs := make([]interface{}, 0, len(columns)+len(args))
	index := 1
	for column, value := range columns {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", column, index))
		bindArgs = append(bindArgs, value)
		index++
	}

	whereClause, whereArgs := buildResourceWhereClauseWithStartIndex(conditions, index, args...)
	bindArgs = append(bindArgs, whereArgs...)

	query := fmt.Sprintf("update %s set %s", m.table, strings.Join(setClauses, ", "))
	if whereClause != "" {
		query += " where " + whereClause
	}

	return rowsAffected(m.ExecNoCacheCtx(ctx, query, bindArgs...))
}
