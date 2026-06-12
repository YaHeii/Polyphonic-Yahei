// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package auth

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

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
	url, err := getOauthAuthorizeURL(l.svcCtx, currentAppName(l.ctx), req.Platform, req.State)
	if err != nil {
		return nil, err
	}

	return &types.GetOauthAuthorizeUrlResp{
		AuthorizeUrl: url,
	}, nil
}
