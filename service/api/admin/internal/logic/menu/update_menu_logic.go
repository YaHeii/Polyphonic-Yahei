// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package menu

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/permissionrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新菜单
func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMenuLogic) UpdateMenu(req *types.NewMenuReq) (resp *types.MenuBackVO, err error) {
	out, err := l.svcCtx.PermissionRpc.UpdateMenu(l.ctx, &permissionrpc.UpdateMenuReq{
		Id:        req.Id,
		ParentId:  req.ParentId,
		Path:      req.Path,
		Name:      req.Name,
		Component: req.Component,
		Redirect:  req.Redirect,
		Meta: &permissionrpc.MenuMeta{
			Type:       req.Type,
			Title:      req.Title,
			Icon:       req.Icon,
			Rank:       req.Rank,
			Perm:       req.Perm,
			Params:     convertMenuPb(req).Meta.Params,
			KeepAlive:  req.KeepAlive == 1,
			AlwaysShow: req.AlwaysShow == 1,
			Visible:    req.Visible == 1,
			Status:     req.Status == 1,
		},
	})
	if err != nil {
		return nil, err
	}

	return convertMenuTypes(out.Menu), nil
}
