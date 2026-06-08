// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package visitor

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindVisitorListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取游客列表
func NewFindVisitorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindVisitorListLogic {
	return &FindVisitorListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindVisitorListLogic) FindVisitorList(req *types.QueryVisitorReq) (resp *types.PageResp, err error) {
	in := &accountrpc.FindVisitorListReq{
		Paginate: &accountrpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		TerminalId: req.TerminalId,
		IpSource:   req.IpSource,
	}

	out, err := l.svcCtx.AccountRpc.FindVisitorList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	list := make([]*types.VisitorBackVO, 0, len(out.List))
	for _, item := range out.List {
		list = append(list, &types.VisitorBackVO{
			Id:         item.Id,
			TerminalId: item.TerminalId,
			Os:         item.Os,
			Browser:    item.Browser,
			IpAddress:  item.IpAddress,
			IpSource:   item.IpSource,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
		})
	}

	return &types.PageResp{
		Page:     out.Pagination.Page,
		PageSize: out.Pagination.PageSize,
		Total:    out.Pagination.Total,
		List:     list,
	}, nil
}
