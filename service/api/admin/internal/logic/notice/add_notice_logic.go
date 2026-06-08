// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package notice

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/noticerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建通知
func NewAddNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddNoticeLogic {
	return &AddNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddNoticeLogic) AddNotice(req *types.AddNoticeReq) (resp *types.NoticeBackVO, err error) {
	in := &noticerpc.AddNoticeReq{
		Title:   req.Title,
		Content: req.Content,
		Type:    req.Type,
		Level:   req.Level,
		AppName: req.AppName,
	}

	out, err := l.svcCtx.NoticeRpc.AddNotice(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convertNoticeOut(out.Notice), nil
}
