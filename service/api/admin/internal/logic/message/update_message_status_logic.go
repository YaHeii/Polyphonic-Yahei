// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package message

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMessageStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新留言状态
func NewUpdateMessageStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMessageStatusLogic {
	return &UpdateMessageStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMessageStatusLogic) UpdateMessageStatus(req *types.UpdateMessageStatusReq) (resp *types.BatchResp, err error) {
	// todo: add your logic here and delete this line

	return
}
