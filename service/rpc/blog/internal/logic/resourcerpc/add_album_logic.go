package resourcerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/resourcerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAlbumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAlbumLogic {
	return &AddAlbumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建相册
func (l *AddAlbumLogic) AddAlbum(in *resourcerpc.AddAlbumReq) (*resourcerpc.AddAlbumResp, error) {
	// todo: add your logic here and delete this line

	return &resourcerpc.AddAlbumResp{}, nil
}
