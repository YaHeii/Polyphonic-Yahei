package articlerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleTopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateArticleTopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleTopLogic {
	return &UpdateArticleTopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新文章置顶
func (l *UpdateArticleTopLogic) UpdateArticleTop(in *articlerpc.UpdateArticleTopReq) (*articlerpc.UpdateArticleTopResp, error) {
	// todo: add your logic here and delete this line

	return &articlerpc.UpdateArticleTopResp{}, nil
}
