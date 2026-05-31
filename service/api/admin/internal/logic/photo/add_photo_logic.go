// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package photo

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPhotoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建照片
func NewAddPhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPhotoLogic {
	return &AddPhotoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPhotoLogic) AddPhoto(req *types.NewPhotoReq) (resp *types.PhotoBackVO, err error) {
	// todo: add your logic here and delete this line

	return
}
