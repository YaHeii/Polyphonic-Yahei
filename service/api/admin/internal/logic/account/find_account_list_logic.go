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

type FindAccountListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询用户列表
func NewFindAccountListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAccountListLogic {
	return &FindAccountListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindAccountListLogic) FindAccountList(req *types.QueryAccountReq) (resp *types.PageResp, err error) {
	in := &accountrpc.FindUserListReq{
		Paginate: &accountrpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		Username: req.Username,
		Nickname: req.Nickname,
		Email:    req.Email,
		Phone:    req.Phone,
		Status:   req.Status,
		UserIds:  req.UserIds,
	}

	out, err := l.svcCtx.AccountRpc.FindUserInfoList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.UserInfoDetail
	for _, v := range out.List {
		m := convertUserInfoTypes(v)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = out.Pagination.Page
	resp.PageSize = out.Pagination.PageSize
	resp.Total = out.Pagination.Total
	resp.List = list
	return resp, nil
}
