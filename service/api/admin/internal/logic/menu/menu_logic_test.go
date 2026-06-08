package menu

import (
	"context"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/permissionrpc"
	"google.golang.org/grpc"
)

type stubMenuPermissionRPC struct {
	permissionrpc.PermissionRpc
	addReq     *permissionrpc.AddMenuReq
	addResp    *permissionrpc.AddMenuResp
	updateReq  *permissionrpc.UpdateMenuReq
	updateResp *permissionrpc.UpdateMenuResp
	deleteReq  *permissionrpc.DeletesMenuReq
	deleteResp *permissionrpc.DeletesMenuResp
	findReq    *permissionrpc.FindMenuListReq
	findResp   *permissionrpc.FindMenuListResp
	syncReq    *permissionrpc.SyncMenuListReq
	syncResp   *permissionrpc.SyncMenuListResp
	cleanReq   *permissionrpc.CleanMenuListReq
	cleanResp  *permissionrpc.CleanMenuListResp
}

func (s *stubMenuPermissionRPC) AddMenu(_ context.Context, in *permissionrpc.AddMenuReq, _ ...grpc.CallOption) (*permissionrpc.AddMenuResp, error) {
	s.addReq = in
	return s.addResp, nil
}

func (s *stubMenuPermissionRPC) UpdateMenu(_ context.Context, in *permissionrpc.UpdateMenuReq, _ ...grpc.CallOption) (*permissionrpc.UpdateMenuResp, error) {
	s.updateReq = in
	return s.updateResp, nil
}

func (s *stubMenuPermissionRPC) DeletesMenu(_ context.Context, in *permissionrpc.DeletesMenuReq, _ ...grpc.CallOption) (*permissionrpc.DeletesMenuResp, error) {
	s.deleteReq = in
	return s.deleteResp, nil
}

func (s *stubMenuPermissionRPC) FindMenuList(_ context.Context, in *permissionrpc.FindMenuListReq, _ ...grpc.CallOption) (*permissionrpc.FindMenuListResp, error) {
	s.findReq = in
	return s.findResp, nil
}

func (s *stubMenuPermissionRPC) SyncMenuList(_ context.Context, in *permissionrpc.SyncMenuListReq, _ ...grpc.CallOption) (*permissionrpc.SyncMenuListResp, error) {
	s.syncReq = in
	return s.syncResp, nil
}

func (s *stubMenuPermissionRPC) CleanMenuList(_ context.Context, in *permissionrpc.CleanMenuListReq, _ ...grpc.CallOption) (*permissionrpc.CleanMenuListResp, error) {
	s.cleanReq = in
	return s.cleanResp, nil
}

func TestAddAndUpdateMenuBuildRequests(t *testing.T) {
	rpc := &stubMenuPermissionRPC{
		addResp: &permissionrpc.AddMenuResp{
			Menu: &permissionrpc.Menu{
				Id:       1,
				Name:     "menu-a",
				Path:     "/a",
				Meta:     &permissionrpc.MenuMeta{Title: "Menu A", Params: `[{"key":"id","value":"1"}]`},
				Children: []*permissionrpc.Menu{{Id: 2, Name: "child", Meta: &permissionrpc.MenuMeta{Title: "Child"}}},
			},
		},
		updateResp: &permissionrpc.UpdateMenuResp{
			Menu: &permissionrpc.Menu{
				Id:   1,
				Name: "menu-b",
				Meta: &permissionrpc.MenuMeta{Title: "Menu B"},
			},
		},
	}
	ctx := context.Background()

	addLogic := NewAddMenuLogic(ctx, &svc.ServiceContext{PermissionRpc: rpc})
	addResp, err := addLogic.AddMenu(&types.NewMenuReq{
		Id:        1,
		ParentId:  0,
		Path:      "/a",
		Name:      "menu-a",
		Component: "Layout",
		Redirect:  "/a/index",
		MenuMeta: types.MenuMeta{
			Type:       "1",
			Title:      "Menu A",
			Icon:       "home",
			Rank:       1,
			Perm:       "system:menu:add",
			Params:     []*types.MenuMetaParams{{Key: "id", Value: "1"}},
			KeepAlive:  1,
			AlwaysShow: 1,
			Visible:    1,
			Status:     1,
		},
		Children: []*types.NewMenuReq{{Id: 2, Name: "child", MenuMeta: types.MenuMeta{Title: "Child"}}},
	})
	if err != nil {
		t.Fatalf("AddMenu returned error: %v", err)
	}
	if rpc.addReq == nil || rpc.addReq.Name != "menu-a" || rpc.addReq.Meta == nil || rpc.addReq.Meta.Title != "Menu A" || len(rpc.addReq.Children) != 1 {
		t.Fatalf("unexpected add request: %#v", rpc.addReq)
	}
	if addResp.Name != "menu-a" || len(addResp.Children) != 1 || len(addResp.Params) != 1 {
		t.Fatalf("unexpected add response: %#v", addResp)
	}

	updateLogic := NewUpdateMenuLogic(ctx, &svc.ServiceContext{PermissionRpc: rpc})
	updateResp, err := updateLogic.UpdateMenu(&types.NewMenuReq{
		Id:   1,
		Name: "menu-b",
		MenuMeta: types.MenuMeta{
			Title: "Menu B",
		},
	})
	if err != nil {
		t.Fatalf("UpdateMenu returned error: %v", err)
	}
	if rpc.updateReq == nil || rpc.updateReq.Name != "menu-b" || rpc.updateReq.Meta == nil || rpc.updateReq.Meta.Title != "Menu B" {
		t.Fatalf("unexpected update request: %#v", rpc.updateReq)
	}
	if updateResp.Name != "menu-b" {
		t.Fatalf("unexpected update response: %#v", updateResp)
	}
}

