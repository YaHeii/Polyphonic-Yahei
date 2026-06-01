package socialrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/socialrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeTalkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeTalkLogic {
	return &LikeTalkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 点赞说说
func (l *LikeTalkLogic) LikeTalk(in *socialrpc.LikeTalkReq) (*socialrpc.LikeTalkResp, error) {
	// todo: add your logic here and delete this line

	return &socialrpc.LikeTalkResp{}, nil
}
