// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package account

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAccountOnlineListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询在线用户列表
func NewFindAccountOnlineListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAccountOnlineListLogic {
	return &FindAccountOnlineListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindAccountOnlineListLogic) FindAccountOnlineList(req *types.QueryAccountReq) (resp *types.AccountPageResp, err error) {
	in := &accountrpc.FindUserListReq{
		Paginate: &accountrpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		Nickname: req.Nickname,
	}

	out, err := l.svcCtx.AccountRpc.FindUserOnlineList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.UserInfoDetail
	for _, v := range out.List {
		m := convertUserInfoTypes(v)
		list = append(list, m)
	}

	resp = &types.AccountPageResp{}
	resp.Page = out.Pagination.Page
	resp.PageSize = out.Pagination.PageSize
	resp.Total = out.Pagination.Total
	resp.List = list
	return resp, nil
}
