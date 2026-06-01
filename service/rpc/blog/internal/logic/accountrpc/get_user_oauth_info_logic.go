package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserOauthInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserOauthInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserOauthInfoLogic {
	return &GetUserOauthInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户第三平台信息
func (l *GetUserOauthInfoLogic) GetUserOauthInfo(in *accountrpc.GetUserOauthInfoReq) (*accountrpc.GetUserOauthInfoResp, error) {
	// todo: add your logic here and delete this line

	return &accountrpc.GetUserOauthInfoResp{}, nil
}
