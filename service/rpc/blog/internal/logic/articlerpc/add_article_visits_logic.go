package articlerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleVisitsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddArticleVisitsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleVisitsLogic {
	return &AddArticleVisitsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加文章访问量
func (l *AddArticleVisitsLogic) AddArticleVisits(in *articlerpc.AddArticleVisitsReq) (*articlerpc.AddArticleVisitsResp, error) {
	// todo: add your logic here and delete this line

	return &articlerpc.AddArticleVisitsResp{}, nil
}
