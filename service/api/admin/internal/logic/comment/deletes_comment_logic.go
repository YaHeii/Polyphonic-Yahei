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

type DeletesCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除评论
func NewDeletesCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesCommentLogic {
	return &DeletesCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletesCommentLogic) DeletesComment(req *types.IdsReq) (resp *types.BatchResp, err error) {
	out, err := l.svcCtx.NewsRpc.DeletesComment(l.ctx, &newsrpc.DeletesCommentReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{SuccessCount: out.SuccessCount}, nil
}
