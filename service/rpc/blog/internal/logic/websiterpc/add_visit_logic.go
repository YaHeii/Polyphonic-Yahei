package websiterpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/common/enums"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/websiterpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVisitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddVisitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVisitLogic {
	return &AddVisitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加用户访问记录
func (l *AddVisitLogic) AddVisit(in *websiterpc.AddVisitReq) (*websiterpc.AddVisitResp, error) {
	date := currentVisitDate()
	if _, err := l.svcCtx.TVisitDailyStatsModel.Increment(l.ctx, date, enums.VisitTypePv, 1); err != nil {
		return nil, err
	}

	isNewVisitor, err := markDailyVisitor(l.ctx, l.svcCtx.Redis, date, resolveVisitorID(l.ctx))
	if err != nil {
		return nil, err
	}
	if isNewVisitor {
		if _, err := l.svcCtx.TVisitDailyStatsModel.Increment(l.ctx, date, enums.VisitTypeUv, 1); err != nil {
			return nil, err
		}
	}

	return &websiterpc.AddVisitResp{}, nil
}
