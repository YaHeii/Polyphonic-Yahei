package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesApiLogic {
	return &DeletesApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除接口
func (l *DeletesApiLogic) DeletesApi(in *permissionrpc.DeletesApiReq) (*permissionrpc.DeletesApiResp, error) {
	if _, err := l.svcCtx.TRoleApiModel.DeleteByApiIDs(l.ctx, in.Ids); err != nil {
		return nil, err
	}

	rows, err := l.svcCtx.TApiModel.Deletes(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &permissionrpc.DeletesApiResp{SuccessCount: rows}, nil
}
