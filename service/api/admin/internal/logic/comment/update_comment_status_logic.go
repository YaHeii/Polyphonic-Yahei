// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package comment

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/newsrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新评论状态
func NewUpdateCommentStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentStatusLogic {
	return &UpdateCommentStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCommentStatusLogic) UpdateCommentStatus(req *types.UpdateCommentStatusReq) (resp *types.BatchResp, err error) {
	out, err := l.svcCtx.NewsRpc.UpdateCommentStatus(l.ctx, &newsrpc.UpdateCommentStatusReq{
		Ids:    req.Ids,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{SuccessCount: out.SuccessCount}, nil
}
