// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package visit_log

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesVisitLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除操作记录
func NewDeletesVisitLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesVisitLogLogic {
	return &DeletesVisitLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletesVisitLogLogic) DeletesVisitLog(req *types.IdsReq) (resp *types.BatchResp, err error) {
	// todo: add your logic here and delete this line

	return
}
