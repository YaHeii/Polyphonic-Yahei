package articlerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateArticleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleDeleteLogic {
	return &UpdateArticleDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新文章删除
func (l *UpdateArticleDeleteLogic) UpdateArticleDelete(in *articlerpc.UpdateArticleDeleteReq) (*articlerpc.UpdateArticleDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &articlerpc.UpdateArticleDeleteResp{}, nil
}
