// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package tag

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/articlerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新标签
func NewUpdateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTagLogic {
	return &UpdateTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTagLogic) UpdateTag(req *types.NewTagReq) (resp *types.TagBackVO, err error) {
	in := &articlerpc.UpdateTagReq{
		Id:      req.Id,
		TagName: req.TagName,
	}

	out, err := l.svcCtx.ArticleRpc.UpdateTag(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convertTagTypes(out.Tag), nil
}
