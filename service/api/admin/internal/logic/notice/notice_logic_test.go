package notice

import (
	"context"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/noticerpc"
	"google.golang.org/grpc"
)

type stubNoticeRPC struct {
	addReq             *noticerpc.AddNoticeReq
	addResp            *noticerpc.AddNoticeResp
	updateReq          *noticerpc.UpdateNoticeReq
	updateResp         *noticerpc.UpdateNoticeResp
	getReq             *noticerpc.GetNoticeReq
	getResp            *noticerpc.GetNoticeResp
	deleteReq          *noticerpc.DeletesNoticeReq
	deleteResp         *noticerpc.DeletesNoticeResp
	findReq            *noticerpc.FindNoticeListReq
	findResp           *noticerpc.FindNoticeListResp
	updateStatusReq    *noticerpc.UpdateNoticeStatusReq
	updateStatusResp   *noticerpc.UpdateNoticeStatusResp
	findUserNoticeReq  *noticerpc.FindUserNoticeListReq
	findUserNoticeResp *noticerpc.FindUserNoticeListResp
}

func (s *stubNoticeRPC) AddNotice(_ context.Context, in *noticerpc.AddNoticeReq, _ ...grpc.CallOption) (*noticerpc.AddNoticeResp, error) {
	s.addReq = in
	return s.addResp, nil
}

func (s *stubNoticeRPC) UpdateNotice(_ context.Context, in *noticerpc.UpdateNoticeReq, _ ...grpc.CallOption) (*noticerpc.UpdateNoticeResp, error) {
	s.updateReq = in
	return s.updateResp, nil
}

func (s *stubNoticeRPC) GetNotice(_ context.Context, in *noticerpc.GetNoticeReq, _ ...grpc.CallOption) (*noticerpc.GetNoticeResp, error) {
	s.getReq = in
	return s.getResp, nil
}

func (s *stubNoticeRPC) DeletesNotice(_ context.Context, in *noticerpc.DeletesNoticeReq, _ ...grpc.CallOption) (*noticerpc.DeletesNoticeResp, error) {
	s.deleteReq = in
	return s.deleteResp, nil
}

func (s *stubNoticeRPC) FindNoticeList(_ context.Context, in *noticerpc.FindNoticeListReq, _ ...grpc.CallOption) (*noticerpc.FindNoticeListResp, error) {
	s.findReq = in
	return s.findResp, nil
}

func (s *stubNoticeRPC) UpdateNoticeStatus(_ context.Context, in *noticerpc.UpdateNoticeStatusReq, _ ...grpc.CallOption) (*noticerpc.UpdateNoticeStatusResp, error) {
	s.updateStatusReq = in
	return s.updateStatusResp, nil
}

func (s *stubNoticeRPC) FindUserNoticeList(_ context.Context, in *noticerpc.FindUserNoticeListReq, _ ...grpc.CallOption) (*noticerpc.FindUserNoticeListResp, error) {
	s.findUserNoticeReq = in
	return s.findUserNoticeResp, nil
}

func TestAddNoticeBuildsRequest(t *testing.T) {
	rpc := &stubNoticeRPC{
		addResp: &noticerpc.AddNoticeResp{
			Notice: &noticerpc.Notice{Id: 1, Title: "title"},
		},
	}
	logic := NewAddNoticeLogic(context.Background(), &svc.ServiceContext{NoticeRpc: rpc})

	resp, err := logic.AddNotice(&types.AddNoticeReq{
		Title:   "title",
		Content: "content",
		Type:    "system",
		Level:   "info",
		AppName: "admin",
	})
	if err != nil {
		t.Fatalf("AddNotice returned error: %v", err)
	}
	if rpc.addReq == nil || rpc.addReq.Title != "title" || rpc.addReq.AppName != "admin" {
		t.Fatalf("unexpected add request: %#v", rpc.addReq)
	}
	if resp.Id != 1 || resp.Title != "title" {
		t.Fatalf("unexpected response: %#v", resp)
	}
}

func TestFindNoticeListBuildsQuery(t *testing.T) {
	rpc := &stubNoticeRPC{
		findResp: &noticerpc.FindNoticeListResp{
			Pagination: &noticerpc.PageResp{Page: 1, PageSize: 10, Total: 1},
			List: []*noticerpc.Notice{
				{Id: 1, Title: "title", Type: "system"},
			},
		},
	}
	logic := NewFindNoticeListLogic(context.Background(), &svc.ServiceContext{NoticeRpc: rpc})

	resp, err := logic.FindNoticeList(&types.QueryNoticeReq{
		PageQuery:     types.PageQuery{Page: 1, PageSize: 10},
		Type:          "system",
		Level:         "info",
		PublishStatus: 2,
		AppName:       "admin",
	})
	if err != nil {
		t.Fatalf("FindNoticeList returned error: %v", err)
	}
	if rpc.findReq == nil || rpc.findReq.Type != "system" || rpc.findReq.PublishStatus != 2 {
		t.Fatalf("unexpected find request: %#v", rpc.findReq)
	}
	list, ok := resp.List.([]*types.NoticeBackVO)
	if !ok || len(list) != 1 || list[0].Id != 1 {
		t.Fatalf("unexpected list payload: %#v", resp.List)
	}
}

