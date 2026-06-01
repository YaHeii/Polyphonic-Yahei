package noticerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/noticerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesNoticeLogic {
	return &DeletesNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除通知
func (l *DeletesNoticeLogic) DeletesNotice(in *noticerpc.DeletesNoticeReq) (*noticerpc.DeletesNoticeResp, error) {
	// todo: add your logic here and delete this line

	return &noticerpc.DeletesNoticeResp{}, nil
}
