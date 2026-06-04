package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type AnalysisUserAreasLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type userAreaRow struct {
	IpSource string `db:"ip_source"`
	Count    int64  `db:"count"`
}

func NewAnalysisUserAreasLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnalysisUserAreasLogic {
	return &AnalysisUserAreasLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户分布区域
func (l *AnalysisUserAreasLogic) AnalysisUserAreas(in *accountrpc.AnalysisUserAreasReq) (*accountrpc.AnalysisUserAreasResp, error) {
	tableName := `"public"."t_user"`
	if in.UserType == 1 {
		tableName = `"public"."t_visitor"`
	}

	query := `
select ip_source, count(*) as count
from ` + tableName + `
group by ip_source
order by count desc`

	var result []userAreaRow
	if err := l.svcCtx.SqlConn.QueryRowsCtx(l.ctx, &result, query); err != nil {
		return nil, err
	}

	areas := make(map[string]int64)
	for _, item := range result {
		if item.IpSource == "" {
			continue
		}

		area := findArea(item.IpSource)
		areas[area] += item.Count
	}

	list := make([]*accountrpc.UserArea, 0, len(areas))
	for k, v := range areas {
		list = append(list, &accountrpc.UserArea{
			Area:  k,
			Count: v,
		})
	}

	return &accountrpc.AnalysisUserAreasResp{
		List: list,
	}, nil
}
