// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package auth

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOauthAuthorizeUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 第三方登录授权地址
func NewGetOauthAuthorizeUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOauthAuthorizeUrlLogic {
	return &GetOauthAuthorizeUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOauthAuthorizeUrlLogic) GetOauthAuthorizeUrl(req *types.GetOauthAuthorizeUrlReq) (resp *types.GetOauthAuthorizeUrlResp, err error) {
	out, err := l.svcCtx.AccountRpc.GetOauthAuthorizeUrl(l.ctx, &accountrpc.GetOauthAuthorizeUrlReq{
		Platform: req.Platform,
		State:    req.State,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetOauthAuthorizeUrlResp{
		AuthorizeUrl: out.GetAuthorizeUrl(),
	}, nil
}
