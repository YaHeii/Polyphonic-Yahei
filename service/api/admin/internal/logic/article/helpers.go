package article

import (
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/articlerpc"
)

func convertArticlePreviewTypes(in *articlerpc.ArticlePreview) *types.ArticleBackVO {
	return &types.ArticleBackVO{
		Id: in.Id,
	}
}

func convertArticleTypes(in *articlerpc.ArticleDetails) (out *types.ArticleBackVO) {
	out = &types.ArticleBackVO{
		Id:             in.Id,
		ArticleCover:   in.ArticleCover,
		ArticleTitle:   in.ArticleTitle,
		ArticleContent: in.ArticleContent,
		ArticleType:    in.ArticleType,
		OriginalUrl:    in.OriginalUrl,
		IsTop:          in.IsTop,
		IsDelete:       in.IsDelete,
		Status:         in.Status,
		CreatedAt:      in.CreatedAt,
		UpdatedAt:      in.UpdatedAt,
		CategoryName:   "",
		TagNameList:    make([]string, 0),
		LikeCount:      in.LikeCount,
		ViewsCount:     in.ViewCount,
	}

	if in.Category != nil {
		out.CategoryName = in.Category.CategoryName
	}

	if in.TagList != nil {
		for _, tag := range in.TagList {
			out.TagNameList = append(out.TagNameList, tag.TagName)
		}
	}

	return
}
