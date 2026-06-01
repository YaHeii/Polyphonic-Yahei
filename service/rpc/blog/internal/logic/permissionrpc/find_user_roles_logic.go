package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserRolesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserRolesLogic {
	return &FindUserRolesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户角色信息
func (l *FindUserRolesLogic) FindUserRoles(in *permissionrpc.FindUserRolesReq) (*permissionrpc.FindUserRolesResp, error) {
	// todo: add your logic here and delete this line

	return &permissionrpc.FindUserRolesResp{}, nil
}
