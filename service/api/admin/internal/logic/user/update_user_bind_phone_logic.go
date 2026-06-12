// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/common/constant"
	authlogic "github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/logic/auth"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserBindPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户绑定手机号
func NewUpdateUserBindPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserBindPhoneLogic {
	return &UpdateUserBindPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserBindPhoneLogic) UpdateUserBindPhone(req *types.UpdateUserBindPhoneReq) (resp *types.EmptyResp, err error) {
	if err := authlogic.VerifyPhoneCodeForUser(l.svcCtx, constant.CodeTypeBindPhone, req.Phone, req.VerifyCode); err != nil {
		return nil, err
	}

	_, err = l.svcCtx.AccountRpc.BindUserPhone(l.ctx, &accountrpc.BindUserPhoneReq{
		Phone: req.Phone,
	})
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
