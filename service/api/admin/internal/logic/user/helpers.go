package user

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/authctx"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/jsonconv"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/permissionrpc"
)

func currentUserID(ctx context.Context) string {
	return authctx.CurrentUserID(ctx)
}

func convertUserApi(in *permissionrpc.Api) *types.UserApi {
	children := make([]*types.UserApi, 0, len(in.Children))
	for _, child := range in.Children {
		children = append(children, convertUserApi(child))
	}

	return &types.UserApi{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Name:      in.Name,
		Path:      in.Path,
		Method:    in.Method,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
		Children:  children,
	}
}

func convertUserMenu(in *permissionrpc.Menu) *types.UserMenu {
	children := make([]*types.UserMenu, 0, len(in.Children))
	for _, child := range in.Children {
		children = append(children, convertUserMenu(child))
	}

	var meta types.UserMenuMeta
	if in.Meta != nil {
		meta = types.UserMenuMeta{
			Title:      in.Meta.Title,
			Icon:       in.Meta.Icon,
			Hidden:     in.Meta.Visible,
			AlwaysShow: in.Meta.AlwaysShow,
			Affix:      false,
			KeepAlive:  in.Meta.KeepAlive,
			Breadcrumb: false,
		}
	}

	return &types.UserMenu{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Path:      in.Path,
		Name:      in.Name,
		Component: in.Component,
		Redirect:  in.Redirect,
		Meta:      meta,
		Children:  children,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}

func extractMenuPerms(menus []*permissionrpc.Menu) []string {
	perms := make([]string, 0)
	for _, menu := range menus {
		if menu.Meta != nil && menu.Meta.Perm != "" {
			perms = append(perms, menu.Meta.Perm)
		}
		perms = append(perms, extractMenuPerms(menu.Children)...)
	}
	return perms
}

func convertUserInfo(user *accountrpc.UserInfo, oauthResp *accountrpc.GetUserOauthInfoResp, menuResp *permissionrpc.FindUserMenusResp) *types.UserInfoResp {
	var ext types.UserInfoExt
	if user.Info != "" {
		_ = jsonconv.JsonToAny(user.Info, &ext)
	}

	thirdParty := make([]*types.UserThirdPartyInfo, 0, len(oauthResp.List))
	for _, item := range oauthResp.List {
		thirdParty = append(thirdParty, &types.UserThirdPartyInfo{
			Platform:  item.Platform,
			OpenId:    item.OpenId,
			Nickname:  item.Nickname,
			Avatar:    item.Avatar,
			CreatedAt: item.CreatedAt,
		})
	}

	roles := make([]string, 0, len(user.Roles))
	for _, role := range user.Roles {
		roles = append(roles, role.RoleKey)
	}

	return &types.UserInfoResp{
		UserId:       user.UserId,
		Username:     user.Username,
		Nickname:     user.Nickname,
		Avatar:       user.Avatar,
		Email:        user.Email,
		Phone:        user.Phone,
		CreatedAt:    user.CreatedAt,
		RegisterType: user.RegisterType,
		UserInfoExt:  ext,
		ThirdParty:   thirdParty,
		Roles:        roles,
		Perms:        extractMenuPerms(menuResp.List),
	}
}
