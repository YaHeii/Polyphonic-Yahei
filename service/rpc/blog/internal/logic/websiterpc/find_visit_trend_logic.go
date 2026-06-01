package websiterpclogic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &websiterpc.FindVisitTrendResp{}, nil
}
