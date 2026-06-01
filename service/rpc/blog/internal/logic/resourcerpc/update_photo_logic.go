package resourcerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/resourcerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePhotoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePhotoLogic {
	return &UpdatePhotoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新照片
func (l *UpdatePhotoLogic) UpdatePhoto(in *resourcerpc.UpdatePhotoReq) (*resourcerpc.UpdatePhotoResp, error) {
	// todo: add your logic here and delete this line

	return &resourcerpc.UpdatePhotoResp{}, nil
}
