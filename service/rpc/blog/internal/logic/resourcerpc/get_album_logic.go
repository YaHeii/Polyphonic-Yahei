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
	record, err := l.svcCtx.TAlbumModel.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	photoCount, err := l.svcCtx.TPhotoModel.FindCount(l.ctx, "album_id = ? and is_delete = ?", in.Id, false)
	if err != nil {
		return nil, err
	}

	return &resourcerpc.GetAlbumResp{Album: convertAlbumOut(record, photoCount)}, nil
}
