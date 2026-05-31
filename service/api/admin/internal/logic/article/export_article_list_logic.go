// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package article

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 导出文章列表
func NewExportArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportArticleListLogic {
	return &ExportArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportArticleListLogic) ExportArticleList(req *types.IdsReq) (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
