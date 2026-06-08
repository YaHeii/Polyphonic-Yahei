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

type DeletesNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除通知
func NewDeletesNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesNoticeLogic {
	return &DeletesNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletesNoticeLogic) DeletesNotice(req *types.IdsReq) (resp *types.BatchResp, err error) {
	in := &noticerpc.DeletesNoticeReq{
		Ids: req.Ids,
	}

	out, err := l.svcCtx.NoticeRpc.DeletesNotice(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{SuccessCount: out.SuccessCount}, nil
}
