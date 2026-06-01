package socialrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/socialrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesTalkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesTalkLogic {
	return &DeletesTalkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除说说
func (l *DeletesTalkLogic) DeletesTalk(in *socialrpc.DeletesTalkReq) (*socialrpc.DeletesTalkResp, error) {
	// todo: add your logic here and delete this line

	return &socialrpc.DeletesTalkResp{}, nil
}
