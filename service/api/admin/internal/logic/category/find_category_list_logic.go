// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package category

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/articlerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取文章分类列表
func NewFindCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCategoryListLogic {
	return &FindCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindCategoryListLogic) FindCategoryList(req *types.QueryCategoryReq) (resp *types.CategoryPageResp, err error) {
	in := &articlerpc.FindCategoryListReq{
		Paginate: &articlerpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		CategoryName: req.CategoryName,
	}

	out, err := l.svcCtx.ArticleRpc.FindCategoryList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	list := make([]*types.CategoryBackVO, 0, len(out.List))
	for _, v := range out.List {
		list = append(list, convertCategoryDetailsTypes(v))
	}

	resp = &types.CategoryPageResp{
		List: list,
	}
	if out.Pagination != nil {
		resp.Page = out.Pagination.Page
		resp.PageSize = out.Pagination.PageSize
		resp.Total = out.Pagination.Total
	}
	return resp, nil
}
