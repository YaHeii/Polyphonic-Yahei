package resourcerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/resourcerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesAlbumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesAlbumLogic {
	return &DeletesAlbumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除相册
func (l *DeletesAlbumLogic) DeletesAlbum(in *resourcerpc.DeletesAlbumReq) (*resourcerpc.DeletesAlbumResp, error) {
	// todo: add your logic here and delete this line

	return &resourcerpc.DeletesAlbumResp{}, nil
}
