package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncMenuListLogic {
	return &SyncMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 同步菜单列表
func (l *SyncMenuListLogic) SyncMenuList(in *permissionrpc.SyncMenuListReq) (*permissionrpc.SyncMenuListResp, error) {
	helper := newPermissionHelper(l.ctx, l.svcCtx)

	var count int64
	for _, menu := range in.Menus {
		_, saved, err := helper.saveAddMenuTree(menu, 0)
		if err != nil {
			return nil, err
		}
		count += saved
	}

	return &permissionrpc.SyncMenuListResp{SuccessCount: count}, nil
}
