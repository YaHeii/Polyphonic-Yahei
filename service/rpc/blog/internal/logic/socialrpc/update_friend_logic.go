package socialrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/socialrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFriendLogic {
	return &UpdateFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新友链
func (l *UpdateFriendLogic) UpdateFriend(in *socialrpc.UpdateFriendReq) (*socialrpc.UpdateFriendResp, error) {
	entity := convertUpdateFriendIn(in)
	if _, err := l.svcCtx.TFriendModel.Save(l.ctx, entity); err != nil {
		return nil, err
	}

	record, err := l.svcCtx.TFriendModel.FindById(l.ctx, entity.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}

	return &socialrpc.UpdateFriendResp{Friend: convertFriendOut(record)}, nil
}
