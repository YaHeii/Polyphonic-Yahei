// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package auth

import (
	"context"
	"time"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/syslogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoffLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 注销
func NewLogoffLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoffLogic {
	return &LogoffLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoffLogic) Logoff(req *types.EmptyReq) (resp *types.EmptyResp, err error) {
	uid := currentUserID(l.ctx)

	_, err = l.svcCtx.AccountRpc.Logoff(l.ctx, &accountrpc.LogoffReq{
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.SyslogRpc.AddLogoutLog(l.ctx, &syslogrpc.AddLogoutLogReq{
		UserId:   uid,
		LogoutAt: time.Now().UnixMilli(),
	})
	if err != nil {
		return nil, err
	}

	if err := l.svcCtx.TokenManager.RevokeToken(uid, false); err != nil {
		return nil, err
	}
	if err := l.svcCtx.TokenManager.RevokeToken(uid, true); err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
