package articlerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleListLogic {
	return &FindArticleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询文章列表
func (l *FindArticleListLogic) FindArticleList(in *articlerpc.FindArticleListReq) (*articlerpc.FindArticleListResp, error) {
	// todo: add your logic here and delete this line

	return &articlerpc.FindArticleListResp{}, nil
}
