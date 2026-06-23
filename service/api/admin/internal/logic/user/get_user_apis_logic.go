// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/authctx"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/permissionrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserApisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户接口权限
func NewGetUserApisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserApisLogic {
	return &GetUserApisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserApisLogic) GetUserApis(req *types.EmptyReq) (resp *types.UserApisResp, err error) {
	out, err := l.svcCtx.PermissionRpc.FindUserApis(l.ctx, &permissionrpc.FindUserApisReq{
		UserId: authctx.CurrentUserID(l.ctx),
	})
	if err != nil {
		return nil, err
	}

	list := make([]*types.UserApi, 0, len(out.List))
	for _, item := range out.List {
		list = append(list, convertUserApi(item))
	}

	return &types.UserApisResp{List: list}, nil
}
