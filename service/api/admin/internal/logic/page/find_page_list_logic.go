// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package page

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/resourcerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取页面列表
func NewFindPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPageListLogic {
	return &FindPageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPageListLogic) FindPageList(req *types.QueryPageReq) (resp *types.PageResp, err error) {
	in := &resourcerpc.FindPageListReq{
		Paginate: &resourcerpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		PageName: req.PageName,
	}

	out, err := l.svcCtx.ResourceRpc.FindPageList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	list := make([]*types.PageBackVO, 0, len(out.List))
	for _, item := range out.List {
		list = append(list, convertPageTypes(item))
	}

	return &types.PageResp{
		Page:     out.Pagination.Page,
		PageSize: out.Pagination.PageSize,
		Total:    out.Pagination.Total,
		List:     list,
	}, nil
}
