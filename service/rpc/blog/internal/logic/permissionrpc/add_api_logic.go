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
	// todo: add your logic here and delete this line

	return &permissionrpc.AddApiResp{}, nil
}
