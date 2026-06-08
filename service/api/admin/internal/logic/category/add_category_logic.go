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

type AddCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建文章分类
func NewAddCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCategoryLogic {
	return &AddCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCategoryLogic) AddCategory(req *types.NewCategoryReq) (resp *types.CategoryBackVO, err error) {
	in := &articlerpc.AddCategoryReq{
		Id:           req.Id,
		CategoryName: req.CategoryName,
	}

	out, err := l.svcCtx.ArticleRpc.AddCategory(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convertCategoryTypes(out.Category), nil
}
