package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UpdateRoleMenusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleMenusLogic {
	return &UpdateRoleMenusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新角色菜单
func (l *UpdateRoleMenusLogic) UpdateRoleMenus(in *permissionrpc.UpdateRoleMenusReq) (*permissionrpc.UpdateRoleMenusResp, error) {
	err := l.svcCtx.SqlConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		conn := sqlx.NewSqlConnFromSession(session)
		_, err := model.NewTRoleMenuModel(conn).ReplaceByRoleID(ctx, in.RoleId, in.MenuIds)
		return err
	})
	if err != nil {
		return nil, err
	}

	return &permissionrpc.UpdateRoleMenusResp{}, nil
}
