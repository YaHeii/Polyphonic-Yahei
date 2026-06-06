package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuLogic {
	return &AddMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建菜单
func (l *AddMenuLogic) AddMenu(in *permissionrpc.AddMenuReq) (*permissionrpc.AddMenuResp, error) {
	menu, _, err := newPermissionHelper(l.ctx, l.svcCtx).saveAddMenuTree(in, 0)
	if err != nil {
		return nil, err
	}

	return &permissionrpc.AddMenuResp{Menu: menu}, nil
}
