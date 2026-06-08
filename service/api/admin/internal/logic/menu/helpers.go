package menu

import (
	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/jsonconv"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/permissionrpc"
)

func convertMenuPb(in *types.NewMenuReq) *permissionrpc.AddMenuReq {
	children := make([]*permissionrpc.AddMenuReq, 0, len(in.Children))
	for _, child := range in.Children {
		children = append(children, convertMenuPb(child))
	}

	return &permissionrpc.AddMenuReq{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Path:      in.Path,
		Name:      in.Name,
		Component: in.Component,
		Redirect:  in.Redirect,
		Children:  children,
		Meta: &permissionrpc.MenuMeta{
			Type:       in.Type,
			Title:      in.Title,
			Icon:       in.Icon,
			Rank:       in.Rank,
			Perm:       in.Perm,
			Params:     jsonconv.AnyToJsonNE(in.Params),
			KeepAlive:  in.KeepAlive == 1,
			AlwaysShow: in.AlwaysShow == 1,
			Visible:    in.Visible == 1,
			Status:     in.Status == 1,
		},
	}
}

func convertMenuTypes(in *permissionrpc.Menu) *types.MenuBackVO {
	children := make([]*types.MenuBackVO, 0, len(in.Children))
	for _, child := range in.Children {
		children = append(children, convertMenuTypes(child))
	}

	var params []*types.MenuMetaParams
	if in.Meta != nil && in.Meta.Params != "" {
		_ = jsonconv.JsonToAny(in.Meta.Params, &params)
	}

	meta := types.MenuMeta{}
	if in.Meta != nil {
		meta = types.MenuMeta{
			Type:       in.Meta.Type,
			Title:      in.Meta.Title,
			Icon:       in.Meta.Icon,
			Rank:       in.Meta.Rank,
			Perm:       in.Meta.Perm,
			Params:     params,
			KeepAlive:  boolToInt64(in.Meta.KeepAlive),
			AlwaysShow: boolToInt64(in.Meta.AlwaysShow),
			Visible:    boolToInt64(in.Meta.Visible),
			Status:     boolToInt64(in.Meta.Status),
		}
	}

	return &types.MenuBackVO{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Path:      in.Path,
		Name:      in.Name,
		Component: in.Component,
		Redirect:  in.Redirect,
		MenuMeta:  meta,
		Children:  children,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}

func boolToInt64(value bool) int64 {
	if value {
		return 1
	}
	return 0
}
