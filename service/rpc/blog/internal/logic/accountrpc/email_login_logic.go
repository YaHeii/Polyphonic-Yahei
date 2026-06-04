package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/common/enums"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/cryptox"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/patternx"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmailLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailLoginLogic {
	return &EmailLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 邮箱登录
func (l *EmailLoginLogic) EmailLogin(in *accountrpc.EmailLoginReq) (*accountrpc.LoginResp, error) {
	if !patternx.IsValidEmail(in.Email) {
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "邮箱格式不正确")
	}

	if !l.svcCtx.CaptchaHolder.VerifyCaptcha(in.CaptchaKey, in.CaptchaCode) {
		return nil, bizerr.NewBizError(bizcode.CodeCaptchaVerify, "验证码错误")
	}

	account, err := l.svcCtx.TUserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, bizerr.NewBizError(bizcode.CodeUserNotExist, "邮箱未注册")
	}

	if !cryptox.BcryptCheck(in.Password, account.Password) {
		return nil, bizerr.NewBizError(bizcode.CodeUserPasswordError, "密码不正确")
	}

	return onLogin(l.ctx, l.svcCtx, account, enums.LoginTypeEmail)
}
