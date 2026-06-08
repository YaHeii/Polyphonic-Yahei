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

type PhoneLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 手机登录
func NewPhoneLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneLoginLogic {
	return &PhoneLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhoneLoginLogic) PhoneLogin(req *types.PhoneLoginReq) (resp *types.LoginResp, err error) {
	out, err := l.svcCtx.AccountRpc.PhoneLogin(l.ctx, &accountrpc.PhoneLoginReq{
		Phone:      req.Phone,
		VerifyCode: req.VerifyCode,
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
