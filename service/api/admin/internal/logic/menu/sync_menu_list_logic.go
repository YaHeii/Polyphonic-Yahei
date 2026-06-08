// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package menu

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/permissionrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 同步菜单列表
func NewSyncMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncMenuListLogic {
	return &SyncMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncMenuListLogic) SyncMenuList(req *types.SyncMenuReq) (resp *types.BatchResp, err error) {
	menus := make([]*permissionrpc.AddMenuReq, 0, len(req.Menus))
	for _, menu := range req.Menus {
		menus = append(menus, convertMenuPb(menu))
	}

	out, err := l.svcCtx.PermissionRpc.SyncMenuList(l.ctx, &permissionrpc.SyncMenuListReq{
		Menus: menus,
	})
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{SuccessCount: out.SuccessCount}, nil
}
