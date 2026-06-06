package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleLogic {
	return &AddRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建角色
func (l *AddRoleLogic) AddRole(in *permissionrpc.AddRoleReq) (*permissionrpc.AddRoleResp, error) {
	entity := convertAddRoleIn(in)
	if _, err := l.svcCtx.TRoleModel.Save(l.ctx, entity); err != nil {
		return nil, err
	}

	return &permissionrpc.AddRoleResp{Role: convertRoleOut(entity)}, nil
}
