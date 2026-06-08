// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package friend

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/socialrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取友链列表
func NewFindFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindFriendListLogic {
	return &FindFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindFriendListLogic) FindFriendList(req *types.QueryFriendReq) (resp *types.PageResp, err error) {
	in := &socialrpc.FindFriendListReq{
		Paginate: &socialrpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		LinkName: req.LinkName,
	}

	out, err := l.svcCtx.SocialRpc.FindFriendList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	list := make([]*types.FriendBackVO, 0, len(out.List))
	for _, item := range out.List {
		list = append(list, convertFriendTypes(item))
	}

	return &types.PageResp{
		Page:     out.Pagination.Page,
		PageSize: out.Pagination.PageSize,
		Total:    out.Pagination.Total,
		List:     list,
	}, nil
}
