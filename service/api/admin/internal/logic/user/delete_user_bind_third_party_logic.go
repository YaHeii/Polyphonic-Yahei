// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserBindThirdPartyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除用户绑定第三方平台账号
func NewDeleteUserBindThirdPartyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserBindThirdPartyLogic {
	return &DeleteUserBindThirdPartyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserBindThirdPartyLogic) DeleteUserBindThirdParty(req *types.DeleteUserBindThirdPartyReq) (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
