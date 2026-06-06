package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UpdateRoleApisLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleApisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleApisLogic {
	return &UpdateRoleApisLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新角色资源
func (l *UpdateRoleApisLogic) UpdateRoleApis(in *permissionrpc.UpdateRoleApisReq) (*permissionrpc.UpdateRoleApisResp, error) {
	err := l.svcCtx.SqlConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		conn := sqlx.NewSqlConnFromSession(session)
		_, err := model.NewTRoleApiModel(conn).ReplaceByRoleID(ctx, in.RoleId, in.ApiIds)
		return err
	})
	if err != nil {
		return nil, err
	}

	return &permissionrpc.UpdateRoleApisResp{}, nil
}
