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

type FindTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取标签列表
func NewFindTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTagListLogic {
	return &FindTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindTagListLogic) FindTagList(req *types.QueryTagReq) (resp *types.PageResp, err error) {
	in := &articlerpc.FindTagListReq{
		Paginate: &articlerpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		TagName: req.TagName,
	}

	out, err := l.svcCtx.ArticleRpc.FindTagList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	list := make([]*types.TagBackVO, 0, len(out.List))
	for _, v := range out.List {
		list = append(list, convertTagDetailsTypes(v))
	}

	resp = &types.PageResp{
		List: list,
	}
	if out.Pagination != nil {
		resp.Page = out.Pagination.Page
		resp.PageSize = out.Pagination.PageSize
		resp.Total = out.Pagination.Total
	}
	return resp, nil
}
