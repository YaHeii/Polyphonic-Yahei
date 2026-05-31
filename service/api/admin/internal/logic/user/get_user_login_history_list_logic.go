// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLoginHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询用户登录历史
func NewGetUserLoginHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLoginHistoryListLogic {
	return &GetUserLoginHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLoginHistoryListLogic) GetUserLoginHistoryList(req *types.QueryUserLoginHistoryReq) (resp *types.PageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
