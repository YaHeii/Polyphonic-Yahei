package resourcerpclogic

import (
	"encoding/json"
	"strings"

	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/common/query"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/resourcerpc"
)

func convertAddPhotoIn(in *resourcerpc.AddPhotoReq) *model.TPhoto {
	return &model.TPhoto{
		Id:        in.Id,
		AlbumId:   in.AlbumId,
		PhotoName: in.PhotoName,
		PhotoDesc: in.PhotoDesc,
		PhotoSrc:  in.PhotoSrc,
		IsDelete:  in.IsDelete,
	}
}

func convertUpdatePhotoIn(in *resourcerpc.UpdatePhotoReq) *model.TPhoto {
	return &model.TPhoto{
		Id:        in.Id,
		AlbumId:   in.AlbumId,
		PhotoName: in.PhotoName,
		PhotoDesc: in.PhotoDesc,
		PhotoSrc:  in.PhotoSrc,
		IsDelete:  in.IsDelete,
	}
}

func convertPhotoOut(record *model.TPhoto) *resourcerpc.Photo {
	if record == nil {
		return nil
	}

	return &resourcerpc.Photo{
		Id:        record.Id,
		AlbumId:   record.AlbumId,
		PhotoName: record.PhotoName,
		PhotoDesc: record.PhotoDesc,
		PhotoSrc:  record.PhotoSrc,
		IsDelete:  record.IsDelete,
		CreatedAt: record.CreatedAt.UnixMilli(),
		UpdatedAt: record.UpdatedAt.UnixMilli(),
	}
}

func convertPhotoListOut(records []*model.TPhoto) []*resourcerpc.Photo {
	list := make([]*resourcerpc.Photo, 0, len(records))
	for _, record := range records {
		list = append(list, convertPhotoOut(record))
	}
	return list
}

func convertAddAlbumIn(in *resourcerpc.AddAlbumReq) *model.TAlbum {
	return &model.TAlbum{
		Id:         in.Id,
		AlbumName:  in.AlbumName,
		AlbumDesc:  in.AlbumDesc,
		AlbumCover: in.AlbumCover,
		IsDelete:   in.IsDelete,
		Status:     in.Status,
	}
}

func convertUpdateAlbumIn(in *resourcerpc.UpdateAlbumReq) *model.TAlbum {
	return &model.TAlbum{
		Id:         in.Id,
		AlbumName:  in.AlbumName,
		AlbumDesc:  in.AlbumDesc,
		AlbumCover: in.AlbumCover,
		IsDelete:   in.IsDelete,
		Status:     in.Status,
	}
}

func convertAlbumOut(record *model.TAlbum, photoCount int64) *resourcerpc.Album {
	if record == nil {
		return nil
	}

	return &resourcerpc.Album{
		Id:         record.Id,
		AlbumName:  record.AlbumName,
		AlbumDesc:  record.AlbumDesc,
		AlbumCover: record.AlbumCover,
		IsDelete:   record.IsDelete,
		Status:     record.Status,
		CreatedAt:  record.CreatedAt.UnixMilli(),
		UpdatedAt:  record.UpdatedAt.UnixMilli(),
		PhotoCount: photoCount,
	}
}

func convertAlbumListOut(records []*model.TAlbum, photoCounts map[int64]int64) []*resourcerpc.Album {
	list := make([]*resourcerpc.Album, 0, len(records))
	for _, record := range records {
		list = append(list, convertAlbumOut(record, photoCounts[record.Id]))
	}
	return list
}

func convertAddPageIn(in *resourcerpc.AddPageReq) *model.TPage {
	return &model.TPage{
		Id:             in.Id,
		PageName:       in.PageName,
		PageLabel:      in.PageLabel,
		PageCover:      in.PageCover,
		IsCarousel:     in.IsCarousel,
		CarouselCovers: encodeCarouselCovers(in.CarouselCovers),
	}
}

