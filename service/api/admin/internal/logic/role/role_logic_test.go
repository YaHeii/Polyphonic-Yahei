package role

import (
	"context"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/permissionx"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/permissionrpc"
	"google.golang.org/grpc"
)

type stubRolePermissionRPC struct {
	addReq         *permissionrpc.AddRoleReq
	addResp        *permissionrpc.AddRoleResp
	updateReq      *permissionrpc.UpdateRoleReq
	updateResp     *permissionrpc.UpdateRoleResp
	deleteReq      *permissionrpc.DeletesRoleReq
	deleteResp     *permissionrpc.DeletesRoleResp
	findReq        *permissionrpc.FindRoleListReq
	findResp       *permissionrpc.FindRoleListResp
	resReq         *permissionrpc.FindRoleResourcesReq
	resResp        *permissionrpc.FindRoleResourcesResp
	updateApisReq  *permissionrpc.UpdateRoleApisReq
	updateMenusReq *permissionrpc.UpdateRoleMenusReq
}

func (s *stubRolePermissionRPC) AddApi(context.Context, *permissionrpc.AddApiReq, ...grpc.CallOption) (*permissionrpc.AddApiResp, error) {
	panic("unexpected call")
}
func (s *stubRolePermissionRPC) UpdateApi(context.Context, *permissionrpc.UpdateApiReq, ...grpc.CallOption) (*permissionrpc.UpdateApiResp, error) {
	panic("unexpected call")
}
func (s *stubRolePermissionRPC) DeletesApi(context.Context, *permissionrpc.DeletesApiReq, ...grpc.CallOption) (*permissionrpc.DeletesApiResp, error) {
	panic("unexpected call")
}
func (s *stubRolePermissionRPC) FindApiList(context.Context, *permissionrpc.FindApiListReq, ...grpc.CallOption) (*permissionrpc.FindApiListResp, error) {
	panic("unexpected call")
}
func (s *stubRolePermissionRPC) SyncApiList(context.Context, *permissionrpc.SyncApiListReq, ...grpc.CallOption) (*permissionrpc.SyncApiListResp, error) {
	panic("unexpected call")
}
func (s *stubRolePermissionRPC) CleanApiList(context.Context, *permissionrpc.CleanApiListReq, ...grpc.CallOption) (*permissionrpc.CleanApiListResp, error) {
	panic("unexpected call")
}
func (s *stubRolePermissionRPC) FindAllApi(context.Context, *permissionrpc.FindAllApiReq, ...grpc.CallOption) (*permissionrpc.FindAllApiResp, error) {
	panic("unexpected call")
}
func (s *stubRolePermissionRPC) AddMenu(context.Context, *permissionrpc.AddMenuReq, ...grpc.CallOption) (*permissionrpc.AddMenuResp, error) {
	panic("unexpected call")
}
func (s *stubRolePermissionRPC) UpdateMenu(context.Context, *permissionrpc.UpdateMenuReq, ...grpc.CallOption) (*permissionrpc.UpdateMenuResp, error) {
	panic("unexpected call")
}
func (s *stubRolePermissionRPC) DeletesMenu(context.Context, *permissionrpc.DeletesMenuReq, ...grpc.CallOption) (*permissionrpc.DeletesMenuResp, error) {
	panic("unexpected call")
}
func (s *stubRolePermissionRPC) FindMenuList(context.Context, *permissionrpc.FindMenuListReq, ...grpc.CallOption) (*permissionrpc.FindMenuListResp, error) {
	panic("unexpected call")
}
func (s *stubRolePermissionRPC) SyncMenuList(context.Context, *permissionrpc.SyncMenuListReq, ...grpc.CallOption) (*permissionrpc.SyncMenuListResp, error) {
	panic("unexpected call")
}
func (s *stubRolePermissionRPC) CleanMenuList(context.Context, *permissionrpc.CleanMenuListReq, ...grpc.CallOption) (*permissionrpc.CleanMenuListResp, error) {
	panic("unexpected call")
}
func (s *stubRolePermissionRPC) FindAllMenu(context.Context, *permissionrpc.FindAllMenuReq, ...grpc.CallOption) (*permissionrpc.FindAllMenuResp, error) {
	panic("unexpected call")
}
func (s *stubRolePermissionRPC) AddRole(_ context.Context, in *permissionrpc.AddRoleReq, _ ...grpc.CallOption) (*permissionrpc.AddRoleResp, error) {
	s.addReq = in
	return s.addResp, nil
}
func (s *stubRolePermissionRPC) UpdateRole(_ context.Context, in *permissionrpc.UpdateRoleReq, _ ...grpc.CallOption) (*permissionrpc.UpdateRoleResp, error) {
	s.updateReq = in
	return s.updateResp, nil
}
func (s *stubRolePermissionRPC) DeletesRole(_ context.Context, in *permissionrpc.DeletesRoleReq, _ ...grpc.CallOption) (*permissionrpc.DeletesRoleResp, error) {
	s.deleteReq = in
	return s.deleteResp, nil
}
func (s *stubRolePermissionRPC) FindRoleList(_ context.Context, in *permissionrpc.FindRoleListReq, _ ...grpc.CallOption) (*permissionrpc.FindRoleListResp, error) {
	s.findReq = in
	return s.findResp, nil
}
func (s *stubRolePermissionRPC) FindAllRole(context.Context, *permissionrpc.FindAllRoleReq, ...grpc.CallOption) (*permissionrpc.FindAllRoleResp, error) {
	panic("unexpected call")
}
func (s *stubRolePermissionRPC) UpdateRoleMenus(_ context.Context, in *permissionrpc.UpdateRoleMenusReq, _ ...grpc.CallOption) (*permissionrpc.UpdateRoleMenusResp, error) {
	s.updateMenusReq = in
	return &permissionrpc.UpdateRoleMenusResp{}, nil
}
func (s *stubRolePermissionRPC) UpdateRoleApis(_ context.Context, in *permissionrpc.UpdateRoleApisReq, _ ...grpc.CallOption) (*permissionrpc.UpdateRoleApisResp, error) {
	s.updateApisReq = in
	return &permissionrpc.UpdateRoleApisResp{}, nil
}
func (s *stubRolePermissionRPC) FindRoleResources(_ context.Context, in *permissionrpc.FindRoleResourcesReq, _ ...grpc.CallOption) (*permissionrpc.FindRoleResourcesResp, error) {
	s.resReq = in
	return s.resResp, nil
}
func (s *stubRolePermissionRPC) UpdateUserRole(context.Context, *permissionrpc.UpdateUserRoleReq, ...grpc.CallOption) (*permissionrpc.UpdateUserRoleResp, error) {
	panic("unexpected call")
}
func (s *stubRolePermissionRPC) FindUserApis(context.Context, *permissionrpc.FindUserApisReq, ...grpc.CallOption) (*permissionrpc.FindUserApisResp, error) {
	panic("unexpected call")
}
func (s *stubRolePermissionRPC) FindUserMenus(context.Context, *permissionrpc.FindUserMenusReq, ...grpc.CallOption) (*permissionrpc.FindUserMenusResp, error) {
	panic("unexpected call")
}
func (s *stubRolePermissionRPC) FindUserRoles(context.Context, *permissionrpc.FindUserRolesReq, ...grpc.CallOption) (*permissionrpc.FindUserRolesResp, error) {
	panic("unexpected call")
}

