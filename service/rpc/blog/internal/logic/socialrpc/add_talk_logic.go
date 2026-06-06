package socialrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/socialrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTalkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTalkLogic {
	return &AddTalkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建说说
func (l *AddTalkLogic) AddTalk(in *socialrpc.AddTalkReq) (*socialrpc.AddTalkResp, error) {
	entity := convertAddTalkIn(in)
	if _, err := l.svcCtx.TTalkModel.Save(l.ctx, entity); err != nil {
		return nil, err
	}

	record, err := l.svcCtx.TTalkModel.FindById(l.ctx, entity.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}

	return &socialrpc.AddTalkResp{Talk: convertTalkOut(record, 0)}, nil
}
