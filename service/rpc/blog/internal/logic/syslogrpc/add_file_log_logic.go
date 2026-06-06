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
	entity := convertAddFileLogIn(l.ctx, in)
	if _, err := l.svcCtx.TFileLogModel.Save(l.ctx, entity); err != nil {
		return nil, err
	}

	record, err := l.svcCtx.TFileLogModel.FindById(l.ctx, entity.Id)
	if err != nil {
		return nil, err
	}

	return &syslogrpc.AddFileLogResp{
		FileLog: convertFileLogOut(record),
	}, nil
}
