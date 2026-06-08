// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package role

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/permissionrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取角色列表
func NewFindRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoleListLogic {
	return &FindRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindRoleListLogic) FindRoleList(req *types.QueryRoleReq) (resp *types.PageResp, err error) {
	in := &permissionrpc.FindRoleListReq{
		Paginate: &permissionrpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		RoleKey:   req.RoleKey,
		RoleLabel: req.RoleLabel,
		Status:    req.Status,
	}

	out, err := l.svcCtx.PermissionRpc.FindRoleList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	list := make([]*types.RoleBackVO, 0, len(out.List))
	for _, item := range out.List {
		list = append(list, convertRoleTypes(item))
	}

	return &types.PageResp{
		Page:     out.Pagination.Page,
		PageSize: out.Pagination.PageSize,
		Total:    out.Pagination.Total,
		List:     list,
	}, nil
}
