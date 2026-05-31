// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package account

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAccountStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户状态
func NewUpdateAccountStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAccountStatusLogic {
	return &UpdateAccountStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAccountStatusLogic) UpdateAccountStatus(req *types.UpdateAccountStatusReq) (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
