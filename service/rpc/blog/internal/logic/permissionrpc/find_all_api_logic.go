package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAllApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAllApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllApiLogic {
	return &FindAllApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查找所有接口
func (l *FindAllApiLogic) FindAllApi(in *permissionrpc.FindAllApiReq) (*permissionrpc.FindAllApiResp, error) {
	records, err := l.svcCtx.TApiModel.FindALL(l.ctx, "")
	if err != nil {
		return nil, err
	}

	return &permissionrpc.FindAllApiResp{List: buildApiTree(records)}, nil
}
