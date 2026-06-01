package syslogrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogoutLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogoutLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogoutLogLogic {
	return &AddLogoutLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新登录记录
func (l *AddLogoutLogLogic) AddLogoutLog(in *syslogrpc.AddLogoutLogReq) (*syslogrpc.AddLogoutLogResp, error) {
	// todo: add your logic here and delete this line

	return &syslogrpc.AddLogoutLogResp{}, nil
}
