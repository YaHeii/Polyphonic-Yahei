// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package website

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVisitStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取访客数据分析
func NewGetVisitStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVisitStatsLogic {
	return &GetVisitStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVisitStatsLogic) GetVisitStats(req *types.EmptyReq) (resp *types.GetVisitStatsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
