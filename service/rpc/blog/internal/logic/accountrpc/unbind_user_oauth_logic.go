package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnbindUserOauthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnbindUserOauthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnbindUserOauthLogic {
	return &UnbindUserOauthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 解绑第三方账号
func (l *UnbindUserOauthLogic) UnbindUserOauth(in *accountrpc.UnbindUserOauthReq) (*accountrpc.UnbindUserOauthResp, error) {
	// todo: add your logic here and delete this line

	return &accountrpc.UnbindUserOauthResp{}, nil
}
