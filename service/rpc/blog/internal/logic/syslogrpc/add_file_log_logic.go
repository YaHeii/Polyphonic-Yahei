package syslogrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFileLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFileLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFileLogLogic {
	return &AddFileLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建文件记录
func (l *AddFileLogLogic) AddFileLog(in *syslogrpc.AddFileLogReq) (*syslogrpc.AddFileLogResp, error) {
	// todo: add your logic here and delete this line

	return &syslogrpc.AddFileLogResp{}, nil
}
