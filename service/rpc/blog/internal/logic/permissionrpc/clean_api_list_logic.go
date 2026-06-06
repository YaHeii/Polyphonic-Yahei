package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CleanApiListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCleanApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CleanApiListLogic {
	return &CleanApiListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 清空接口列表
func (l *CleanApiListLogic) CleanApiList(in *permissionrpc.CleanApiListReq) (*permissionrpc.CleanApiListResp, error) {
	if _, err := l.svcCtx.TRoleApiModel.Clean(l.ctx); err != nil {
		return nil, err
	}

	rows, err := l.svcCtx.TApiModel.Clean(l.ctx)
	if err != nil {
		return nil, err
	}

	return &permissionrpc.CleanApiListResp{SuccessCount: rows}, nil
}
