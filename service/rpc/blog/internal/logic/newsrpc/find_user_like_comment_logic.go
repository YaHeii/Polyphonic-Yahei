package newsrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/newsrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLikeCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLikeCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLikeCommentLogic {
	return &FindUserLikeCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户点赞的评论
func (l *FindUserLikeCommentLogic) FindUserLikeComment(in *newsrpc.FindUserLikeCommentReq) (*newsrpc.FindLikeCommentResp, error) {
	// todo: add your logic here and delete this line

	return &newsrpc.FindLikeCommentResp{}, nil
}
