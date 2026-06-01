package articlerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleRelationLogic {
	return &GetArticleRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询关联文章
func (l *GetArticleRelationLogic) GetArticleRelation(in *articlerpc.GetArticleRelationReq) (*articlerpc.GetArticleRelationResp, error) {
	// todo: add your logic here and delete this line

	return &articlerpc.GetArticleRelationResp{}, nil
}
