package photo

import (
	"context"
	"errors"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/resourcerpc"
	"google.golang.org/grpc"
)

type stubPhotoResourceRPC struct {
	addPhotoReq           *resourcerpc.AddPhotoReq
	addPhotoResp          *resourcerpc.AddPhotoResp
	addPhotoErr           error
	updatePhotoReq        *resourcerpc.UpdatePhotoReq
	updatePhotoResp       *resourcerpc.UpdatePhotoResp
	updatePhotoErr        error
	updatePhotoDeleteReq  *resourcerpc.UpdatePhotoDeleteReq
	updatePhotoDeleteResp *resourcerpc.UpdatePhotoDeleteResp
	updatePhotoDeleteErr  error
	deletePhotoReq        *resourcerpc.DeletesPhotoReq
	deletePhotoResp       *resourcerpc.DeletesPhotoResp
	deletePhotoErr        error
	findPhotoReq          *resourcerpc.FindPhotoListReq
	findPhotoResp         *resourcerpc.FindPhotoListResp
	findPhotoErr          error
}

func (s *stubPhotoResourceRPC) AddPhoto(_ context.Context, in *resourcerpc.AddPhotoReq, _ ...grpc.CallOption) (*resourcerpc.AddPhotoResp, error) {
	s.addPhotoReq = in
	return s.addPhotoResp, s.addPhotoErr
}

func (s *stubPhotoResourceRPC) UpdatePhoto(_ context.Context, in *resourcerpc.UpdatePhotoReq, _ ...grpc.CallOption) (*resourcerpc.UpdatePhotoResp, error) {
	s.updatePhotoReq = in
	return s.updatePhotoResp, s.updatePhotoErr
}

func (s *stubPhotoResourceRPC) UpdatePhotoDelete(_ context.Context, in *resourcerpc.UpdatePhotoDeleteReq, _ ...grpc.CallOption) (*resourcerpc.UpdatePhotoDeleteResp, error) {
	s.updatePhotoDeleteReq = in
	return s.updatePhotoDeleteResp, s.updatePhotoDeleteErr
}

func (s *stubPhotoResourceRPC) DeletesPhoto(_ context.Context, in *resourcerpc.DeletesPhotoReq, _ ...grpc.CallOption) (*resourcerpc.DeletesPhotoResp, error) {
	s.deletePhotoReq = in
	return s.deletePhotoResp, s.deletePhotoErr
}

func (s *stubPhotoResourceRPC) FindPhotoList(_ context.Context, in *resourcerpc.FindPhotoListReq, _ ...grpc.CallOption) (*resourcerpc.FindPhotoListResp, error) {
	s.findPhotoReq = in
	return s.findPhotoResp, s.findPhotoErr
}

func (s *stubPhotoResourceRPC) AddAlbum(context.Context, *resourcerpc.AddAlbumReq, ...grpc.CallOption) (*resourcerpc.AddAlbumResp, error) {
	panic("unexpected call")
}

func (s *stubPhotoResourceRPC) UpdateAlbum(context.Context, *resourcerpc.UpdateAlbumReq, ...grpc.CallOption) (*resourcerpc.UpdateAlbumResp, error) {
	panic("unexpected call")
}

func (s *stubPhotoResourceRPC) UpdateAlbumDelete(context.Context, *resourcerpc.UpdateAlbumDeleteReq, ...grpc.CallOption) (*resourcerpc.UpdateAlbumDeleteResp, error) {
	panic("unexpected call")
}

func (s *stubPhotoResourceRPC) GetAlbum(context.Context, *resourcerpc.GetAlbumReq, ...grpc.CallOption) (*resourcerpc.GetAlbumResp, error) {
	panic("unexpected call")
}

func (s *stubPhotoResourceRPC) DeletesAlbum(context.Context, *resourcerpc.DeletesAlbumReq, ...grpc.CallOption) (*resourcerpc.DeletesAlbumResp, error) {
	panic("unexpected call")
}

