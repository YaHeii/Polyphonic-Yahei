// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package article

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizheader"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/articlerpc"
	"github.com/spf13/cast"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加文章
func NewAddArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleLogic {
	return &AddArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddArticleLogic) AddArticle(req *types.NewArticleReq) (resp *types.ArticleBackVO, err error) {
	in := &articlerpc.AddArticleReq{
		Id:             req.Id,
		UserId:         cast.ToString(l.ctx.Value(bizheader.HeaderUid)),
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

	out, err := l.svcCtx.ArticleRpc.AddArticle(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convertArticlePreviewTypes(out.Article), nil
}
