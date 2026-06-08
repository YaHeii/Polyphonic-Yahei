// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package notice

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/noticerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserNoticeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询用户通知列表
func NewFindUserNoticeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserNoticeListLogic {
	return &FindUserNoticeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindUserNoticeListLogic) FindUserNoticeList(req *types.QueryUserNoticeReq) (resp *types.PageResp, err error) {
	in := &noticerpc.FindUserNoticeListReq{
		Paginate: &noticerpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
	}

	out, err := l.svcCtx.NoticeRpc.FindUserNoticeList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	list := make([]*types.NoticeBackVO, 0, len(out.List))
	for _, item := range out.List {
		list = append(list, convertNoticeOut(item))
	}

	return &types.PageResp{
		Page:     out.Pagination.Page,
		PageSize: out.Pagination.PageSize,
		Total:    out.Pagination.Total,
		List:     list,
	}, nil
}
