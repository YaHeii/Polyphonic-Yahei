package noticerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/noticerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNoticeStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNoticeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNoticeStatusLogic {
	return &UpdateNoticeStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新通知状态
func (l *UpdateNoticeStatusLogic) UpdateNoticeStatus(in *noticerpc.UpdateNoticeStatusReq) (*noticerpc.UpdateNoticeStatusResp, error) {
	// todo: add your logic here and delete this line

	return &noticerpc.UpdateNoticeStatusResp{}, nil
}
