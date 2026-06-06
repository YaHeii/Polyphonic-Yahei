package resourcerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/resourcerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAlbumListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAlbumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAlbumListLogic {
	return &FindAlbumListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询相册列表
func (l *FindAlbumListLogic) FindAlbumList(in *resourcerpc.FindAlbumListReq) (*resourcerpc.FindAlbumListResp, error) {
	page, size, sorts, conditions, params := buildAlbumQuery(in)
	records, total, err := l.svcCtx.TAlbumModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	photoCounts, err := l.svcCtx.TPhotoModel.CountByAlbumIDs(l.ctx, collectAlbumIDs(records))
	if err != nil {
		return nil, err
	}

	return &resourcerpc.FindAlbumListResp{
		Pagination: buildPageResp(page, size, total),
		List:       convertAlbumListOut(records, photoCounts),
	}, nil
}
