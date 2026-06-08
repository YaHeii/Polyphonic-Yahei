// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package album

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/resourcerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAlbumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建相册
func NewAddAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAlbumLogic {
	return &AddAlbumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAlbumLogic) AddAlbum(req *types.NewAlbumReq) (resp *types.AlbumBackVO, err error) {
	in := &resourcerpc.AddAlbumReq{
		Id:         req.Id,
		AlbumName:  req.AlbumName,
		AlbumDesc:  req.AlbumDesc,
		AlbumCover: req.AlbumCover,
		IsDelete:   req.IsDelete,
		Status:     req.Status,
	}

	out, err := l.svcCtx.ResourceRpc.AddAlbum(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convertAlbumTypes(out.Album), nil
}
