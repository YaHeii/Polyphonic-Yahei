package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThirdLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThirdLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThirdLoginLogic {
	return &ThirdLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 第三方登录
func (l *ThirdLoginLogic) ThirdLogin(in *accountrpc.ThirdLoginReq) (*accountrpc.LoginResp, error) {
	// todo: add your logic here and delete this line

	return &accountrpc.LoginResp{}, nil
}
