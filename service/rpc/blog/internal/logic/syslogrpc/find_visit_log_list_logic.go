package syslogrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindVisitLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindVisitLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindVisitLogListLogic {
	return &FindVisitLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询操作访问列表
func (l *FindVisitLogListLogic) FindVisitLogList(in *syslogrpc.FindVisitLogListReq) (*syslogrpc.FindVisitLogListResp, error) {
	page, size, sorts, conditions, params := buildVisitLogQuery(in)
	records, total, err := l.svcCtx.TVisitLogModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	return &syslogrpc.FindVisitLogListResp{
		Pagination: buildPageResp(page, size, total),
		List:       convertVisitLogListOut(records),
	}, nil
}
