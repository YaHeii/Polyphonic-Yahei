package newsrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/newsrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMessageLogic {
	return &AddMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建留言
func (l *AddMessageLogic) AddMessage(in *newsrpc.AddMessageReq) (*newsrpc.AddMessageResp, error) {
	// todo: add your logic here and delete this line

	return &newsrpc.AddMessageResp{}, nil
}
