// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package account

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAccountPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户密码
func NewUpdateAccountPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAccountPasswordLogic {
	return &UpdateAccountPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAccountPasswordLogic) UpdateAccountPassword(req *types.UpdateAccountPasswordReq) (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
