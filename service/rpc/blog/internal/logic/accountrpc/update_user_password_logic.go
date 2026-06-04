package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/cryptox"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPasswordLogic {
	return &UpdateUserPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户密码
func (l *UpdateUserPasswordLogic) UpdateUserPassword(in *accountrpc.UpdateUserPasswordReq) (*accountrpc.UpdateUserPasswordResp, error) {
	user, err := getCurrentUser(l.ctx, l.svcCtx)
	if err != nil {
		return nil, err
	}
	if in.OldPassword == "" || in.NewPassword == "" {
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "密码不能为空")
	}
	if !cryptox.BcryptCheck(in.OldPassword, user.Password) {
		return nil, bizerr.NewBizError(bizcode.CodeUserPasswordError, "旧密码错误")
	}

	if err := l.svcCtx.TUserModel.UpdatePasswordByUserID(l.ctx, user.UserId, cryptox.BcryptHash(in.NewPassword)); err != nil {
		return nil, err
	}

	return &accountrpc.UpdateUserPasswordResp{}, nil
}