func (s *stubPhotoResourceRPC) FindAlbumList(context.Context, *resourcerpc.FindAlbumListReq, ...grpc.CallOption) (*resourcerpc.FindAlbumListResp, error) {
	panic("unexpected call")
}

func (s *stubPhotoResourceRPC) AddPage(context.Context, *resourcerpc.AddPageReq, ...grpc.CallOption) (*resourcerpc.AddPageResp, error) {
	panic("unexpected call")
}

func (s *stubPhotoResourceRPC) UpdatePage(context.Context, *resourcerpc.UpdatePageReq, ...grpc.CallOption) (*resourcerpc.UpdatePageResp, error) {
	panic("unexpected call")
}

func (s *stubPhotoResourceRPC) DeletesPage(context.Context, *resourcerpc.DeletesPageReq, ...grpc.CallOption) (*resourcerpc.DeletesPageResp, error) {
	panic("unexpected call")
}

func (s *stubPhotoResourceRPC) FindPageList(context.Context, *resourcerpc.FindPageListReq, ...grpc.CallOption) (*resourcerpc.FindPageListResp, error) {
	panic("unexpected call")
}

func TestAddPhotoBuildsRequest(t *testing.T) {
	rpc := &stubPhotoResourceRPC{
		addPhotoResp: &resourcerpc.AddPhotoResp{
			Photo: &resourcerpc.Photo{Id: 1, AlbumId: 2, PhotoName: "cover"},
		},
	}
	logic := NewAddPhotoLogic(context.Background(), &svc.ServiceContext{ResourceRpc: rpc})

	resp, err := logic.AddPhoto(&types.NewPhotoReq{
		Id:        1,
		AlbumId:   2,
		PhotoName: "cover",
		PhotoDesc: "desc",
		PhotoSrc:  "src",
		IsDelete:  true,
	})
	if err != nil {
		t.Fatalf("AddPhoto returned error: %v", err)
	}

	if rpc.addPhotoReq == nil || rpc.addPhotoReq.AlbumId != 2 || rpc.addPhotoReq.PhotoName != "cover" {
		t.Fatalf("unexpected add request: %#v", rpc.addPhotoReq)
	}

	if resp.Id != 1 || resp.AlbumId != 2 {
		t.Fatalf("unexpected response: %#v", resp)
	}
}

func TestUpdatePhotoDeleteBuildsRequest(t *testing.T) {
	rpc := &stubPhotoResourceRPC{
		updatePhotoDeleteResp: &resourcerpc.UpdatePhotoDeleteResp{SuccessCount: 2},
	}
	logic := NewUpdatePhotoDeleteLogic(context.Background(), &svc.ServiceContext{ResourceRpc: rpc})

	resp, err := logic.UpdatePhotoDelete(&types.UpdatePhotoDeleteReq{
		Ids:      []int64{1, 2},
		IsDelete: true,
	})
	if err != nil {
		t.Fatalf("UpdatePhotoDelete returned error: %v", err)
	}

	if rpc.updatePhotoDeleteReq == nil || !rpc.updatePhotoDeleteReq.IsDelete || len(rpc.updatePhotoDeleteReq.Ids) != 2 {
		t.Fatalf("unexpected update delete request: %#v", rpc.updatePhotoDeleteReq)
	}

	if resp.SuccessCount != 2 {
		t.Fatalf("unexpected success count: %d", resp.SuccessCount)
	}
}

func TestFindPhotoListPropagatesRPCError(t *testing.T) {
	wantErr := errors.New("find failed")
	logic := NewFindPhotoListLogic(context.Background(), &svc.ServiceContext{
		ResourceRpc: &stubPhotoResourceRPC{findPhotoErr: wantErr},
	})

	_, err := logic.FindPhotoList(&types.QueryPhotoReq{})
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected %v, got %v", wantErr, err)
	}
}
