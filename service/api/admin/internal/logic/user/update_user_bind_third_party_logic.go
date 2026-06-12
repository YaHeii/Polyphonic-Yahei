// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	authlogic "github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/logic/auth"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserBindThirdPartyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户绑定第三方平台账号
func NewUpdateUserBindThirdPartyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserBindThirdPartyLogic {
	return &UpdateUserBindThirdPartyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserBindThirdPartyLogic) UpdateUserBindThirdParty(req *types.UpdateUserBindThirdPartyReq) (resp *types.EmptyResp, err error) {
	info, err := authlogic.GetOauthIdentityForUser(l.ctx, l.svcCtx, req.Platform, req.Code)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.AccountRpc.BindUserOauth(l.ctx, &accountrpc.BindUserOauthReq{
		Platform: req.Platform,
		OpenId:   info.OpenId,
		Nickname: info.NickName,
		Avatar:   info.Avatar,
	})
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
