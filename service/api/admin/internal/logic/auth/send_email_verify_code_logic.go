// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package auth

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailVerifyCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发送邮件验证码
func NewSendEmailVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailVerifyCodeLogic {
	return &SendEmailVerifyCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendEmailVerifyCodeLogic) SendEmailVerifyCode(req *types.SendEmailVerifyCodeReq) (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
