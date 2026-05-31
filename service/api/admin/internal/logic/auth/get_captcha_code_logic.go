// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package auth

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
