package category

import (
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/articlerpc"
)

func convertCategoryTypes(out *articlerpc.Category) *types.CategoryBackVO {
	return &types.CategoryBackVO{
		Id:           out.Id,
		CategoryName: out.CategoryName,
		ArticleCount: 0,
		CreatedAt:    out.CreatedAt,
		UpdatedAt:    out.UpdatedAt,
	}
}

func convertCategoryDetailsTypes(out *articlerpc.CategoryDetails) *types.CategoryBackVO {
	return &types.CategoryBackVO{
		Id:           out.Id,
		CategoryName: out.CategoryName,
		ArticleCount: out.ArticleCount,
		CreatedAt:    out.CreatedAt,
		UpdatedAt:    out.UpdatedAt,
	}
}
