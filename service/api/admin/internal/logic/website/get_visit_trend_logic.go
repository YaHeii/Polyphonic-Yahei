// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package website

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVisitTrendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取访客数据趋势
func NewGetVisitTrendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVisitTrendLogic {
	return &GetVisitTrendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVisitTrendLogic) GetVisitTrend(req *types.GetVisitTrendReq) (resp *types.GetVisitTrendResp, err error) {
	// todo: add your logic here and delete this line

	return
}
