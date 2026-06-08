// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package role

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/permissionrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRoleResourcesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取角色资源列表
func NewFindRoleResourcesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoleResourcesLogic {
	return &FindRoleResourcesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindRoleResourcesLogic) FindRoleResources(req *types.IdReq) (resp *types.RoleResourcesResp, err error) {
	in := &permissionrpc.FindRoleResourcesReq{
		Id: req.Id,
	}

	out, err := l.svcCtx.PermissionRpc.FindRoleResources(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.RoleResourcesResp{
		RoleId:  out.RoleId,
		ApiIds:  out.ApiIds,
		MenuIds: out.MenuIds,
	}, nil
}
