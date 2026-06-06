package syslogrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindFileLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindFileLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindFileLogListLogic {
	return &FindFileLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询文件记录列表
func (l *FindFileLogListLogic) FindFileLogList(in *syslogrpc.FindFileLogListReq) (*syslogrpc.FindFileLogListResp, error) {
	page, size, sorts, conditions, params := buildFileLogQuery(in)
	records, total, err := l.svcCtx.TFileLogModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	return &syslogrpc.FindFileLogListResp{
		Pagination: buildPageResp(page, size, total),
		List:       convertFileLogListOut(records),
	}, nil
}
