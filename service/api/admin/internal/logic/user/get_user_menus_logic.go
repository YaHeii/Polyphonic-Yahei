// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizheader"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/permissionrpc"
	"github.com/spf13/cast"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户菜单权限
func NewGetUserMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserMenusLogic {
	return &GetUserMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserMenusLogic) GetUserMenus(req *types.EmptyReq) (resp *types.UserMenusResp, err error) {
	out, err := l.svcCtx.PermissionRpc.FindUserMenus(l.ctx, &permissionrpc.FindUserMenusReq{
		UserId: cast.ToString(l.ctx.Value(bizheader.HeaderUid)),
	})
	if err != nil {
		return nil, err
	}

	list := make([]*types.UserMenu, 0, len(out.List))
	for _, item := range out.List {
		list = append(list, convertUserMenu(item))
	}

	return &types.UserMenusResp{List: list}, nil
}
