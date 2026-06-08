package page

import (
	"context"
	"errors"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/resourcerpc"
	"google.golang.org/grpc"
)

type stubResourceRPC struct {
	addPageReq     *resourcerpc.AddPageReq
	addPageResp    *resourcerpc.AddPageResp
	addPageErr     error
	updatePageReq  *resourcerpc.UpdatePageReq
	updatePageResp *resourcerpc.UpdatePageResp
	updatePageErr  error
	deletePageReq  *resourcerpc.DeletesPageReq
	deletePageResp *resourcerpc.DeletesPageResp
	deletePageErr  error
	findPageReq    *resourcerpc.FindPageListReq
	findPageResp   *resourcerpc.FindPageListResp
	findPageErr    error
}

func (s *stubResourceRPC) AddPage(_ context.Context, in *resourcerpc.AddPageReq, _ ...grpc.CallOption) (*resourcerpc.AddPageResp, error) {
	s.addPageReq = in
	return s.addPageResp, s.addPageErr
}

func (s *stubResourceRPC) UpdatePage(_ context.Context, in *resourcerpc.UpdatePageReq, _ ...grpc.CallOption) (*resourcerpc.UpdatePageResp, error) {
	s.updatePageReq = in
	return s.updatePageResp, s.updatePageErr
}

func (s *stubResourceRPC) DeletesPage(_ context.Context, in *resourcerpc.DeletesPageReq, _ ...grpc.CallOption) (*resourcerpc.DeletesPageResp, error) {
	s.deletePageReq = in
	return s.deletePageResp, s.deletePageErr
}

func (s *stubResourceRPC) FindPageList(_ context.Context, in *resourcerpc.FindPageListReq, _ ...grpc.CallOption) (*resourcerpc.FindPageListResp, error) {
	s.findPageReq = in
	return s.findPageResp, s.findPageErr
}

func (s *stubResourceRPC) AddPhoto(context.Context, *resourcerpc.AddPhotoReq, ...grpc.CallOption) (*resourcerpc.AddPhotoResp, error) {
	panic("unexpected call")
}

func (s *stubResourceRPC) UpdatePhoto(context.Context, *resourcerpc.UpdatePhotoReq, ...grpc.CallOption) (*resourcerpc.UpdatePhotoResp, error) {
	panic("unexpected call")
}

func (s *stubResourceRPC) UpdatePhotoDelete(context.Context, *resourcerpc.UpdatePhotoDeleteReq, ...grpc.CallOption) (*resourcerpc.UpdatePhotoDeleteResp, error) {
	panic("unexpected call")
}

func (s *stubResourceRPC) DeletesPhoto(context.Context, *resourcerpc.DeletesPhotoReq, ...grpc.CallOption) (*resourcerpc.DeletesPhotoResp, error) {
	panic("unexpected call")
}

func (s *stubResourceRPC) FindPhotoList(context.Context, *resourcerpc.FindPhotoListReq, ...grpc.CallOption) (*resourcerpc.FindPhotoListResp, error) {
	panic("unexpected call")
}

func (s *stubResourceRPC) AddAlbum(context.Context, *resourcerpc.AddAlbumReq, ...grpc.CallOption) (*resourcerpc.AddAlbumResp, error) {
	panic("unexpected call")
}

func (s *stubResourceRPC) UpdateAlbum(context.Context, *resourcerpc.UpdateAlbumReq, ...grpc.CallOption) (*resourcerpc.UpdateAlbumResp, error) {
	panic("unexpected call")
}

func (s *stubResourceRPC) UpdateAlbumDelete(context.Context, *resourcerpc.UpdateAlbumDeleteReq, ...grpc.CallOption) (*resourcerpc.UpdateAlbumDeleteResp, error) {
	panic("unexpected call")
}

func (s *stubResourceRPC) GetAlbum(context.Context, *resourcerpc.GetAlbumReq, ...grpc.CallOption) (*resourcerpc.GetAlbumResp, error) {
	panic("unexpected call")
}

func (s *stubResourceRPC) DeletesAlbum(context.Context, *resourcerpc.DeletesAlbumReq, ...grpc.CallOption) (*resourcerpc.DeletesAlbumResp, error) {
	panic("unexpected call")
}

