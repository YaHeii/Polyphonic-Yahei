// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package auth

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/syslogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 邮箱登录
func NewEmailLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailLoginLogic {
	return &EmailLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailLoginLogic) EmailLogin(req *types.EmailLoginReq) (resp *types.LoginResp, err error) {
	out, err := l.svcCtx.AccountRpc.EmailLogin(l.ctx, &accountrpc.EmailLoginReq{
		Email:       req.Email,
		Password:    req.Password,
		CaptchaKey:  req.CaptchaKey,
		CaptchaCode: req.CaptchaCode,
	})
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.SyslogRpc.AddLoginLog(l.ctx, &syslogrpc.AddLoginLogReq{
		UserId:    out.GetUser().GetUserId(),
		LoginType: out.GetLoginType(),
	})
	if err != nil {
		return nil, err
	}

	return onLogin(l.ctx, l.svcCtx, out)
}
