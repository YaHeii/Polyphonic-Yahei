// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package auth

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

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
	if l.svcCtx.CaptchaHolder == nil {
		return nil, bizerr.NewBizError(bizcode.CodeInternalServerError, "验证码服务未初始化")
	}

	key, base64, code, err := l.svcCtx.CaptchaHolder.GetMathImageCaptcha(int(req.Height), int(req.Width))
	if err != nil {
		return nil, err
	}

	return &types.GetCaptchaCodeResp{
		CaptchaKey:    key,
		CaptchaBase64: base64,
		CaptchaCode:   code,
	}, nil
}
