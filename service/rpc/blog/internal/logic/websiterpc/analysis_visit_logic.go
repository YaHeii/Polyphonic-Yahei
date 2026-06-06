package websiterpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/common/enums"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/websiterpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnalysisVisitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnalysisVisitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnalysisVisitLogic {
	return &AnalysisVisitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户日浏览量分析
func (l *AnalysisVisitLogic) AnalysisVisit(in *websiterpc.AnalysisVisitReq) (*websiterpc.AnalysisVisitResp, error) {
	today := currentVisitDate()
	yesterday := previousVisitDate()

	todayUV, err := l.svcCtx.TVisitDailyStatsModel.FindCount(l.ctx, today, enums.VisitTypeUv)
	if err != nil {
		return nil, err
	}
	yesterdayUV, err := l.svcCtx.TVisitDailyStatsModel.FindCount(l.ctx, yesterday, enums.VisitTypeUv)
	if err != nil {
		return nil, err
	}
	totalUV, err := l.svcCtx.TVisitDailyStatsModel.SumByVisitType(l.ctx, enums.VisitTypeUv)
	if err != nil {
		return nil, err
	}

	todayPV, err := l.svcCtx.TVisitDailyStatsModel.FindCount(l.ctx, today, enums.VisitTypePv)
	if err != nil {
		return nil, err
	}
	yesterdayPV, err := l.svcCtx.TVisitDailyStatsModel.FindCount(l.ctx, yesterday, enums.VisitTypePv)
	if err != nil {
		return nil, err
	}
	totalPV, err := l.svcCtx.TVisitDailyStatsModel.SumByVisitType(l.ctx, enums.VisitTypePv)
	if err != nil {
		return nil, err
	}

	return &websiterpc.AnalysisVisitResp{
		TodayUvCount: todayUV,
		TotalUvCount: totalUV,
		UvGrowthRate: calculateGrowthRate(todayUV, yesterdayUV),
		TodayPvCount: todayPV,
		TotalPvCount: totalPV,
		PvGrowthRate: calculateGrowthRate(todayPV, yesterdayPV),
	}, nil
}
