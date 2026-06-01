package newsrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/newsrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMessageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMessageListLogic {
	return &FindMessageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询留言列表
func (l *FindMessageListLogic) FindMessageList(in *newsrpc.FindMessageListReq) (*newsrpc.FindMessageListResp, error) {
	// todo: add your logic here and delete this line

	return &newsrpc.FindMessageListResp{}, nil
}
