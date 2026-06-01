package resourcerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/resourcerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAlbumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAlbumLogic {
	return &GetAlbumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取相册
func (l *GetAlbumLogic) GetAlbum(in *resourcerpc.GetAlbumReq) (*resourcerpc.GetAlbumResp, error) {
	// todo: add your logic here and delete this line

	return &resourcerpc.GetAlbumResp{}, nil
}
