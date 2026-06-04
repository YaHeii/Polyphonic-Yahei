package articlerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新文章
func (l *UpdateArticleLogic) UpdateArticle(in *articlerpc.UpdateArticleReq) (*articlerpc.UpdateArticleResp, error) {
	helper := NewArticleHelperLogic(l.ctx, l.svcCtx)

	entity, err := l.svcCtx.TArticleModel.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	// 插入文章分类
	categoryId, err := helper.findOrAddCategory(in.CategoryName)
	if err != nil {
		return nil, err
	}

	tagNames, err := helper.ensureTags(in.TagNameList)
	if err != nil {
		return nil, err
	}

	entity.ArticleTitle = in.ArticleTitle
	entity.ArticleContent = in.ArticleContent
	entity.ArticleCover = in.ArticleCover
	entity.ArticleType = in.ArticleType
	entity.OriginalUrl = in.OriginalUrl
	entity.Tags = tagNames
	entity.IsTop = in.IsTop
	entity.Status = in.Status
	entity.CategoryId = categoryId

	err = l.svcCtx.TArticleModel.Update(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return &articlerpc.UpdateArticleResp{
		Article: helper.convertArticlePreviewOut(entity),
	}, nil
}
