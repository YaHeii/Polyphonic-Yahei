package resourcerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/resourcerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesPhotoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesPhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesPhotoLogic {
	return &DeletesPhotoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除照片
func (l *DeletesPhotoLogic) DeletesPhoto(in *resourcerpc.DeletesPhotoReq) (*resourcerpc.DeletesPhotoResp, error) {
	rows, err := l.svcCtx.TPhotoModel.Deletes(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &resourcerpc.DeletesPhotoResp{SuccessCount: rows}, nil
}