type stubPermissionHolder struct {
	reloadCount int
}

func (s *stubPermissionHolder) LoadPolicy() error   { return nil }
func (s *stubPermissionHolder) ReloadPolicy() error { s.reloadCount++; return nil }
func (s *stubPermissionHolder) Enforce(string, string, string) error {
	return nil
}

var _ permissionx.PermissionHolder = (*stubPermissionHolder)(nil)

func TestAddAndUpdateRoleBuildRequests(t *testing.T) {
	rpc := &stubRolePermissionRPC{
		addResp:    &permissionrpc.AddRoleResp{Role: &permissionrpc.Role{Id: 1, RoleKey: "admin"}},
		updateResp: &permissionrpc.UpdateRoleResp{Role: &permissionrpc.Role{Id: 1, RoleKey: "editor"}},
	}
	ctx := context.Background()

	addLogic := NewAddRoleLogic(ctx, &svc.ServiceContext{PermissionRpc: rpc})
	addResp, err := addLogic.AddRole(&types.NewRoleReq{
		Id:          1,
		ParentId:    2,
		RoleKey:     "admin",
		RoleLabel:   "Admin",
		RoleComment: "root",
		Status:      1,
		IsDefault:   true,
	})
	if err != nil {
		t.Fatalf("AddRole returned error: %v", err)
	}
	if rpc.addReq == nil || rpc.addReq.RoleKey != "admin" || addResp.Id != 1 {
		t.Fatalf("unexpected add flow: req=%#v resp=%#v", rpc.addReq, addResp)
	}

	updateLogic := NewUpdateRoleLogic(ctx, &svc.ServiceContext{PermissionRpc: rpc})
	updateResp, err := updateLogic.UpdateRole(&types.NewRoleReq{
		Id:          1,
		RoleKey:     "editor",
		RoleLabel:   "Editor",
		RoleComment: "edit",
		Status:      0,
	})
	if err != nil {
		t.Fatalf("UpdateRole returned error: %v", err)
	}
	if rpc.updateReq == nil || rpc.updateReq.RoleKey != "editor" || updateResp.RoleKey != "editor" {
		t.Fatalf("unexpected update flow: req=%#v resp=%#v", rpc.updateReq, updateResp)
	}
}

