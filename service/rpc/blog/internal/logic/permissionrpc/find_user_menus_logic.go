package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserMenusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserMenusLogic {
	return &FindUserMenusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户菜单权限
func (l *FindUserMenusLogic) FindUserMenus(in *permissionrpc.FindUserMenusReq) (*permissionrpc.FindUserMenusResp, error) {
	// todo: add your logic here and delete this line

	return &permissionrpc.FindUserMenusResp{}, nil
}
