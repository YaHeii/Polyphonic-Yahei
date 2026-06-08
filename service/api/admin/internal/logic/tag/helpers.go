package tag

import (
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/articlerpc"
)

func convertTagTypes(out *articlerpc.Tag) *types.TagBackVO {
	return &types.TagBackVO{
		Id:           out.Id,
		TagName:      out.TagName,
		ArticleCount: 0,
		CreatedAt:    out.CreatedAt,
		UpdatedAt:    out.UpdatedAt,
	}
}

func convertTagDetailsTypes(out *articlerpc.TagDetails) *types.TagBackVO {
	return &types.TagBackVO{
		Id:           out.Id,
		TagName:      out.TagName,
		ArticleCount: out.ArticleCount,
		CreatedAt:    out.CreatedAt,
		UpdatedAt:    out.UpdatedAt,
	}
}
