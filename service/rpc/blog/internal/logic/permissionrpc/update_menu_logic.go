package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新菜单
func (l *UpdateMenuLogic) UpdateMenu(in *permissionrpc.UpdateMenuReq) (*permissionrpc.UpdateMenuResp, error) {
	entity := convertUpdateMenuIn(in)
	if _, err := l.svcCtx.TMenuModel.Save(l.ctx, entity); err != nil {
		return nil, err
	}

	return &permissionrpc.UpdateMenuResp{Menu: convertMenuOut(entity)}, nil
}
