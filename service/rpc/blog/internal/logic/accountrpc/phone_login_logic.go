package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/common/enums"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/patternx"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPhoneLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneLoginLogic {
	return &PhoneLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 手机号登录
func (l *PhoneLoginLogic) PhoneLogin(in *accountrpc.PhoneLoginReq) (*accountrpc.LoginResp, error) {
	// 校验参数
	if !patternx.IsValidPhone(in.Phone) {
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "手机号格式不正确")
	}

	// 验证用户是否存在
	account, err := l.svcCtx.TUserModel.FindOneByPhone(l.ctx, in.Phone)
	if err != nil {
		return nil, bizerr.NewBizError(bizcode.CodeUserNotExist, "手机号未绑定邮箱账号")
	}

	return onLogin(l.ctx, l.svcCtx, account, enums.LoginTypePhone)
}
