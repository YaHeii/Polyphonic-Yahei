package comment

import (
	"context"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/newsrpc"
	"google.golang.org/grpc"
)

type stubCommentNewsRPC struct {
	newsrpc.NewsRpc
	deleteReq  *newsrpc.DeletesCommentReq
	deleteResp *newsrpc.DeletesCommentResp
	deleteErr  error
	findReq    *newsrpc.FindCommentListReq
	findResp   *newsrpc.FindCommentListResp
	findErr    error
	updateReq  *newsrpc.UpdateCommentStatusReq
	updateResp *newsrpc.UpdateCommentStatusResp
	updateErr  error
}

func (s *stubCommentNewsRPC) DeletesComment(_ context.Context, in *newsrpc.DeletesCommentReq, _ ...grpc.CallOption) (*newsrpc.DeletesCommentResp, error) {
	s.deleteReq = in
	return s.deleteResp, s.deleteErr
}

func (s *stubCommentNewsRPC) FindCommentList(_ context.Context, in *newsrpc.FindCommentListReq, _ ...grpc.CallOption) (*newsrpc.FindCommentListResp, error) {
	s.findReq = in
	return s.findResp, s.findErr
}

func (s *stubCommentNewsRPC) UpdateCommentStatus(_ context.Context, in *newsrpc.UpdateCommentStatusReq, _ ...grpc.CallOption) (*newsrpc.UpdateCommentStatusResp, error) {
	s.updateReq = in
	return s.updateResp, s.updateErr
}

type stubCommentAccountRPC struct {
	accountrpc.AccountRpc
	userReq     *accountrpc.FindUserListReq
	userResp    *accountrpc.FindUserListResp
	userErr     error
	visitorReq  *accountrpc.FindVisitorListReq
	visitorResp *accountrpc.FindVisitorListResp
	visitorErr  error
}

func (s *stubCommentAccountRPC) FindUserList(_ context.Context, in *accountrpc.FindUserListReq, _ ...grpc.CallOption) (*accountrpc.FindUserListResp, error) {
	s.userReq = in
	return s.userResp, s.userErr
}

func (s *stubCommentAccountRPC) FindVisitorList(_ context.Context, in *accountrpc.FindVisitorListReq, _ ...grpc.CallOption) (*accountrpc.FindVisitorListResp, error) {
	s.visitorReq = in
	return s.visitorResp, s.visitorErr
}

type stubCommentArticleRPC struct {
	articlerpc.ArticleRpc
	previewReq  *articlerpc.FindArticleListReq
	previewResp *articlerpc.FindArticlePreviewListResp
	previewErr  error
}

func (s *stubCommentArticleRPC) FindArticlePreviewList(_ context.Context, in *articlerpc.FindArticleListReq, _ ...grpc.CallOption) (*articlerpc.FindArticlePreviewListResp, error) {
	s.previewReq = in
	return s.previewResp, s.previewErr
}

func TestDeletesCommentBuildsRequest(t *testing.T) {
	rpc := &stubCommentNewsRPC{deleteResp: &newsrpc.DeletesCommentResp{SuccessCount: 2}}
	logic := NewDeletesCommentLogic(context.Background(), &svc.ServiceContext{NewsRpc: rpc})

	resp, err := logic.DeletesComment(&types.IdsReq{Ids: []int64{1, 2}})
	if err != nil {
		t.Fatalf("DeletesComment returned error: %v", err)
	}
	if rpc.deleteReq == nil || len(rpc.deleteReq.Ids) != 2 {
		t.Fatalf("unexpected delete request: %#v", rpc.deleteReq)
	}
	if resp.SuccessCount != 2 {
		t.Fatalf("unexpected success count: %d", resp.SuccessCount)
	}
}

