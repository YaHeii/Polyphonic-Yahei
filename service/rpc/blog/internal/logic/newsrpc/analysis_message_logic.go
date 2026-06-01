package newsrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/newsrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnalysisMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnalysisMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnalysisMessageLogic {
	return &AnalysisMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 消息数据分析
func (l *AnalysisMessageLogic) AnalysisMessage(in *newsrpc.AnalysisMessageReq) (*newsrpc.AnalysisMessageResp, error) {
	// todo: add your logic here and delete this line

	return &newsrpc.AnalysisMessageResp{}, nil
}
