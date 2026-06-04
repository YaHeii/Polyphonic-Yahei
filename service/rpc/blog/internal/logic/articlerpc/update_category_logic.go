package articlerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新文章分类
func (l *UpdateCategoryLogic) UpdateCategory(in *articlerpc.UpdateCategoryReq) (*articlerpc.UpdateCategoryResp, error) {
	entity, err := l.svcCtx.TCategoryModel.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	entity.CategoryName = in.CategoryName
	_, err = l.svcCtx.TCategoryModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return &articlerpc.UpdateCategoryResp{
		Category: &articlerpc.Category{
			Id:           entity.Id,
			CategoryName: entity.CategoryName,
			CreatedAt:    entity.CreatedAt.UnixMilli(),
			UpdatedAt:    entity.UpdatedAt.UnixMilli(),
		},
	}, nil
}