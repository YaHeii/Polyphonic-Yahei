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

type FindNoticeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取通知列表
func NewFindNoticeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindNoticeListLogic {
	return &FindNoticeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindNoticeListLogic) FindNoticeList(req *types.QueryNoticeReq) (resp *types.PageResp, err error) {
	in := &noticerpc.FindNoticeListReq{
		Paginate: &noticerpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		Type:          req.Type,
		Level:         req.Level,
		PublishStatus: req.PublishStatus,
		AppName:       req.AppName,
	}

	out, err := l.svcCtx.NoticeRpc.FindNoticeList(l.ctx, in)
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
