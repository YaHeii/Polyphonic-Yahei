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

type FindMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取菜单列表
func NewFindMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMenuListLogic {
	return &FindMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindMenuListLogic) FindMenuList(req *types.QueryMenuReq) (resp *types.MenuPageResp, err error) {
	out, err := l.svcCtx.PermissionRpc.FindMenuList(l.ctx, &permissionrpc.FindMenuListReq{
		Name:  req.Name,
		Title: req.Title,
	})
	if err != nil {
		return nil, err
	}

	list := make([]*types.MenuBackVO, 0, len(out.List))
	for _, item := range out.List {
		list = append(list, convertMenuTypes(item))
	}

	return &types.MenuPageResp{
		PageMeta: types.PageMeta{
			Page:     0,
			PageSize: int64(len(list)),
			Total:    int64(len(list)),
		},
		List: list,
	}, nil
}
