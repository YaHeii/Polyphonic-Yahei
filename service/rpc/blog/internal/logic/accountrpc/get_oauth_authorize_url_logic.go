package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOauthAuthorizeUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOauthAuthorizeUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOauthAuthorizeUrlLogic {
	return &GetOauthAuthorizeUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取第三方登录授权地址
func (l *GetOauthAuthorizeUrlLogic) GetOauthAuthorizeUrl(in *accountrpc.GetOauthAuthorizeUrlReq) (*accountrpc.GetOauthAuthorizeUrlResp, error) {
	// todo: add your logic here and delete this line

	return &accountrpc.GetOauthAuthorizeUrlResp{}, nil
}
