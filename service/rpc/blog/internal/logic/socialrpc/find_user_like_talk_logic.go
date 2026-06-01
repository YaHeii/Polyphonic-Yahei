package socialrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/socialrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLikeTalkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLikeTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLikeTalkLogic {
	return &FindUserLikeTalkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户点赞的说说
func (l *FindUserLikeTalkLogic) FindUserLikeTalk(in *socialrpc.FindUserLikeTalkReq) (*socialrpc.FindUserLikeTalkResp, error) {
	// todo: add your logic here and delete this line

	return &socialrpc.FindUserLikeTalkResp{}, nil
}
