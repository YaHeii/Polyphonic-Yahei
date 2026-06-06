package syslogrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOperationLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOperationLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOperationLogLogic {
	return &AddOperationLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建操作记录
func (l *AddOperationLogLogic) AddOperationLog(in *syslogrpc.AddOperationLogReq) (*syslogrpc.AddOperationLogResp, error) {
	entity := convertAddOperationLogIn(l.ctx, in)
	if _, err := l.svcCtx.TOperationLogModel.Save(l.ctx, entity); err != nil {
		return nil, err
	}

	record, err := l.svcCtx.TOperationLogModel.FindById(l.ctx, entity.Id)
	if err != nil {
		return nil, err
	}

	return &syslogrpc.AddOperationLogResp{
		OperationLog: convertOperationLogOut(record),
	}, nil
}
