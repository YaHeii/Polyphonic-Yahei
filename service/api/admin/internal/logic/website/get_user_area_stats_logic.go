// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package website

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"

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
	out, err := l.svcCtx.AccountRpc.AnalysisUserAreas(l.ctx, &accountrpc.AnalysisUserAreasReq{
		UserType: req.UserType,
	})
	if err != nil {
		return nil, err
	}

	list := make([]*types.UserAreaVO, 0, len(out.List))
	for _, item := range out.List {
		list = append(list, &types.UserAreaVO{
			Name:  item.Area,
			Value: item.Count,
		})
	}

	return &types.GetUserAreaStatsResp{
		UserAreas:    list,
		TouristAreas: list,
	}, nil
}
