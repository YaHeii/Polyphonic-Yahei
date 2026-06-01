package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserApisLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserApisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserApisLogic {
	return &FindUserApisLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户接口权限
func (l *FindUserApisLogic) FindUserApis(in *permissionrpc.FindUserApisReq) (*permissionrpc.FindUserApisResp, error) {
	// todo: add your logic here and delete this line

	return &permissionrpc.FindUserApisResp{}, nil
}
