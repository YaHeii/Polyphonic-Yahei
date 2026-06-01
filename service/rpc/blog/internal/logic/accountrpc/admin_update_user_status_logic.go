package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminUpdateUserStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminUpdateUserStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminUpdateUserStatusLogic {
	return &AdminUpdateUserStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户状态
func (l *AdminUpdateUserStatusLogic) AdminUpdateUserStatus(in *accountrpc.AdminUpdateUserStatusReq) (*accountrpc.AdminUpdateUserStatusResp, error) {
	// todo: add your logic here and delete this line

	return &accountrpc.AdminUpdateUserStatusResp{}, nil
}
