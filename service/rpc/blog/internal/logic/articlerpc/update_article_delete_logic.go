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
	record, err := l.svcCtx.TArticleModel.FindById(l.ctx, in.ArticleId)
	if err != nil {
		return nil, err
	}

	record.IsDelete = in.IsDelete
	err = l.svcCtx.TArticleModel.Update(l.ctx, record)
	if err != nil {
		return nil, err
	}

	helper := NewArticleHelperLogic(l.ctx, l.svcCtx)
	return &articlerpc.UpdateArticleDeleteResp{
		Article: helper.convertArticlePreviewOut(record),
	}, nil
}
