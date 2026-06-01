package syslogrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindLoginLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindLoginLogListLogic {
	return &FindLoginLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询登录记录列表
func (l *FindLoginLogListLogic) FindLoginLogList(in *syslogrpc.FindLoginLogListReq) (*syslogrpc.FindLoginLogListResp, error) {
	// todo: add your logic here and delete this line

	return &syslogrpc.FindLoginLogListResp{}, nil
}
