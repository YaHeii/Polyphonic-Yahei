package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesRoleLogic {
	return &DeletesRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除角色
func (l *DeletesRoleLogic) DeletesRole(in *permissionrpc.DeletesRoleReq) (*permissionrpc.DeletesRoleResp, error) {
	// todo: add your logic here and delete this line

	return &permissionrpc.DeletesRoleResp{}, nil
}
