// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package page

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/resourcerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建页面
func NewAddPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPageLogic {
	return &AddPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPageLogic) AddPage(req *types.NewPageReq) (resp *types.PageBackVO, err error) {
	in := &resourcerpc.AddPageReq{
		Id:             req.Id,
		PageName:       req.PageName,
		PageLabel:      req.PageLabel,
		PageCover:      req.PageCover,
		IsCarousel:     req.IsCarousel != 0,
		CarouselCovers: req.CarouselCovers,
	}

	out, err := l.svcCtx.ResourceRpc.AddPage(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convertPageTypes(out.Page), nil
}
