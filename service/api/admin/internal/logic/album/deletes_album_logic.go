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

type DeletesAlbumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除相册
func NewDeletesAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesAlbumLogic {
	return &DeletesAlbumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletesAlbumLogic) DeletesAlbum(req *types.IdsReq) (resp *types.BatchResp, err error) {
	in := &resourcerpc.DeletesAlbumReq{
		Ids: req.Ids,
	}

	out, err := l.svcCtx.ResourceRpc.DeletesAlbum(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}
	return resp, nil
}
