package socialrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/socialrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTalkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTalkLogic {
	return &GetTalkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询说说
func (l *GetTalkLogic) GetTalk(in *socialrpc.GetTalkReq) (*socialrpc.GetTalkResp, error) {
	record, err := l.svcCtx.TTalkModel.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	commentCount, err := l.svcCtx.TCommentModel.FindCount(l.ctx, "topic_id = ? and type = ? and status != ?", in.Id, talkCommentType(), 2)
	if err != nil {
		return nil, err
	}

	return &socialrpc.GetTalkResp{Talk: convertTalkOut(record, commentCount)}, nil
}
