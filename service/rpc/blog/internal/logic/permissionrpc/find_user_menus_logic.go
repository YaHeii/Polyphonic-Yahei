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
	records, err := l.svcCtx.TMenuModel.FindByUserID(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	return &permissionrpc.FindUserMenusResp{List: buildMenuTree(records)}, nil
}
