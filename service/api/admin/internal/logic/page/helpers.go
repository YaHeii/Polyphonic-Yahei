package page

import (
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/resourcerpc"
)

func convertPageTypes(out *resourcerpc.Page) *types.PageBackVO {
	if out == nil {
		return nil
	}

	return &types.PageBackVO{
		Id:             out.Id,
		PageName:       out.PageName,
		PageLabel:      out.PageLabel,
		PageCover:      out.PageCover,
		IsCarousel:     boolToInt64(out.IsCarousel),
		CarouselCovers: out.CarouselCovers,
		CreatedAt:      out.CreatedAt,
		UpdatedAt:      out.UpdatedAt,
	}
}

func boolToInt64(v bool) int64 {
	if v {
		return 1
	}
	return 0
}
