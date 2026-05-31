// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package visitor

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindVisitorListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取游客列表
func NewFindVisitorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindVisitorListLogic {
	return &FindVisitorListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindVisitorListLogic) FindVisitorList(req *types.QueryVisitorReq) (resp *types.PageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
