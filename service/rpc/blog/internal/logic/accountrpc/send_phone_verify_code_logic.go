package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/common/constant"
	"github.com/YaHeii/Polyphonic-Yahei/common/rediskey"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/patternx"
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
	if !patternx.IsValidPhone(in.Phone) {
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "手机号格式不正确")
	}

	switch in.Type {
	case constant.CodeTypeRegister, constant.CodeTypeBindPhone:
		exist, _ := l.svcCtx.TUserModel.FindOneByPhone(l.ctx, in.Phone)
		if exist != nil {
			return nil, bizerr.NewBizError(bizcode.CodeUserAlreadyExist, "手机号已被绑定")
		}
	case constant.CodeTypeResetPwd:
		exist, _ := l.svcCtx.TUserModel.FindOneByPhone(l.ctx, in.Phone)
		if exist == nil {
			return nil, bizerr.NewBizError(bizcode.CodeUserNotExist, "手机号未注册")
		}
	default:
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "验证码类型不正确")
	}

	key := rediskey.GetCaptchaKey(in.Type, in.Phone)
	if _, err := l.svcCtx.CaptchaHolder.GetCodeCaptcha(key); err != nil {
		return nil, err
	}

	return &accountrpc.SendPhoneVerifyCodeResp{}, nil
}
