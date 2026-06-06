package websiterpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/common/enums"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/websiterpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindVisitTrendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindVisitTrendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindVisitTrendLogic {
	return &FindVisitTrendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户访问趋势
func (l *FindVisitTrendLogic) FindVisitTrend(in *websiterpc.FindVisitTrendReq) (*websiterpc.FindVisitTrendResp, error) {
	start, end, err := normalizeVisitRange(in.StartDate, in.EndDate)
	if err != nil {
		return nil, err
	}

	uvRecords, err := l.svcCtx.TVisitDailyStatsModel.FindByDateRange(l.ctx, start.Format(visitDateLayout), end.Format(visitDateLayout), enums.VisitTypeUv)
	if err != nil {
		return nil, err
	}
	pvRecords, err := l.svcCtx.TVisitDailyStatsModel.FindByDateRange(l.ctx, start.Format(visitDateLayout), end.Format(visitDateLayout), enums.VisitTypePv)
	if err != nil {
		return nil, err
	}

	return &websiterpc.FindVisitTrendResp{
		UvTrend: buildVisitTrend(uvRecords, start, end),
		PvTrend: buildVisitTrend(pvRecords, start, end),
	}, nil
}
