package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAllMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAllMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllMenuLogic {
	return &FindAllMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查找所有菜单
func (l *FindAllMenuLogic) FindAllMenu(in *permissionrpc.FindAllMenuReq) (*permissionrpc.FindAllMenuResp, error) {
	// todo: add your logic here and delete this line

	return &permissionrpc.FindAllMenuResp{}, nil
}
