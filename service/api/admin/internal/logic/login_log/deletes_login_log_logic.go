// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package login_log

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesLoginLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除登录日志
func NewDeletesLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesLoginLogLogic {
	return &DeletesLoginLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletesLoginLogLogic) DeletesLoginLog(req *types.IdsReq) (resp *types.BatchResp, err error) {
	// todo: add your logic here and delete this line

	return
}
