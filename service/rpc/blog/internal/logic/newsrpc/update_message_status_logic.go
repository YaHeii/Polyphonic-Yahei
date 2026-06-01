package newsrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/newsrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMessageStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMessageStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMessageStatusLogic {
	return &UpdateMessageStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新留言状态
func (l *UpdateMessageStatusLogic) UpdateMessageStatus(in *newsrpc.UpdateMessageStatusReq) (*newsrpc.UpdateMessageStatusResp, error) {
	// todo: add your logic here and delete this line

	return &newsrpc.UpdateMessageStatusResp{}, nil
}
