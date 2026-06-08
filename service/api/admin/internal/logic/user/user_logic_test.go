package user

import (
	"context"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizheader"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/permissionrpc"
	"google.golang.org/grpc"
)

type stubUserAccountRPC struct {
	accountrpc.AccountRpc
	getUserInfoReq  *accountrpc.GetUserInfoReq
	getUserInfoResp *accountrpc.GetUserInfoResp
	getOauthReq     *accountrpc.GetUserOauthInfoReq
	getOauthResp    *accountrpc.GetUserOauthInfoResp
}

func (s *stubUserAccountRPC) GetUserInfo(_ context.Context, in *accountrpc.GetUserInfoReq, _ ...grpc.CallOption) (*accountrpc.GetUserInfoResp, error) {
	s.getUserInfoReq = in
	return s.getUserInfoResp, nil
}

func (s *stubUserAccountRPC) GetUserOauthInfo(_ context.Context, in *accountrpc.GetUserOauthInfoReq, _ ...grpc.CallOption) (*accountrpc.GetUserOauthInfoResp, error) {
	s.getOauthReq = in
	return s.getOauthResp, nil
}

type stubUserPermissionRPC struct {
	permissionrpc.PermissionRpc
	apisReq   *permissionrpc.FindUserApisReq
	apisResp  *permissionrpc.FindUserApisResp
	menusReq  *permissionrpc.FindUserMenusReq
	menusResp *permissionrpc.FindUserMenusResp
	rolesReq  *permissionrpc.FindUserRolesReq
	rolesResp *permissionrpc.FindUserRolesResp
}

func (s *stubUserPermissionRPC) FindUserApis(_ context.Context, in *permissionrpc.FindUserApisReq, _ ...grpc.CallOption) (*permissionrpc.FindUserApisResp, error) {
	s.apisReq = in
	return s.apisResp, nil
}

func (s *stubUserPermissionRPC) FindUserMenus(_ context.Context, in *permissionrpc.FindUserMenusReq, _ ...grpc.CallOption) (*permissionrpc.FindUserMenusResp, error) {
	s.menusReq = in
	return s.menusResp, nil
}

func (s *stubUserPermissionRPC) FindUserRoles(_ context.Context, in *permissionrpc.FindUserRolesReq, _ ...grpc.CallOption) (*permissionrpc.FindUserRolesResp, error) {
	s.rolesReq = in
	return s.rolesResp, nil
}

func TestGetUserApisBuildsRequestAndMapsTree(t *testing.T) {
	ctx := context.WithValue(context.Background(), bizheader.HeaderUid, "u-1")
	permissionRPC := &stubUserPermissionRPC{
		apisResp: &permissionrpc.FindUserApisResp{
			List: []*permissionrpc.Api{
				{
					Id:       1,
					Name:     "root",
					Path:     "/root",
					Method:   "GET",
					Children: []*permissionrpc.Api{{Id: 2, ParentId: 1, Name: "child"}},
				},
			},
		},
	}
	logic := NewGetUserApisLogic(ctx, &svc.ServiceContext{PermissionRpc: permissionRPC})

	resp, err := logic.GetUserApis(&types.EmptyReq{})
	if err != nil {
		t.Fatalf("GetUserApis returned error: %v", err)
	}
	if permissionRPC.apisReq == nil || permissionRPC.apisReq.UserId != "u-1" {
		t.Fatalf("unexpected apis request: %#v", permissionRPC.apisReq)
	}
	if len(resp.List) != 1 || len(resp.List[0].Children) != 1 || resp.List[0].Children[0].Id != 2 {
		t.Fatalf("unexpected apis response: %#v", resp)
	}
}

func TestGetUserMenusBuildsRequestAndMapsTree(t *testing.T) {
	ctx := context.WithValue(context.Background(), bizheader.HeaderUid, "u-1")
	permissionRPC := &stubUserPermissionRPC{
		menusResp: &permissionrpc.FindUserMenusResp{
			List: []*permissionrpc.Menu{
				{
					Id:        1,
					Path:      "/dashboard",
					Name:      "Dashboard",
					Component: "Layout",
					Meta: &permissionrpc.MenuMeta{
						Title:      "Dashboard",
						Icon:       "home",
						Visible:    true,
						AlwaysShow: true,
					},
					Children: []*permissionrpc.Menu{
						{
							Id:   2,
							Name: "Analysis",
							Meta: &permissionrpc.MenuMeta{Title: "Analysis"},
						},
					},
				},
			},
		},
	}
	logic := NewGetUserMenusLogic(ctx, &svc.ServiceContext{PermissionRpc: permissionRPC})

	resp, err := logic.GetUserMenus(&types.EmptyReq{})
	if err != nil {
		t.Fatalf("GetUserMenus returned error: %v", err)
	}
	if permissionRPC.menusReq == nil || permissionRPC.menusReq.UserId != "u-1" {
		t.Fatalf("unexpected menus request: %#v", permissionRPC.menusReq)
	}
	if len(resp.List) != 1 || !resp.List[0].Meta.Hidden || !resp.List[0].Meta.AlwaysShow || len(resp.List[0].Children) != 1 {
		t.Fatalf("unexpected menus response: %#v", resp)
	}
}

