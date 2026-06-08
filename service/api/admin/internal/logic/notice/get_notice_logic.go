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

type GetNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询通知详情
func NewGetNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNoticeLogic {
	return &GetNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNoticeLogic) GetNotice(req *types.IdReq) (resp *types.NoticeBackVO, err error) {
	in := &noticerpc.GetNoticeReq{Id: req.Id}

	out, err := l.svcCtx.NoticeRpc.GetNotice(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convertNoticeOut(out.Notice), nil
}
