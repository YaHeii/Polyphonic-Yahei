package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnalysisUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnalysisUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnalysisUserLogic {
	return &AnalysisUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户数量
func (l *AnalysisUserLogic) AnalysisUser(in *accountrpc.AnalysisUserReq) (*accountrpc.AnalysisUserResp, error) {
	tableName := `"public"."t_user"`
	if in.UserType == 1 {
		tableName = `"public"."t_visitor"`
	}

	var result struct {
		Count int64 `db:"count"`
	}
	if err := l.svcCtx.SqlConn.QueryRowCtx(l.ctx, &result, `select count(*) as count from `+tableName); err != nil {
		return nil, err
	}

	return &accountrpc.AnalysisUserResp{
		UserCount: result.Count,
	}, nil
}
