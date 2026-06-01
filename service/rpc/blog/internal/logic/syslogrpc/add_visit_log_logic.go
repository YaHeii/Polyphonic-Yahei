package syslogrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVisitLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddVisitLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVisitLogLogic {
	return &AddVisitLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建访问记录
func (l *AddVisitLogLogic) AddVisitLog(in *syslogrpc.AddVisitLogReq) (*syslogrpc.AddVisitLogResp, error) {
	// todo: add your logic here and delete this line

	return &syslogrpc.AddVisitLogResp{}, nil
}
