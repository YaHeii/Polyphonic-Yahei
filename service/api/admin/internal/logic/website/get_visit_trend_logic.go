// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package website

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/websiterpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVisitTrendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取访客数据趋势
func NewGetVisitTrendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVisitTrendLogic {
	return &GetVisitTrendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVisitTrendLogic) GetVisitTrend(req *types.GetVisitTrendReq) (resp *types.GetVisitTrendResp, err error) {
	out, err := l.svcCtx.WebsiteRpc.FindVisitTrend(l.ctx, &websiterpc.FindVisitTrendReq{
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetVisitTrendResp{
		VisitTrend: mergeVisitTrend(out.PvTrend, out.UvTrend),
	}, nil
}
