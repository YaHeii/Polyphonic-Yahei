// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizheader"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/permissionrpc"
	"github.com/spf13/cast"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.EmptyReq) (resp *types.UserInfoResp, err error) {
	userID := cast.ToString(l.ctx.Value(bizheader.HeaderUid))

	info, err := l.svcCtx.AccountRpc.GetUserInfo(l.ctx, &accountrpc.GetUserInfoReq{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	oauthInfo, err := l.svcCtx.AccountRpc.GetUserOauthInfo(l.ctx, &accountrpc.GetUserOauthInfoReq{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	menus, err := l.svcCtx.PermissionRpc.FindUserMenus(l.ctx, &permissionrpc.FindUserMenusReq{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	return convertUserInfo(info.User, oauthInfo, menus), nil
}
