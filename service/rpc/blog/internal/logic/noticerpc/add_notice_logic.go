package noticerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/common/rpcutils"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/noticerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddNoticeLogic {
	return &AddNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建通知
func (l *AddNoticeLogic) AddNotice(in *noticerpc.AddNoticeReq) (*noticerpc.AddNoticeResp, error) {
	publisherID, _ := rpcutils.GetUserIdFromCtx(l.ctx)

	entity := &model.TSystemNotice{
		Id:            in.Id,
		Title:         in.Title,
		Content:       in.Content,
		Type:          in.Type,
		Level:         in.Level,
		AppName:       in.AppName,
		PublisherId:   publisherID,
		PublishStatus: noticeStatusDraft,
	}

	_, err := l.svcCtx.TSystemNoticeModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return &noticerpc.AddNoticeResp{
		Notice: convertNoticeOut(entity),
	}, nil
}
