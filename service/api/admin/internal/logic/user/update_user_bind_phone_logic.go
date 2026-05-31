// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
