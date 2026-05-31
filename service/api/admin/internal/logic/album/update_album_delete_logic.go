// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package album

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAlbumDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新相册删除状态
func NewUpdateAlbumDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAlbumDeleteLogic {
	return &UpdateAlbumDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAlbumDeleteLogic) UpdateAlbumDelete(req *types.UpdateAlbumDeleteReq) (resp *types.BatchResp, err error) {
	// todo: add your logic here and delete this line

	return
}
