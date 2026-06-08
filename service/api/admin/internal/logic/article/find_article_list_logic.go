// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package article

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/articlerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询文章列表
func NewFindArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleListLogic {
	return &FindArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleListLogic) FindArticleList(req *types.QueryArticleReq) (resp *types.PageResp, err error) {
	in := &articlerpc.FindArticleListReq{
		Paginate: &articlerpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		ArticleTitle: req.ArticleTitle,
		ArticleType:  req.ArticleType,
		CategoryName: req.CategoryName,
		TagName:      req.TagName,
		IsTop:        &req.IsTop,
		IsDelete:     &req.IsDelete,
		Status:       req.Status,
	}

	out, err := l.svcCtx.ArticleRpc.FindArticleList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	list := make([]*types.ArticleBackVO, 0, len(out.List))
	for _, v := range out.List {
		list = append(list, convertArticleTypes(v))
	}

	resp = &types.PageResp{
		List: list,
	}
	if out.Pagination != nil {
		resp.Page = out.Pagination.Page
		resp.PageSize = out.Pagination.PageSize
		resp.Total = out.Pagination.Total
	}
	return resp, nil
}
