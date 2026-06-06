package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRoleResourcesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindRoleResourcesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoleResourcesLogic {
	return &FindRoleResourcesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询角色资源权限
func (l *FindRoleResourcesLogic) FindRoleResources(in *permissionrpc.FindRoleResourcesReq) (*permissionrpc.FindRoleResourcesResp, error) {
	apiIDs, err := l.svcCtx.TRoleApiModel.FindApiIDsByRoleID(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	menuIDs, err := l.svcCtx.TRoleMenuModel.FindMenuIDsByRoleID(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &permissionrpc.FindRoleResourcesResp{
		RoleId:  in.Id,
		ApiIds:  apiIDs,
		MenuIds: menuIDs,
	}, nil
}
