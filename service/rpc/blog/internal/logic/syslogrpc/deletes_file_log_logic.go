package syslogrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesFileLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesFileLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesFileLogLogic {
	return &DeletesFileLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除文件记录
func (l *DeletesFileLogLogic) DeletesFileLog(in *syslogrpc.DeletesFileLogReq) (*syslogrpc.DeletesFileLogResp, error) {
	// todo: add your logic here and delete this line

	return &syslogrpc.DeletesFileLogResp{}, nil
}
