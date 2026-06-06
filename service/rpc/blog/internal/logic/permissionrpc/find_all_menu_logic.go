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
	records, err := l.svcCtx.TMenuModel.FindALL(l.ctx, "")
	if err != nil {
		return nil, err
	}

	return &permissionrpc.FindAllMenuResp{List: buildMenuTree(records)}, nil
}
