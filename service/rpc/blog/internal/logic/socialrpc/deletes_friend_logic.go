package socialrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/socialrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesFriendLogic {
	return &DeletesFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除友链
func (l *DeletesFriendLogic) DeletesFriend(in *socialrpc.DeletesFriendReq) (*socialrpc.DeletesFriendResp, error) {
	// todo: add your logic here and delete this line

	return &socialrpc.DeletesFriendResp{}, nil
}
