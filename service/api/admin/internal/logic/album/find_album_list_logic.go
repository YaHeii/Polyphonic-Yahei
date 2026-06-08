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

type FindAlbumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取相册列表
func NewFindAlbumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAlbumListLogic {
	return &FindAlbumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindAlbumListLogic) FindAlbumList(req *types.QueryAlbumReq) (resp *types.PageResp, err error) {
	in := &resourcerpc.FindAlbumListReq{
		Paginate: &resourcerpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		AlbumName: req.AlbumName,
		IsDelete:  &req.IsDelete,
	}

	out, err := l.svcCtx.ResourceRpc.FindAlbumList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.AlbumBackVO
	for _, v := range out.List {
		list = append(list, convertAlbumTypes(v))
	}

	resp = &types.PageResp{}
	resp.Page = out.Pagination.Page
	resp.PageSize = out.Pagination.PageSize
	resp.Total = out.Pagination.Total
	resp.List = list
	return resp, nil
}
