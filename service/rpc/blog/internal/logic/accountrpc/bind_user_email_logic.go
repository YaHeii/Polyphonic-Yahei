package accountrpclogic

import (
	"context"
	"errors"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/patternx"
	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindUserEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindUserEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindUserEmailLogic {
	return &BindUserEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户登录邮箱
func (l *BindUserEmailLogic) BindUserEmail(in *accountrpc.BindUserEmailReq) (*accountrpc.BindUserEmailResp, error) {
	user, err := getCurrentUser(l.ctx, l.svcCtx)
	if err != nil {
		return nil, err
	}
	if !patternx.IsValidEmail(in.Email) {
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "邮箱格式不正确")
	}

	exist, err := l.svcCtx.TUserModel.FindOneByEmail(l.ctx, in.Email)
	if err == nil && exist != nil && exist.UserId != user.UserId {
		return nil, bizerr.NewBizError(bizcode.CodeUserAlreadyExist, "邮箱已被绑定")
	}
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, err
	}

	if err := l.svcCtx.TUserModel.UpdateEmailByUserID(l.ctx, user.UserId, in.Email); err != nil {
		return nil, err
	}

	return &accountrpc.BindUserEmailResp{}, nil
}
