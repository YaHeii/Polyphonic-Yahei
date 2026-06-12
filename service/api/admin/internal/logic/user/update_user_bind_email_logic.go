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

type UpdateUserBindEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户绑定邮箱
func NewUpdateUserBindEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserBindEmailLogic {
	return &UpdateUserBindEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserBindEmailLogic) UpdateUserBindEmail(req *types.UpdateUserBindEmailReq) (resp *types.EmptyResp, err error) {
	if err := authlogic.VerifyEmailCodeForUser(l.svcCtx, constant.CodeTypeBindEmail, req.Email, req.VerifyCode); err != nil {
		return nil, err
	}

	_, err = l.svcCtx.AccountRpc.BindUserEmail(l.ctx, &accountrpc.BindUserEmailReq{
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
