package resourcerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/resourcerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePhotoDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePhotoDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePhotoDeleteLogic {
	return &UpdatePhotoDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新照片删除状态
func (l *UpdatePhotoDeleteLogic) UpdatePhotoDelete(in *resourcerpc.UpdatePhotoDeleteReq) (*resourcerpc.UpdatePhotoDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &resourcerpc.UpdatePhotoDeleteResp{}, nil
}
