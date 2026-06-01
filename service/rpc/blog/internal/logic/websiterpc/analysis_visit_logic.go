package websiterpclogic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &websiterpc.AnalysisVisitResp{}, nil
}
