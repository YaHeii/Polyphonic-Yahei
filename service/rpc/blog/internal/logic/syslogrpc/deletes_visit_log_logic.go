package syslogrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesVisitLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesVisitLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesVisitLogLogic {
	return &DeletesVisitLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除访问记录
func (l *DeletesVisitLogLogic) DeletesVisitLog(in *syslogrpc.DeletesVisitLogReq) (*syslogrpc.DeletesVisitLogResp, error) {
	// todo: add your logic here and delete this line

	return &syslogrpc.DeletesVisitLogResp{}, nil
}
