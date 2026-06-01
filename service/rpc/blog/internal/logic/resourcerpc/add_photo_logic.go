package resourcerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/resourcerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPhotoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPhotoLogic {
	return &AddPhotoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建照片
func (l *AddPhotoLogic) AddPhoto(in *resourcerpc.AddPhotoReq) (*resourcerpc.AddPhotoResp, error) {
	// todo: add your logic here and delete this line

	return &resourcerpc.AddPhotoResp{}, nil
}
