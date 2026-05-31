// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserBindThirdPartyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户绑定第三方平台账号
func NewUpdateUserBindThirdPartyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserBindThirdPartyLogic {
	return &UpdateUserBindThirdPartyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserBindThirdPartyLogic) UpdateUserBindThirdParty(req *types.UpdateUserBindThirdPartyReq) (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
