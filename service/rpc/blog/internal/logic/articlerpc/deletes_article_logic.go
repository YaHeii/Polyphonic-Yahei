package articlerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesArticleLogic {
	return &DeletesArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除文章
func (l *DeletesArticleLogic) DeletesArticle(in *articlerpc.DeletesArticleReq) (*articlerpc.DeletesArticleResp, error) {
	// todo: add your logic here and delete this line

	return &articlerpc.DeletesArticleResp{}, nil
}
