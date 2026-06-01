package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleApisLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleApisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleApisLogic {
	return &UpdateRoleApisLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新角色资源
func (l *UpdateRoleApisLogic) UpdateRoleApis(in *permissionrpc.UpdateRoleApisReq) (*permissionrpc.UpdateRoleApisResp, error) {
	// todo: add your logic here and delete this line

	return &permissionrpc.UpdateRoleApisResp{}, nil
}
