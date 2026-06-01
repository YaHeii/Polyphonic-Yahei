package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailVerifyCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailVerifyCodeLogic {
	return &SendEmailVerifyCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送邮件验证码
func (l *SendEmailVerifyCodeLogic) SendEmailVerifyCode(in *accountrpc.SendEmailVerifyCodeReq) (*accountrpc.SendEmailVerifyCodeResp, error) {
	// todo: add your logic here and delete this line

	return &accountrpc.SendEmailVerifyCodeResp{}, nil
}
