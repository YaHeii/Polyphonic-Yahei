package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnalysisUserAreasLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnalysisUserAreasLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnalysisUserAreasLogic {
	return &AnalysisUserAreasLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户分布区域
func (l *AnalysisUserAreasLogic) AnalysisUserAreas(in *accountrpc.AnalysisUserAreasReq) (*accountrpc.AnalysisUserAreasResp, error) {
	// todo: add your logic here and delete this line

	return &accountrpc.AnalysisUserAreasResp{}, nil
}
