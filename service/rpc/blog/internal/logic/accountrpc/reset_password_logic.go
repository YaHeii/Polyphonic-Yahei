package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/cryptox"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/patternx"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 重置密码
func (l *ResetPasswordLogic) ResetPassword(in *accountrpc.ResetPasswordReq) (*accountrpc.ResetPasswordResp, error) {
	// 校验邮箱格式
	if !patternx.IsValidEmail(in.Email) {
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "邮箱格式不正确")
	}

	// 验证用户是否存在
	exist, _ := l.svcCtx.TUserModel.FindOneByEmail(l.ctx, in.Email)
	if exist == nil {
		return nil, bizerr.NewBizError(bizcode.CodeUserNotExist, "用户不存在")
	}

	// 更新密码
	exist.Password = cryptox.BcryptHash(in.Password)

	if err := l.svcCtx.TUserModel.Update(l.ctx, exist); err != nil {
		return nil, err
	}

	return &accountrpc.ResetPasswordResp{}, nil
}
