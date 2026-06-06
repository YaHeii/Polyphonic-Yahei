package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TPhotoModel = (*customTPhotoModel)(nil)

type (
	// TPhotoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTPhotoModel.
	TPhotoModel interface {
		tPhotoModel
		FindById(ctx context.Context, id int64) (*TPhoto, error)
		FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TPhoto, int64, error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (int64, error)
		Save(ctx context.Context, data *TPhoto) (int64, error)
		Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error)
		Updates(ctx context.Context, columns map[string]interface{}, conditions string, args ...interface{}) (int64, error)
		CountByAlbumIDs(ctx context.Context, ids []int64) (map[int64]int64, error)
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

func (m *customTPhotoModel) FindById(ctx context.Context, id int64) (*TPhoto, error) {
	return m.FindOne(ctx, id)
}

func (m *customTPhotoModel) FindListAndTotal(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) ([]*TPhoto, int64, error) {
	whereClause, bindArgs := buildResourceWhereClause(conditions, args...)

	countQuery := fmt.Sprintf("select count(*) from %s", m.table)
	if whereClause != "" {
		countQuery += " where " + whereClause
	}
	var total int64
	if err := m.QueryRowNoCacheCtx(ctx, &total, countQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf("select %s from %s", tPhotoRows, m.table)
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

	var list []*TPhoto
	if err := m.QueryRowsNoCacheCtx(ctx, &list, listQuery, bindArgs...); err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (m *customTPhotoModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (int64, error) {
	whereClause, bindArgs := buildResourceWhereClause(conditions, args...)

	query := fmt.Sprintf("select count(*) from %s", m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}

	var total int64
	if err := m.QueryRowNoCacheCtx(ctx, &total, query, bindArgs...); err != nil {
		return 0, err
	}

	return total, nil
}

func (m *customTPhotoModel) Save(ctx context.Context, data *TPhoto) (int64, error) {
	if data.Id > 0 {
		query := fmt.Sprintf(`insert into %s (id, album_id, photo_name, photo_desc, photo_src, is_delete)
values ($1, $2, $3, $4, $5, $6)
on conflict (id) do update set
	album_id = excluded.album_id,
	photo_name = excluded.photo_name,
	photo_desc = excluded.photo_desc,
	photo_src = excluded.photo_src,
	is_delete = excluded.is_delete,
	updated_at = now()
returning id`, m.table)
		if err := m.QueryRowNoCacheCtx(ctx, &data.Id, query, data.Id, data.AlbumId, data.PhotoName, data.PhotoDesc, data.PhotoSrc, data.IsDelete); err != nil {
			return 0, err
		}
		return data.Id, nil
	}

	query := fmt.Sprintf(`insert into %s (album_id, photo_name, photo_desc, photo_src, is_delete)
values ($1, $2, $3, $4, $5)
returning id`, m.table)
	if err := m.QueryRowNoCacheCtx(ctx, &data.Id, query, data.AlbumId, data.PhotoName, data.PhotoDesc, data.PhotoSrc, data.IsDelete); err != nil {
		return 0, err
	}

	return data.Id, nil
}

func (m *customTPhotoModel) Deletes(ctx context.Context, conditions string, args ...interface{}) (int64, error) {
	whereClause, bindArgs := buildResourceWhereClause(conditions, args...)

	query := fmt.Sprintf("delete from %s", m.table)
	if whereClause != "" {
		query += " where " + whereClause
	}

	return rowsAffected(m.ExecNoCacheCtx(ctx, query, bindArgs...))
}

func (m *customTPhotoModel) Updates(ctx context.Context, columns map[string]interface{}, conditions string, args ...interface{}) (int64, error) {
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

func (m *customTPhotoModel) CountByAlbumIDs(ctx context.Context, ids []int64) (map[int64]int64, error) {
	if len(ids) == 0 {
		return map[int64]int64{}, nil
	}

	var rows []struct {
		AlbumID int64 `db:"album_id"`
		Total   int64 `db:"total"`
	}
	query := fmt.Sprintf("select album_id, count(*) as total from %s where album_id = any($1) and is_delete = false group by album_id", m.table)
	if err := m.QueryRowsNoCacheCtx(ctx, &rows, query, pq.Array(ids)); err != nil {
		return nil, err
	}

	result := make(map[int64]int64, len(rows))
	for _, row := range rows {
		result[row.AlbumID] = row.Total
	}

	return result, nil
}
