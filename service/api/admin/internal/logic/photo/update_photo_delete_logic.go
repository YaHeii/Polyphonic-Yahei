// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package photo

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePhotoDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新照片删除状态
func NewUpdatePhotoDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePhotoDeleteLogic {
	return &UpdatePhotoDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePhotoDeleteLogic) UpdatePhotoDelete(req *types.UpdatePhotoDeleteReq) (resp *types.BatchResp, err error) {
	// todo: add your logic here and delete this line

	return
}
