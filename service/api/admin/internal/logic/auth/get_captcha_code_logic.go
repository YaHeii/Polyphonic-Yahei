// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package auth

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取验证码
func NewGetCaptchaCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaCodeLogic {
	return &GetCaptchaCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaCodeLogic) GetCaptchaCode(req *types.GetCaptchaCodeReq) (resp *types.GetCaptchaCodeResp, err error) {
	out, err := l.svcCtx.AccountRpc.GenerateCaptchaCode(l.ctx, &accountrpc.GenerateCaptchaCodeReq{
		Width:  req.Width,
		Height: req.Height,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetCaptchaCodeResp{
		CaptchaKey:    out.GetCaptchaKey(),
		CaptchaBase64: out.GetCaptchaBase64(),
		CaptchaCode:   out.GetCaptchaCode(),
	}, nil
}