func TestFindRoleListAndResources(t *testing.T) {
	rpc := &stubRolePermissionRPC{
		findResp: &permissionrpc.FindRoleListResp{
			Pagination: &permissionrpc.PageResp{Page: 1, PageSize: 10, Total: 1},
			List: []*permissionrpc.Role{
				{Id: 1, RoleKey: "admin"},
			},
		},
		resResp: &permissionrpc.FindRoleResourcesResp{
			RoleId:  1,
			ApiIds:  []int64{2},
			MenuIds: []int64{3},
		},
	}
	ctx := context.Background()

	findLogic := NewFindRoleListLogic(ctx, &svc.ServiceContext{PermissionRpc: rpc})
	findResp, err := findLogic.FindRoleList(&types.QueryRoleReq{
		PageQuery: types.PageQuery{Page: 1, PageSize: 10},
		RoleKey:   "admin",
	})
	if err != nil {
		t.Fatalf("FindRoleList returned error: %v", err)
	}
	if rpc.findReq == nil || rpc.findReq.RoleKey != "admin" {
		t.Fatalf("unexpected find request: %#v", rpc.findReq)
	}
	list, ok := findResp.List.([]*types.RoleBackVO)
	if !ok || len(list) != 1 || list[0].Id != 1 {
		t.Fatalf("unexpected role list: %#v", findResp.List)
	}

	resLogic := NewFindRoleResourcesLogic(ctx, &svc.ServiceContext{PermissionRpc: rpc})
	resResp, err := resLogic.FindRoleResources(&types.IdReq{Id: 1})
	if err != nil {
		t.Fatalf("FindRoleResources returned error: %v", err)
	}
	if rpc.resReq == nil || rpc.resReq.Id != 1 || resResp.RoleId != 1 || len(resResp.ApiIds) != 1 {
		t.Fatalf("unexpected resources flow: req=%#v resp=%#v", rpc.resReq, resResp)
	}
}

func TestUpdateRolePoliciesAndDelete(t *testing.T) {
	rpc := &stubRolePermissionRPC{
		deleteResp: &permissionrpc.DeletesRoleResp{SuccessCount: 1},
	}
	holder := &stubPermissionHolder{}
	ctx := context.Background()

	apiLogic := NewUpdateRoleApisLogic(ctx, &svc.ServiceContext{
		PermissionRpc:    rpc,
		PermissionHolder: holder,
	})
	if _, err := apiLogic.UpdateRoleApis(&types.UpdateRoleApisReq{RoleId: 1, ApiIds: []int64{2, 3}}); err != nil {
		t.Fatalf("UpdateRoleApis returned error: %v", err)
	}
	if rpc.updateApisReq == nil || rpc.updateApisReq.RoleId != 1 || holder.reloadCount != 1 {
		t.Fatalf("unexpected update apis flow: req=%#v reload=%d", rpc.updateApisReq, holder.reloadCount)
	}

	menuLogic := NewUpdateRoleMenusLogic(ctx, &svc.ServiceContext{
		PermissionRpc:    rpc,
		PermissionHolder: holder,
	})
	if _, err := menuLogic.UpdateRoleMenus(&types.UpdateRoleMenusReq{RoleId: 1, MenuIds: []int64{4}}); err != nil {
		t.Fatalf("UpdateRoleMenus returned error: %v", err)
	}
	if rpc.updateMenusReq == nil || rpc.updateMenusReq.RoleId != 1 || holder.reloadCount != 2 {
		t.Fatalf("unexpected update menus flow: req=%#v reload=%d", rpc.updateMenusReq, holder.reloadCount)
	}

	deleteLogic := NewDeletesRoleLogic(ctx, &svc.ServiceContext{PermissionRpc: rpc})
	deleteResp, err := deleteLogic.DeletesRole(&types.IdsReq{Ids: []int64{1}})
	if err != nil {
		t.Fatalf("DeletesRole returned error: %v", err)
	}
	if rpc.deleteReq == nil || len(rpc.deleteReq.Ids) != 1 || deleteResp.SuccessCount != 1 {
		t.Fatalf("unexpected delete flow: req=%#v resp=%#v", rpc.deleteReq, deleteResp)
	}
}
