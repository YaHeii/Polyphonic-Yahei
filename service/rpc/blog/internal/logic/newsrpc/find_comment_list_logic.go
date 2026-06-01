package newsrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/newsrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCommentListLogic {
	return &FindCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询评论列表
func (l *FindCommentListLogic) FindCommentList(in *newsrpc.FindCommentListReq) (*newsrpc.FindCommentListResp, error) {
	// todo: add your logic here and delete this line

	return &newsrpc.FindCommentListResp{}, nil
}
