package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminResetUserPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminResetUserPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminResetUserPasswordLogic {
	return &AdminResetUserPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 管理员重置用户密码
func (l *AdminResetUserPasswordLogic) AdminResetUserPassword(in *accountrpc.AdminResetUserPasswordReq) (*accountrpc.AdminResetUserPasswordResp, error) {
	// todo: add your logic here and delete this line

	return &accountrpc.AdminResetUserPasswordResp{}, nil
}
