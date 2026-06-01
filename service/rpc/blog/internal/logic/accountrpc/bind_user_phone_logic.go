package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindUserPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindUserPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindUserPhoneLogic {
	return &BindUserPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户登录手机号
func (l *BindUserPhoneLogic) BindUserPhone(in *accountrpc.BindUserPhoneReq) (*accountrpc.BindUserPhoneResp, error) {
	// todo: add your logic here and delete this line

	return &accountrpc.BindUserPhoneResp{}, nil
}
