package socialrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/socialrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建友链
func (l *AddFriendLogic) AddFriend(in *socialrpc.AddFriendReq) (*socialrpc.AddFriendResp, error) {
	entity := convertAddFriendIn(in)
	if _, err := l.svcCtx.TFriendModel.Save(l.ctx, entity); err != nil {
		return nil, err
	}

	record, err := l.svcCtx.TFriendModel.FindById(l.ctx, entity.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}

	return &socialrpc.AddFriendResp{Friend: convertFriendOut(record)}, nil
}
