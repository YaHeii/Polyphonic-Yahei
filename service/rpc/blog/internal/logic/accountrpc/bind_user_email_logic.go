package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindUserEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindUserEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindUserEmailLogic {
	return &BindUserEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户登录邮箱
func (l *BindUserEmailLogic) BindUserEmail(in *accountrpc.BindUserEmailReq) (*accountrpc.BindUserEmailResp, error) {
	// todo: add your logic here and delete this line

	return &accountrpc.BindUserEmailResp{}, nil
}
