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
	userID := in.UserId
	if userID == "" {
		userID = metadataUserID(l.ctx)
	}

	if _, err := l.svcCtx.TLoginLogModel.UpdateLatestLogout(l.ctx, userID, logoutTime(in)); err != nil {
		return nil, err
	}

	return &syslogrpc.AddLogoutLogResp{}, nil
}
