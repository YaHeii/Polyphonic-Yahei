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
	// todo: add your logic here and delete this line

	return &permissionrpc.DeletesMenuResp{}, nil
}
