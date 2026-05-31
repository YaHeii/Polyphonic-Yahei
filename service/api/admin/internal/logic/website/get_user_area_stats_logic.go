// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package website

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAreaStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户分布地区
func NewGetUserAreaStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAreaStatsLogic {
	return &GetUserAreaStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserAreaStatsLogic) GetUserAreaStats(req *types.GetUserAreaStatsReq) (resp *types.GetUserAreaStatsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
