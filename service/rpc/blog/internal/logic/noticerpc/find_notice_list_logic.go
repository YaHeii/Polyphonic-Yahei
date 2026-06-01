package noticerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/noticerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindNoticeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindNoticeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindNoticeListLogic {
	return &FindNoticeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询通知列表
func (l *FindNoticeListLogic) FindNoticeList(in *noticerpc.FindNoticeListReq) (*noticerpc.FindNoticeListResp, error) {
	// todo: add your logic here and delete this line

	return &noticerpc.FindNoticeListResp{}, nil
}
