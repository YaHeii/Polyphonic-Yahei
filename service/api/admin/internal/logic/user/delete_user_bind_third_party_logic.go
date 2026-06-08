// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserBindThirdPartyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除用户绑定第三方平台账号
func NewDeleteUserBindThirdPartyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserBindThirdPartyLogic {
	return &DeleteUserBindThirdPartyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserBindThirdPartyLogic) DeleteUserBindThirdParty(req *types.DeleteUserBindThirdPartyReq) (resp *types.EmptyResp, err error) {
	_, err = l.svcCtx.AccountRpc.UnbindUserOauth(l.ctx, &accountrpc.UnbindUserOauthReq{
		Platform: req.Platform,
	})
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
