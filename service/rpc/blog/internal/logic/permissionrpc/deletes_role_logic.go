package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesRoleLogic {
	return &DeletesRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除角色
func (l *DeletesRoleLogic) DeletesRole(in *permissionrpc.DeletesRoleReq) (*permissionrpc.DeletesRoleResp, error) {
	if _, err := l.svcCtx.TRoleApiModel.DeleteByRoleIDs(l.ctx, in.Ids); err != nil {
		return nil, err
	}
	if _, err := l.svcCtx.TRoleMenuModel.DeleteByRoleIDs(l.ctx, in.Ids); err != nil {
		return nil, err
	}
	if _, err := l.svcCtx.TUserRoleModel.DeleteByRoleIDs(l.ctx, in.Ids); err != nil {
		return nil, err
	}

	rows, err := l.svcCtx.TRoleModel.Deletes(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &permissionrpc.DeletesRoleResp{SuccessCount: rows}, nil
}
