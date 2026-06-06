package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CleanMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCleanMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CleanMenuListLogic {
	return &CleanMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 清空菜单列表
func (l *CleanMenuListLogic) CleanMenuList(in *permissionrpc.CleanMenuListReq) (*permissionrpc.CleanMenuListResp, error) {
	if _, err := l.svcCtx.TRoleMenuModel.Clean(l.ctx); err != nil {
		return nil, err
	}

	rows, err := l.svcCtx.TMenuModel.Clean(l.ctx)
	if err != nil {
		return nil, err
	}

	return &permissionrpc.CleanMenuListResp{SuccessCount: rows}, nil
}
