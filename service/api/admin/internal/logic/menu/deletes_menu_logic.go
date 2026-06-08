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

type DeletesMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除菜单
func NewDeletesMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesMenuLogic {
	return &DeletesMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletesMenuLogic) DeletesMenu(req *types.IdsReq) (resp *types.BatchResp, err error) {
	out, err := l.svcCtx.PermissionRpc.DeletesMenu(l.ctx, &permissionrpc.DeletesMenuReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{SuccessCount: out.SuccessCount}, nil
}
