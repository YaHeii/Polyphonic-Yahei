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

type CleanMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 清空菜单列表
func NewCleanMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CleanMenuListLogic {
	return &CleanMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CleanMenuListLogic) CleanMenuList(req *types.EmptyReq) (resp *types.BatchResp, err error) {
	out, err := l.svcCtx.PermissionRpc.CleanMenuList(l.ctx, &permissionrpc.CleanMenuListReq{})
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{SuccessCount: out.SuccessCount}, nil
}
