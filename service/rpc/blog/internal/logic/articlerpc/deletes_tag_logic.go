package articlerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesTagLogic {
	return &DeletesTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除标签
func (l *DeletesTagLogic) DeletesTag(in *articlerpc.DeletesTagReq) (*articlerpc.DeletesTagResp, error) {
	// todo: add your logic here and delete this line

	return &articlerpc.DeletesTagResp{}, nil
}