func TestUpdateCommentStatusBuildsRequest(t *testing.T) {
	rpc := &stubCommentNewsRPC{updateResp: &newsrpc.UpdateCommentStatusResp{SuccessCount: 1}}
	logic := NewUpdateCommentStatusLogic(context.Background(), &svc.ServiceContext{NewsRpc: rpc})

	resp, err := logic.UpdateCommentStatus(&types.UpdateCommentStatusReq{
		Ids:    []int64{9},
		Status: 2,
	})
	if err != nil {
		t.Fatalf("UpdateCommentStatus returned error: %v", err)
	}
	if rpc.updateReq == nil || rpc.updateReq.Status != 2 || len(rpc.updateReq.Ids) != 1 {
		t.Fatalf("unexpected update request: %#v", rpc.updateReq)
	}
	if resp.SuccessCount != 1 {
		t.Fatalf("unexpected success count: %d", resp.SuccessCount)
	}
}

func TestFindCommentBackListBuildsQueryAndMapsPage(t *testing.T) {
	newsRPC := &stubCommentNewsRPC{
		findResp: &newsrpc.FindCommentListResp{
			Pagination: &newsrpc.PageResp{Page: 1, PageSize: 10, Total: 1},
			List: []*newsrpc.Comment{
				{
					Id:             1,
					UserId:         "u-1",
					TerminalId:     "t-1",
					TopicId:        101,
					ReplyUserId:    "u-2",
					CommentContent: "hello",
					Type:           1,
					Status:         2,
					CreatedAt:      123,
				},
			},
		},
	}
	accountRPC := &stubCommentAccountRPC{
		userResp: &accountrpc.FindUserListResp{
			List: []*accountrpc.User{
				{UserId: "u-1", Username: "user-1", Nickname: "User 1"},
				{UserId: "u-2", Username: "user-2", Nickname: "User 2"},
			},
		},
		visitorResp: &accountrpc.FindVisitorListResp{
			List: []*accountrpc.VisitorInfo{
				{TerminalId: "t-1", Os: "mac", Browser: "chrome"},
			},
		},
	}
	articleRPC := &stubCommentArticleRPC{
		previewResp: &articlerpc.FindArticlePreviewListResp{
			List: []*articlerpc.ArticlePreview{
				{Id: 101, ArticleTitle: "topic-title"},
			},
		},
	}
	logic := NewFindCommentBackListLogic(context.Background(), &svc.ServiceContext{
		NewsRpc:    newsRPC,
		AccountRpc: accountRPC,
		ArticleRpc: articleRPC,
	})

	resp, err := logic.FindCommentBackList(&types.QueryCommentReq{
		PageQuery: types.PageQuery{Page: 1, PageSize: 10, Sorts: []string{"created_at desc"}},
		UserId:    "u-1",
		Status:    2,
		Type:      1,
	})
	if err != nil {
		t.Fatalf("FindCommentBackList returned error: %v", err)
	}
	if newsRPC.findReq == nil || newsRPC.findReq.UserId != "u-1" || newsRPC.findReq.Status != 2 || newsRPC.findReq.Type != 1 {
		t.Fatalf("unexpected find request: %#v", newsRPC.findReq)
	}
	if accountRPC.userReq == nil || accountRPC.userReq.Paginate == nil || accountRPC.userReq.Paginate.PageSize != 2 {
		t.Fatalf("unexpected user request: %#v", accountRPC.userReq)
	}
	if accountRPC.visitorReq == nil || accountRPC.visitorReq.Paginate == nil || accountRPC.visitorReq.Paginate.PageSize != 1 {
		t.Fatalf("unexpected visitor request: %#v", accountRPC.visitorReq)
	}
	if articleRPC.previewReq == nil || len(articleRPC.previewReq.Ids) != 1 || articleRPC.previewReq.Ids[0] != 101 {
		t.Fatalf("unexpected article preview request: %#v", articleRPC.previewReq)
	}

	list, ok := resp.List.([]*types.CommentBackVO)
	if !ok || len(list) != 1 {
		t.Fatalf("unexpected list payload: %#v", resp.List)
	}
	if list[0].TopicTitle != "topic-title" || list[0].UserInfo == nil || list[0].ReplyUserInfo == nil || list[0].ClientInfo == nil {
		t.Fatalf("unexpected mapped comment: %#v", list[0])
	}
}
