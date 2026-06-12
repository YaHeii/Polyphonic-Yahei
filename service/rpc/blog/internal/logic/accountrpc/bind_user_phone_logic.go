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

type BindUserPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindUserPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindUserPhoneLogic {
	return &BindUserPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户登录手机号
func (l *BindUserPhoneLogic) BindUserPhone(in *accountrpc.BindUserPhoneReq) (*accountrpc.BindUserPhoneResp, error) {
	user, err := getCurrentUser(l.ctx, l.svcCtx)
	if err != nil {
		return nil, err
	}
	if !patternx.IsValidPhone(in.Phone) {
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "手机号格式不正确")
	}

	exist, err := l.svcCtx.TUserModel.FindOneByPhone(l.ctx, in.Phone)
	if err == nil && exist != nil && exist.UserId != user.UserId {
		return nil, bizerr.NewBizError(bizcode.CodeUserAlreadyExist, "手机号已被绑定")
	}
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, err
	}

	if err := l.svcCtx.TUserModel.UpdatePhoneByUserID(l.ctx, user.UserId, in.Phone); err != nil {
		return nil, err
	}

	return &accountrpc.BindUserPhoneResp{}, nil
}