func convertUpdatePageIn(in *resourcerpc.UpdatePageReq) *model.TPage {
	return &model.TPage{
		Id:             in.Id,
		PageName:       in.PageName,
		PageLabel:      in.PageLabel,
		PageCover:      in.PageCover,
		IsCarousel:     in.IsCarousel,
		CarouselCovers: encodeCarouselCovers(in.CarouselCovers),
	}
}

func convertPageOut(record *model.TPage) *resourcerpc.Page {
	if record == nil {
		return nil
	}

	return &resourcerpc.Page{
		Id:             record.Id,
		PageName:       record.PageName,
		PageLabel:      record.PageLabel,
		PageCover:      record.PageCover,
		IsCarousel:     record.IsCarousel,
		CarouselCovers: decodeCarouselCovers(record.CarouselCovers),
		CreatedAt:      record.CreatedAt.UnixMilli(),
		UpdatedAt:      record.UpdatedAt.UnixMilli(),
	}
}

func convertPageListOut(records []*model.TPage) []*resourcerpc.Page {
	list := make([]*resourcerpc.Page, 0, len(records))
	for _, record := range records {
		list = append(list, convertPageOut(record))
	}
	return list
}

func encodeCarouselCovers(values []string) string {
	if len(values) == 0 {
		return ""
	}
	data, err := json.Marshal(values)
	if err != nil {
		return ""
	}
	return string(data)
}

func decodeCarouselCovers(value string) []string {
	if value == "" {
		return []string{}
	}

	var values []string
	if err := json.Unmarshal([]byte(value), &values); err == nil {
		return values
	}

	parts := strings.Split(value, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		if trimmed := strings.TrimSpace(part); trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func buildPhotoQuery(in *resourcerpc.FindPhotoListReq) (int, int, string, string, []any) {
	var opts []query.Option
	if in.Paginate != nil {
		opts = append(opts, query.WithPage(int(in.Paginate.Page)))
		opts = append(opts, query.WithSize(int(in.Paginate.PageSize)))
		opts = append(opts, query.WithSorts(in.Paginate.Sorts...))
	}
	if in.AlbumId != 0 {
		opts = append(opts, query.WithCondition("album_id = ?", in.AlbumId))
	}
	if in.IsDelete != nil {
		opts = append(opts, query.WithCondition("is_delete = ?", *in.IsDelete))
	}
	return query.NewQueryBuilder(opts...).Build()
}

func buildAlbumQuery(in *resourcerpc.FindAlbumListReq) (int, int, string, string, []any) {
	var opts []query.Option
	if in.Paginate != nil {
		opts = append(opts, query.WithPage(int(in.Paginate.Page)))
		opts = append(opts, query.WithSize(int(in.Paginate.PageSize)))
		opts = append(opts, query.WithSorts(in.Paginate.Sorts...))
	}
	if in.AlbumName != "" {
		opts = append(opts, query.WithCondition("album_name like ?", "%"+in.AlbumName+"%"))
	}
	if in.IsDelete != nil {
		opts = append(opts, query.WithCondition("is_delete = ?", *in.IsDelete))
	}
	return query.NewQueryBuilder(opts...).Build()
}

func buildPageQuery(in *resourcerpc.FindPageListReq) (int, int, string, string, []any) {
	var opts []query.Option
	if in.Paginate != nil {
		opts = append(opts, query.WithPage(int(in.Paginate.Page)))
		opts = append(opts, query.WithSize(int(in.Paginate.PageSize)))
		opts = append(opts, query.WithSorts(in.Paginate.Sorts...))
	}
	if in.PageName != "" {
		opts = append(opts, query.WithCondition("page_name like ?", "%"+in.PageName+"%"))
	}
	return query.NewQueryBuilder(opts...).Build()
}

func buildPageResp(page int, size int, total int64) *resourcerpc.PageResp {
	return &resourcerpc.PageResp{
		Page:     int64(page),
		PageSize: int64(size),
		Total:    total,
	}
}

func collectAlbumIDs(records []*model.TAlbum) []int64 {
	ids := make([]int64, 0, len(records))
	for _, record := range records {
		ids = append(ids, record.Id)
	}
	return ids
}
