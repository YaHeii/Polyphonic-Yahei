// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package article

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新文章删除状态
func NewUpdateArticleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleDeleteLogic {
	return &UpdateArticleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateArticleDeleteLogic) UpdateArticleDelete(req *types.UpdateArticleDeleteReq) (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
