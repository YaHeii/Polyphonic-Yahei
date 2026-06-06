package socialrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/socialrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTalkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTalkLogic {
	return &UpdateTalkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新说说
func (l *UpdateTalkLogic) UpdateTalk(in *socialrpc.UpdateTalkReq) (*socialrpc.UpdateTalkResp, error) {
	entity := convertUpdateTalkIn(in)
	if _, err := l.svcCtx.TTalkModel.Save(l.ctx, entity); err != nil {
		return nil, err
	}

	record, err := l.svcCtx.TTalkModel.FindById(l.ctx, entity.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}

	commentCount, err := l.svcCtx.TCommentModel.FindCount(l.ctx, "topic_id = ? and type = ? and status != ?", entity.Id, talkCommentType(), 2)
	if err != nil {
		return nil, err
	}

	return &socialrpc.UpdateTalkResp{Talk: convertTalkOut(record, commentCount)}, nil
}