func TestGetUserRolesBuildsRequestAndMapsRoles(t *testing.T) {
	ctx := context.WithValue(context.Background(), bizheader.HeaderUid, "u-1")
	permissionRPC := &stubUserPermissionRPC{
		rolesResp: &permissionrpc.FindUserRolesResp{
			List: []*permissionrpc.Role{
				{Id: 1, ParentId: 2, RoleKey: "admin", RoleLabel: "Admin", RoleComment: "root"},
			},
		},
	}
	logic := NewGetUserRolesLogic(ctx, &svc.ServiceContext{PermissionRpc: permissionRPC})

	resp, err := logic.GetUserRoles(&types.EmptyReq{})
	if err != nil {
		t.Fatalf("GetUserRoles returned error: %v", err)
	}
	if permissionRPC.rolesReq == nil || permissionRPC.rolesReq.UserId != "u-1" {
		t.Fatalf("unexpected roles request: %#v", permissionRPC.rolesReq)
	}
	if len(resp.List) != 1 || resp.List[0].RoleKey != "admin" {
		t.Fatalf("unexpected roles response: %#v", resp)
	}
}

func TestGetUserInfoBuildsRequestsAndAggregatesResponse(t *testing.T) {
	ctx := context.WithValue(context.Background(), bizheader.HeaderUid, "u-1")
	accountRPC := &stubUserAccountRPC{
		getUserInfoResp: &accountrpc.GetUserInfoResp{
			User: &accountrpc.UserInfo{
				UserId:       "u-1",
				Username:     "demo",
				Nickname:     "Demo",
				Avatar:       "avatar",
				Email:        "demo@example.com",
				Phone:        "188",
				RegisterType: "email",
				CreatedAt:    100,
				Info:         `{"intro":"hello","website":"https://example.com"}`,
				Roles: []*accountrpc.UserRoleLabel{
					{RoleKey: "admin", RoleLabel: "Admin"},
				},
			},
		},
		getOauthResp: &accountrpc.GetUserOauthInfoResp{
			List: []*accountrpc.UserOauthInfo{
				{Platform: "github", OpenId: "oid-1", Nickname: "gh", Avatar: "a", CreatedAt: 10},
			},
		},
	}
	permissionRPC := &stubUserPermissionRPC{
		menusResp: &permissionrpc.FindUserMenusResp{
			List: []*permissionrpc.Menu{
				{
					Id:   1,
					Name: "Root",
					Meta: &permissionrpc.MenuMeta{Perm: "system:user:list"},
					Children: []*permissionrpc.Menu{
						{Id: 2, Name: "Child", Meta: &permissionrpc.MenuMeta{Perm: "system:user:edit"}},
					},
				},
			},
		},
	}
	logic := NewGetUserInfoLogic(ctx, &svc.ServiceContext{
		AccountRpc:    accountRPC,
		PermissionRpc: permissionRPC,
	})

	resp, err := logic.GetUserInfo(&types.EmptyReq{})
	if err != nil {
		t.Fatalf("GetUserInfo returned error: %v", err)
	}
	if accountRPC.getUserInfoReq == nil || accountRPC.getUserInfoReq.UserId != "u-1" {
		t.Fatalf("unexpected get user info request: %#v", accountRPC.getUserInfoReq)
	}
	if accountRPC.getOauthReq == nil || accountRPC.getOauthReq.UserId != "u-1" {
		t.Fatalf("unexpected get oauth request: %#v", accountRPC.getOauthReq)
	}
	if permissionRPC.menusReq == nil || permissionRPC.menusReq.UserId != "u-1" {
		t.Fatalf("unexpected get menus request: %#v", permissionRPC.menusReq)
	}
	if resp.UserId != "u-1" || resp.UserInfoExt.Intro != "hello" || len(resp.ThirdParty) != 1 || len(resp.Roles) != 1 || len(resp.Perms) != 2 {
		t.Fatalf("unexpected user info response: %#v", resp)
	}
}
