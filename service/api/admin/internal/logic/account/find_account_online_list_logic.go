// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package account

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAccountOnlineListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询在线用户列表
func NewFindAccountOnlineListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAccountOnlineListLogic {
	return &FindAccountOnlineListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindAccountOnlineListLogic) FindAccountOnlineList(req *types.QueryAccountReq) (resp *types.PageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
