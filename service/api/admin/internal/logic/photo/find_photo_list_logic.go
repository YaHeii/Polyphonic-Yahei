// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package photo

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/resourcerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPhotoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取照片列表
func NewFindPhotoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPhotoListLogic {
	return &FindPhotoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPhotoListLogic) FindPhotoList(req *types.QueryPhotoReq) (resp *types.PhotoPageResp, err error) {
	in := &resourcerpc.FindPhotoListReq{
		Paginate: &resourcerpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		AlbumId:  req.AlbumId,
		IsDelete: &req.IsDelete,
	}

	out, err := l.svcCtx.ResourceRpc.FindPhotoList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	list := make([]*types.PhotoBackVO, 0, len(out.List))
	for _, item := range out.List {
		list = append(list, convertPhotoTypes(item))
	}

	return &types.PhotoPageResp{
		PageMeta: types.PageMeta{
			Page:     out.Pagination.Page,
			PageSize: out.Pagination.PageSize,
			Total:    out.Pagination.Total,
		},
		List: list,
	}, nil
}
