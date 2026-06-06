package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesMenuLogic {
	return &DeletesMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除菜单
func (l *DeletesMenuLogic) DeletesMenu(in *permissionrpc.DeletesMenuReq) (*permissionrpc.DeletesMenuResp, error) {
	if _, err := l.svcCtx.TRoleMenuModel.DeleteByMenuIDs(l.ctx, in.Ids); err != nil {
		return nil, err
	}

	rows, err := l.svcCtx.TMenuModel.Deletes(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &permissionrpc.DeletesMenuResp{SuccessCount: rows}, nil
}
