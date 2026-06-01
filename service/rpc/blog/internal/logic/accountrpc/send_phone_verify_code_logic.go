package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendPhoneVerifyCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendPhoneVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendPhoneVerifyCodeLogic {
	return &SendPhoneVerifyCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送手机号验证码
func (l *SendPhoneVerifyCodeLogic) SendPhoneVerifyCode(in *accountrpc.SendPhoneVerifyCodeReq) (*accountrpc.SendPhoneVerifyCodeResp, error) {
	// todo: add your logic here and delete this line

	return &accountrpc.SendPhoneVerifyCodeResp{}, nil
}
