package resourcerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/resourcerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAlbumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAlbumLogic {
	return &UpdateAlbumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新相册
func (l *UpdateAlbumLogic) UpdateAlbum(in *resourcerpc.UpdateAlbumReq) (*resourcerpc.UpdateAlbumResp, error) {
	entity := convertUpdateAlbumIn(in)
	if _, err := l.svcCtx.TAlbumModel.Save(l.ctx, entity); err != nil {
		return nil, err
	}

	record, err := l.svcCtx.TAlbumModel.FindById(l.ctx, entity.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}

	photoCount, err := l.svcCtx.TPhotoModel.FindCount(l.ctx, "album_id = ? and is_delete = ?", entity.Id, false)
	if err != nil {
		return nil, err
	}

	return &resourcerpc.UpdateAlbumResp{Album: convertAlbumOut(record, photoCount)}, nil
}
