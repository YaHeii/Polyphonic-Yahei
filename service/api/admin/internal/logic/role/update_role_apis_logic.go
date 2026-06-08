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

type UpdateRoleApisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新角色接口权限
func NewUpdateRoleApisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleApisLogic {
	return &UpdateRoleApisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleApisLogic) UpdateRoleApis(req *types.UpdateRoleApisReq) (resp *types.EmptyResp, err error) {
	in := &permissionrpc.UpdateRoleApisReq{
		RoleId: req.RoleId,
		ApiIds: req.ApiIds,
	}

	if _, err = l.svcCtx.PermissionRpc.UpdateRoleApis(l.ctx, in); err != nil {
		return nil, err
	}

	if err = l.svcCtx.PermissionHolder.ReloadPolicy(); err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
