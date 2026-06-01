package articlerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLikeArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLikeArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLikeArticleLogic {
	return &FindUserLikeArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户点赞的文章
func (l *FindUserLikeArticleLogic) FindUserLikeArticle(in *articlerpc.FindUserLikeArticleReq) (*articlerpc.FindLikeArticleResp, error) {
	// todo: add your logic here and delete this line

	return &articlerpc.FindLikeArticleResp{}, nil
}
