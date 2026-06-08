package album

import (
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/resourcerpc"
)

func convertAlbumTypes(out *resourcerpc.Album) *types.AlbumBackVO {
	return &types.AlbumBackVO{
		Id:         out.Id,
		AlbumName:  out.AlbumName,
		AlbumDesc:  out.AlbumDesc,
		AlbumCover: out.AlbumCover,
		IsDelete:   out.IsDelete,
		Status:     out.Status,
		CreatedAt:  out.CreatedAt,
		UpdatedAt:  out.UpdatedAt,
		PhotoCount: out.PhotoCount,
	}
}
