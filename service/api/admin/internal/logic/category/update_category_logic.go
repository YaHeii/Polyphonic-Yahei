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

type UpdateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新文章分类
func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(req *types.NewCategoryReq) (resp *types.CategoryBackVO, err error) {
	in := &articlerpc.UpdateCategoryReq{
		Id:           req.Id,
		CategoryName: req.CategoryName,
	}

	out, err := l.svcCtx.ArticleRpc.UpdateCategory(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convertCategoryTypes(out.Category), nil
}
