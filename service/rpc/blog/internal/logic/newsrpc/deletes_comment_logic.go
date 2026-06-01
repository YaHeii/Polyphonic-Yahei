package newsrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/newsrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesCommentLogic {
	return &DeletesCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除评论
func (l *DeletesCommentLogic) DeletesComment(in *newsrpc.DeletesCommentReq) (*newsrpc.DeletesCommentResp, error) {
	// todo: add your logic here and delete this line

	return &newsrpc.DeletesCommentResp{}, nil
}
