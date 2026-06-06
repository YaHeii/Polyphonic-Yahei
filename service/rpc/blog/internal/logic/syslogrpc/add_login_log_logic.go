package syslogrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLoginLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLoginLogLogic {
	return &AddLoginLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建登录记录
func (l *AddLoginLogLogic) AddLoginLog(in *syslogrpc.AddLoginLogReq) (*syslogrpc.AddLoginLogResp, error) {
	entity := convertAddLoginLogIn(l.ctx, in)
	if _, err := l.svcCtx.TLoginLogModel.Save(l.ctx, entity); err != nil {
		return nil, err
	}

	record, err := l.svcCtx.TLoginLogModel.FindById(l.ctx, entity.Id)
	if err != nil {
		return nil, err
	}

	return &syslogrpc.AddLoginLogResp{
		LoginLog: convertLoginLogOut(record),
	}, nil
}
