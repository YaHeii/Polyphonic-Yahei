package noticerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/noticerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserNoticeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserNoticeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserNoticeListLogic {
	return &FindUserNoticeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户可见通知列表
func (l *FindUserNoticeListLogic) FindUserNoticeList(in *noticerpc.FindUserNoticeListReq) (*noticerpc.FindUserNoticeListResp, error) {
	// todo: add your logic here and delete this line

	return &noticerpc.FindUserNoticeListResp{}, nil
}
