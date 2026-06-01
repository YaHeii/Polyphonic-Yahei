package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClientInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetClientInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClientInfoLogic {
	return &GetClientInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取客户端信息
func (l *GetClientInfoLogic) GetClientInfo(in *accountrpc.GetClientInfoReq) (*accountrpc.GetClientInfoResp, error) {
	// todo: add your logic here and delete this line

	return &accountrpc.GetClientInfoResp{}, nil
}
