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
	// todo: add your logic here and delete this line

	return &syslogrpc.AddOperationLogResp{}, nil
}