func (s *stubResourceRPC) FindAlbumList(context.Context, *resourcerpc.FindAlbumListReq, ...grpc.CallOption) (*resourcerpc.FindAlbumListResp, error) {
	panic("unexpected call")
}

func TestAddPageBuildsRequest(t *testing.T) {
	rpc := &stubResourceRPC{
		addPageResp: &resourcerpc.AddPageResp{
			Page: &resourcerpc.Page{Id: 2, PageName: "about"},
		},
	}
	logic := NewAddPageLogic(context.Background(), &svc.ServiceContext{ResourceRpc: rpc})

	resp, err := logic.AddPage(&types.NewPageReq{
		Id:             2,
		PageName:       "about",
		PageLabel:      "About",
		PageCover:      "cover",
		IsCarousel:     1,
		CarouselCovers: []string{"a", "b"},
	})
	if err != nil {
		t.Fatalf("AddPage returned error: %v", err)
	}

	if rpc.addPageReq == nil || rpc.addPageReq.PageName != "about" || len(rpc.addPageReq.CarouselCovers) != 2 {
		t.Fatalf("unexpected add request: %#v", rpc.addPageReq)
	}

	if resp.Id != 2 || resp.PageName != "about" {
		t.Fatalf("unexpected response: %#v", resp)
	}
}

func TestDeletePageBuildsDeleteRequest(t *testing.T) {
	rpc := &stubResourceRPC{
		deletePageResp: &resourcerpc.DeletesPageResp{SuccessCount: 1},
	}
	logic := NewDeletePageLogic(context.Background(), &svc.ServiceContext{ResourceRpc: rpc})

	resp, err := logic.DeletePage(&types.IdReq{Id: 9})
	if err != nil {
		t.Fatalf("DeletePage returned error: %v", err)
	}

	if rpc.deletePageReq == nil || len(rpc.deletePageReq.Ids) != 1 || rpc.deletePageReq.Ids[0] != 9 {
		t.Fatalf("unexpected delete request: %#v", rpc.deletePageReq)
	}

	if resp.SuccessCount != 1 {
		t.Fatalf("unexpected success count: %d", resp.SuccessCount)
	}
}

func TestFindPageListPropagatesRPCError(t *testing.T) {
	wantErr := errors.New("find failed")
	logic := NewFindPageListLogic(context.Background(), &svc.ServiceContext{
		ResourceRpc: &stubResourceRPC{findPageErr: wantErr},
	})

	_, err := logic.FindPageList(&types.QueryPageReq{})
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected %v, got %v", wantErr, err)
	}
}

func TestFindPageListBuildsQueryAndMapsPage(t *testing.T) {
	rpc := &stubResourceRPC{
		findPageResp: &resourcerpc.FindPageListResp{
			Pagination: &resourcerpc.PageResp{Page: 3, PageSize: 5, Total: 11},
			List: []*resourcerpc.Page{
				{Id: 4, PageName: "index", CarouselCovers: []string{"x"}},
			},
		},
	}
	logic := NewFindPageListLogic(context.Background(), &svc.ServiceContext{ResourceRpc: rpc})

	resp, err := logic.FindPageList(&types.QueryPageReq{
		PageQuery: types.PageQuery{
			Page:     3,
			PageSize: 5,
			Sorts:    []string{"created_at desc"},
		},
		PageName: "ind",
	})
	if err != nil {
		t.Fatalf("FindPageList returned error: %v", err)
	}

	if rpc.findPageReq == nil || rpc.findPageReq.Paginate == nil {
		t.Fatal("expected find request to be sent")
	}

	if rpc.findPageReq.Paginate.Page != 3 || rpc.findPageReq.Paginate.PageSize != 5 || rpc.findPageReq.PageName != "ind" {
		t.Fatalf("unexpected find request: %#v", rpc.findPageReq)
	}

	list, ok := resp.List.([]*types.PageBackVO)
	if !ok || len(list) != 1 || list[0].Id != 4 {
		t.Fatalf("unexpected list payload: %#v", resp.List)
	}
}
