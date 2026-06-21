// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package website

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/newsrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminHomeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取后台首页信息
func NewGetAdminHomeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminHomeInfoLogic {
	return &GetAdminHomeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminHomeInfoLogic) GetAdminHomeInfo(req *types.EmptyReq) (resp *types.AdminHomeInfo, err error) {
	users, err := l.svcCtx.AccountRpc.AnalysisUser(l.ctx, &accountrpc.AnalysisUserReq{})
	if err != nil {
		return nil, err
	}

	articles, err := l.svcCtx.ArticleRpc.AnalysisArticle(l.ctx, &articlerpc.AnalysisArticleReq{})
	if err != nil {
		return nil, err
	}

	messages, err := l.svcCtx.NewsRpc.AnalysisMessage(l.ctx, &newsrpc.AnalysisMessageReq{})
	if err != nil {
		return nil, err
	}

	ranks := make([]*types.ArticleViewVO, 0, len(articles.ArticleRankList))
	for _, item := range articles.ArticleRankList {
		ranks = append(ranks, &types.ArticleViewVO{
			Id:           item.Id,
			ArticleTitle: item.ArticleTitle,
			ViewCount:    item.ViewCount,
		})
	}

	categoryList := make([]*types.CategoryVO, 0, len(articles.CategoryList))
	for _, item := range articles.CategoryList {
		categoryList = append(categoryList, &types.CategoryVO{
			Id:           item.Id,
			CategoryName: item.CategoryName,
			ArticleCount: item.ArticleCount,
		})
	}

	tagList := make([]*types.TagVO, 0, len(articles.TagList))
	for _, item := range articles.TagList {
		tagList = append(tagList, &types.TagVO{
			TagName:      item.TagName,
			ArticleCount: item.ArticleCount,
		})
	}

	archiveList, err := l.svcCtx.ArticleRpc.FindArticleList(l.ctx, &articlerpc.FindArticleListReq{})
	if err != nil {
		return nil, err
	}

	return &types.AdminHomeInfo{
		UserCount:         users.UserCount,
		ArticleCount:      articles.ArticleCount,
		MessageCount:      messages.MessageCount,
		CategoryList:      categoryList,
		TagList:           tagList,
		ArticleViewRanks:  ranks,
		ArticleStatistics: buildArticleStatistics(archiveList.List),
	}, nil
}
