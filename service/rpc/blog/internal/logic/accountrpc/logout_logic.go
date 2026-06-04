package accountrpclogic

import (
	"context"
	"time"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登出
func (l *LogoutLogic) Logout(in *accountrpc.LogoutReq) (*accountrpc.LogoutResp, error) {
	if in.UserId == "" {
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "用户id不能为空")
	}

	if err := l.svcCtx.OnlineUserService.Logout(l.ctx, in.UserId); err != nil {
		return nil, err
	}

	return &accountrpc.LogoutResp{
		LogoutAt: time.Now().Unix(),
	}, nil
}
