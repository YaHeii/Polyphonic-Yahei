package syslogrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesLoginLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesLoginLogLogic {
	return &DeletesLoginLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除登录记录
func (l *DeletesLoginLogLogic) DeletesLoginLog(in *syslogrpc.DeletesLoginLogReq) (*syslogrpc.DeletesLoginLogResp, error) {
	// todo: add your logic here and delete this line

	return &syslogrpc.DeletesLoginLogResp{}, nil
}
