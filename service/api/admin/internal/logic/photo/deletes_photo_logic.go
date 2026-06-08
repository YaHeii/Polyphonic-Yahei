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

type DeletesPhotoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除照片
func NewDeletesPhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesPhotoLogic {
	return &DeletesPhotoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletesPhotoLogic) DeletesPhoto(req *types.IdsReq) (resp *types.BatchResp, err error) {
	in := &resourcerpc.DeletesPhotoReq{
		Ids: req.Ids,
	}

	out, err := l.svcCtx.ResourceRpc.DeletesPhoto(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{SuccessCount: out.SuccessCount}, nil
}
