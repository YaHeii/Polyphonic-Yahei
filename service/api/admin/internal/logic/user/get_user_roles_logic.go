// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/authctx"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/permissionrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户角色
func NewGetUserRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRolesLogic {
	return &GetUserRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserRolesLogic) GetUserRoles(req *types.EmptyReq) (resp *types.UserRolesResp, err error) {
	out, err := l.svcCtx.PermissionRpc.FindUserRoles(l.ctx, &permissionrpc.FindUserRolesReq{
		UserId: authctx.CurrentUserID(l.ctx),
	})
	if err != nil {
		return nil, err
	}

	list := make([]*types.UserRole, 0, len(out.List))
	for _, item := range out.List {
		list = append(list, &types.UserRole{
			Id:          item.Id,
			RoleKey:     item.RoleKey,
			RoleComment: item.RoleComment,
		})
	}

	return &types.UserRolesResp{List: list}, nil
}
