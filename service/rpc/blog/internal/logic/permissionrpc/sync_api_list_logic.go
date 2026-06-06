package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncApiListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncApiListLogic {
	return &SyncApiListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 同步接口列表
func (l *SyncApiListLogic) SyncApiList(in *permissionrpc.SyncApiListReq) (*permissionrpc.SyncApiListResp, error) {
	helper := newPermissionHelper(l.ctx, l.svcCtx)

	var count int64
	for _, api := range in.Apis {
		_, saved, err := helper.saveAddApiTree(api, 0)
		if err != nil {
			return nil, err
		}
		count += saved
	}

	return &permissionrpc.SyncApiListResp{SuccessCount: count}, nil
}
