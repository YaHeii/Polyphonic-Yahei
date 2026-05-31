// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package account

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAccountListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询用户列表
func NewFindAccountListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAccountListLogic {
	return &FindAccountListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindAccountListLogic) FindAccountList(req *types.QueryAccountReq) (resp *types.PageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
