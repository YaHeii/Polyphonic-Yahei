// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package message

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/newsrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除留言
func NewDeletesMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesMessageLogic {
	return &DeletesMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletesMessageLogic) DeletesMessage(req *types.IdsReq) (resp *types.BatchResp, err error) {
	in := &newsrpc.DeletesMessageReq{
		Ids: req.Ids,
	}

	out, err := l.svcCtx.NewsRpc.DeletesMessage(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{SuccessCount: out.SuccessCount}, nil
}