func TestFindUserNoticeListBuildsQuery(t *testing.T) {
	rpc := &stubNoticeRPC{
		findUserNoticeResp: &noticerpc.FindUserNoticeListResp{
			Pagination: &noticerpc.PageResp{Page: 2, PageSize: 5, Total: 1},
			List: []*noticerpc.Notice{
				{Id: 9, Title: "notice"},
			},
		},
	}
	logic := NewFindUserNoticeListLogic(context.Background(), &svc.ServiceContext{NoticeRpc: rpc})

	resp, err := logic.FindUserNoticeList(&types.QueryUserNoticeReq{
		PageQuery: types.PageQuery{Page: 2, PageSize: 5},
	})
	if err != nil {
		t.Fatalf("FindUserNoticeList returned error: %v", err)
	}
	if rpc.findUserNoticeReq == nil || rpc.findUserNoticeReq.Paginate == nil || rpc.findUserNoticeReq.Paginate.Page != 2 {
		t.Fatalf("unexpected user notice request: %#v", rpc.findUserNoticeReq)
	}
	list, ok := resp.List.([]*types.NoticeBackVO)
	if !ok || len(list) != 1 || list[0].Id != 9 {
		t.Fatalf("unexpected list payload: %#v", resp.List)
	}
}

func TestGetUpdateDeleteNotice(t *testing.T) {
	rpc := &stubNoticeRPC{
		getResp:          &noticerpc.GetNoticeResp{Notice: &noticerpc.Notice{Id: 7, Title: "detail"}},
		updateResp:       &noticerpc.UpdateNoticeResp{Notice: &noticerpc.Notice{Id: 7, Title: "new"}},
		updateStatusResp: &noticerpc.UpdateNoticeStatusResp{Notice: &noticerpc.Notice{Id: 7, PublishStatus: 2}},
		deleteResp:       &noticerpc.DeletesNoticeResp{SuccessCount: 1},
	}
	ctx := context.Background()

	getLogic := NewGetNoticeLogic(ctx, &svc.ServiceContext{NoticeRpc: rpc})
	getResp, err := getLogic.GetNotice(&types.IdReq{Id: 7})
	if err != nil {
		t.Fatalf("GetNotice returned error: %v", err)
	}
	if rpc.getReq == nil || rpc.getReq.Id != 7 || getResp.Title != "detail" {
		t.Fatalf("unexpected get flow: req=%#v resp=%#v", rpc.getReq, getResp)
	}

	updateLogic := NewUpdateNoticeLogic(ctx, &svc.ServiceContext{NoticeRpc: rpc})
	updateResp, err := updateLogic.UpdateNotice(&types.UpdateNoticeReq{
		Id:      7,
		Title:   "new",
		Content: "content",
		Type:    "system",
		Level:   "info",
		AppName: "admin",
	})
	if err != nil {
		t.Fatalf("UpdateNotice returned error: %v", err)
	}
	if rpc.updateReq == nil || rpc.updateReq.Id != 7 || updateResp.Title != "new" {
		t.Fatalf("unexpected update flow: req=%#v resp=%#v", rpc.updateReq, updateResp)
	}

	statusLogic := NewUpdateNoticeStatusLogic(ctx, &svc.ServiceContext{NoticeRpc: rpc})
	statusResp, err := statusLogic.UpdateNoticeStatus(&types.UpdateNoticeStatusReq{Id: 7, PublishStatus: 2})
	if err != nil {
		t.Fatalf("UpdateNoticeStatus returned error: %v", err)
	}
	if rpc.updateStatusReq == nil || rpc.updateStatusReq.PublishStatus != 2 || statusResp.PublishStatus != 2 {
		t.Fatalf("unexpected status flow: req=%#v resp=%#v", rpc.updateStatusReq, statusResp)
	}

	deleteLogic := NewDeletesNoticeLogic(ctx, &svc.ServiceContext{NoticeRpc: rpc})
	deleteResp, err := deleteLogic.DeletesNotice(&types.IdsReq{Ids: []int64{7}})
	if err != nil {
		t.Fatalf("DeletesNotice returned error: %v", err)
	}
	if rpc.deleteReq == nil || len(rpc.deleteReq.Ids) != 1 || deleteResp.SuccessCount != 1 {
		t.Fatalf("unexpected delete flow: req=%#v resp=%#v", rpc.deleteReq, deleteResp)
	}
}
