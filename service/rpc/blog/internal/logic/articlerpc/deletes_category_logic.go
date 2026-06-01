package articlerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesCategoryLogic {
	return &DeletesCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除文章分类
func (l *DeletesCategoryLogic) DeletesCategory(in *articlerpc.DeletesCategoryReq) (*articlerpc.DeletesCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &articlerpc.DeletesCategoryResp{}, nil
}
