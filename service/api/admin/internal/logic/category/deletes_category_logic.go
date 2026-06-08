// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package category

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/articlerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除文章分类
func NewDeletesCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesCategoryLogic {
	return &DeletesCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletesCategoryLogic) DeletesCategory(req *types.IdsReq) (resp *types.BatchResp, err error) {
	in := &articlerpc.DeletesCategoryReq{
		Ids: req.Ids,
	}

	out, err := l.svcCtx.ArticleRpc.DeletesCategory(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}
	return resp, nil
}
