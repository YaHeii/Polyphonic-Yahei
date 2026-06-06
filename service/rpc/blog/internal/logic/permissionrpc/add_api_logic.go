package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddApiLogic {
	return &AddApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建接口
func (l *AddApiLogic) AddApi(in *permissionrpc.AddApiReq) (*permissionrpc.AddApiResp, error) {
	api, _, err := newPermissionHelper(l.ctx, l.svcCtx).saveAddApiTree(in, 0)
	if err != nil {
		return nil, err
	}

	return &permissionrpc.AddApiResp{Api: api}, nil
}
