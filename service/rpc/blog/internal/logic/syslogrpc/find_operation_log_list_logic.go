package syslogrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOperationLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOperationLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOperationLogListLogic {
	return &FindOperationLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询操作记录列表
func (l *FindOperationLogListLogic) FindOperationLogList(in *syslogrpc.FindOperationLogListReq) (*syslogrpc.FindOperationLogListResp, error) {
	page, size, sorts, conditions, params := buildOperationLogQuery(in)
	records, total, err := l.svcCtx.TOperationLogModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	return &syslogrpc.FindOperationLogListResp{
		Pagination: buildPageResp(page, size, total),
		List:       convertOperationLogListOut(records),
	}, nil
}
