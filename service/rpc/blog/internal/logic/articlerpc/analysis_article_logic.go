package articlerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnalysisArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnalysisArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnalysisArticleLogic {
	return &AnalysisArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分析文章数量
func (l *AnalysisArticleLogic) AnalysisArticle(in *articlerpc.AnalysisArticleReq) (*articlerpc.AnalysisArticleResp, error) {
	// todo: add your logic here and delete this line

	return &articlerpc.AnalysisArticleResp{}, nil
}