func TestFindMenuListBuildsRequestAndMapsTree(t *testing.T) {
	rpc := &stubMenuPermissionRPC{
		findResp: &permissionrpc.FindMenuListResp{
			List: []*permissionrpc.Menu{
				{
					Id:   1,
					Name: "menu-a",
					Meta: &permissionrpc.MenuMeta{Title: "Menu A", Params: `[{"key":"id","value":"1"}]`},
					Children: []*permissionrpc.Menu{
						{Id: 2, Name: "child", Meta: &permissionrpc.MenuMeta{Title: "Child"}},
					},
				},
			},
		},
	}
	logic := NewFindMenuListLogic(context.Background(), &svc.ServiceContext{PermissionRpc: rpc})

	resp, err := logic.FindMenuList(&types.QueryMenuReq{Name: "menu", Title: "Menu"})
	if err != nil {
		t.Fatalf("FindMenuList returned error: %v", err)
	}
	if rpc.findReq == nil || rpc.findReq.Name != "menu" || rpc.findReq.Title != "Menu" {
		t.Fatalf("unexpected find request: %#v", rpc.findReq)
	}
	list, ok := resp.List.([]*types.MenuBackVO)
	if !ok || len(list) != 1 || len(list[0].Children) != 1 || len(list[0].Params) != 1 {
		t.Fatalf("unexpected menu list response: %#v", resp)
	}
}

func TestSyncCleanAndDeleteMenuBuildRequests(t *testing.T) {
	rpc := &stubMenuPermissionRPC{
		deleteResp: &permissionrpc.DeletesMenuResp{SuccessCount: 2},
		syncResp:   &permissionrpc.SyncMenuListResp{SuccessCount: 3},
		cleanResp:  &permissionrpc.CleanMenuListResp{SuccessCount: 4},
	}
	ctx := context.Background()

	deleteLogic := NewDeletesMenuLogic(ctx, &svc.ServiceContext{PermissionRpc: rpc})
	deleteResp, err := deleteLogic.DeletesMenu(&types.IdsReq{Ids: []int64{1, 2}})
	if err != nil {
		t.Fatalf("DeletesMenu returned error: %v", err)
	}
	if rpc.deleteReq == nil || len(rpc.deleteReq.Ids) != 2 || deleteResp.SuccessCount != 2 {
		t.Fatalf("unexpected delete flow: req=%#v resp=%#v", rpc.deleteReq, deleteResp)
	}

	syncLogic := NewSyncMenuListLogic(ctx, &svc.ServiceContext{PermissionRpc: rpc})
	syncResp, err := syncLogic.SyncMenuList(&types.SyncMenuReq{
		Menus: []*types.NewMenuReq{
			{Name: "root", MenuMeta: types.MenuMeta{Title: "Root"}},
			{Name: "child", MenuMeta: types.MenuMeta{Title: "Child"}},
		},
	})
	if err != nil {
		t.Fatalf("SyncMenuList returned error: %v", err)
	}
	if rpc.syncReq == nil || len(rpc.syncReq.Menus) != 2 || syncResp.SuccessCount != 3 {
		t.Fatalf("unexpected sync flow: req=%#v resp=%#v", rpc.syncReq, syncResp)
	}

	cleanLogic := NewCleanMenuListLogic(ctx, &svc.ServiceContext{PermissionRpc: rpc})
	cleanResp, err := cleanLogic.CleanMenuList(&types.CleanMenuReq{})
	if err != nil {
		t.Fatalf("CleanMenuList returned error: %v", err)
	}
	if rpc.cleanReq == nil || cleanResp.SuccessCount != 4 {
		t.Fatalf("unexpected clean flow: req=%#v resp=%#v", rpc.cleanReq, cleanResp)
	}
}
