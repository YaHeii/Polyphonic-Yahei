// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package article

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/authctx"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/articlerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 保存文章
func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateArticleLogic) UpdateArticle(req *types.NewArticleReq) (resp *types.ArticleBackVO, err error) {
	in := &articlerpc.UpdateArticleReq{
		Id:             req.Id,
		UserId:         authctx.CurrentUserID(l.ctx),
		ArticleCover:   req.ArticleCover,
		ArticleTitle:   req.ArticleTitle,
		ArticleContent: req.ArticleContent,
		ArticleType:    req.ArticleType,
		OriginalUrl:    req.OriginalUrl,
		IsTop:          req.IsTop,
		Status:         req.Status,
		CategoryName:   req.CategoryName,
		TagNameList:    req.TagNameList,
	}

	out, err := l.svcCtx.ArticleRpc.UpdateArticle(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convertArticlePreviewTypes(out.Article), nil
}
