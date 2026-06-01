package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindUserOauthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindUserOauthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindUserOauthLogic {
	return &BindUserOauthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户第三方账号
func (l *BindUserOauthLogic) BindUserOauth(in *accountrpc.BindUserOauthReq) (*accountrpc.BindUserOauthResp, error) {
	// todo: add your logic here and delete this line

	return &accountrpc.BindUserOauthResp{}, nil
}
