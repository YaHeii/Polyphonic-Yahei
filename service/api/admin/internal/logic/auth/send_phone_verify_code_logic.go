// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package auth

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendPhoneVerifyCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发送手机验证码
func NewSendPhoneVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendPhoneVerifyCodeLogic {
	return &SendPhoneVerifyCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendPhoneVerifyCodeLogic) SendPhoneVerifyCode(req *types.SendPhoneVerifyCodeReq) (resp *types.EmptyResp, err error) {
	if err := sendPhoneCode(l.svcCtx, req.Phone, req.Type); err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
