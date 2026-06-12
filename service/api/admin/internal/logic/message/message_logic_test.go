package message

import (
	"context"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/newsrpc"
	"google.golang.org/grpc"
)

type stubMessageNewsRPC struct {
	deleteReq  *newsrpc.DeletesMessageReq
	deleteResp *newsrpc.DeletesMessageResp
	findReq    *newsrpc.FindMessageListReq
	findResp   *newsrpc.FindMessageListResp
	updateReq  *newsrpc.UpdateMessageStatusReq
	updateResp *newsrpc.UpdateMessageStatusResp
}

func (s *stubMessageNewsRPC) AnalysisMessage(context.Context, *newsrpc.AnalysisMessageReq, ...grpc.CallOption) (*newsrpc.AnalysisMessageResp, error) {
	panic("unexpected call")
}
func (s *stubMessageNewsRPC) AddMessage(context.Context, *newsrpc.AddMessageReq, ...grpc.CallOption) (*newsrpc.AddMessageResp, error) {
	panic("unexpected call")
}
func (s *stubMessageNewsRPC) UpdateMessage(context.Context, *newsrpc.UpdateMessageReq, ...grpc.CallOption) (*newsrpc.UpdateMessageResp, error) {
	panic("unexpected call")
}
func (s *stubMessageNewsRPC) GetMessage(context.Context, *newsrpc.GetMessageReq, ...grpc.CallOption) (*newsrpc.GetMessageResp, error) {
	panic("unexpected call")
}
func (s *stubMessageNewsRPC) DeletesMessage(_ context.Context, in *newsrpc.DeletesMessageReq, _ ...grpc.CallOption) (*newsrpc.DeletesMessageResp, error) {
	s.deleteReq = in
	return s.deleteResp, nil
}
func (s *stubMessageNewsRPC) FindMessageList(_ context.Context, in *newsrpc.FindMessageListReq, _ ...grpc.CallOption) (*newsrpc.FindMessageListResp, error) {
	s.findReq = in
	return s.findResp, nil
}
func (s *stubMessageNewsRPC) UpdateMessageStatus(_ context.Context, in *newsrpc.UpdateMessageStatusReq, _ ...grpc.CallOption) (*newsrpc.UpdateMessageStatusResp, error) {
	s.updateReq = in
	return s.updateResp, nil
}
func (s *stubMessageNewsRPC) AddComment(context.Context, *newsrpc.AddCommentReq, ...grpc.CallOption) (*newsrpc.AddCommentResp, error) {
	panic("unexpected call")
}
func (s *stubMessageNewsRPC) UpdateComment(context.Context, *newsrpc.UpdateCommentReq, ...grpc.CallOption) (*newsrpc.UpdateCommentResp, error) {
	panic("unexpected call")
}
func (s *stubMessageNewsRPC) GetComment(context.Context, *newsrpc.GetCommentReq, ...grpc.CallOption) (*newsrpc.GetCommentResp, error) {
	panic("unexpected call")
}
func (s *stubMessageNewsRPC) DeletesComment(context.Context, *newsrpc.DeletesCommentReq, ...grpc.CallOption) (*newsrpc.DeletesCommentResp, error) {
	panic("unexpected call")
}
func (s *stubMessageNewsRPC) FindCommentList(context.Context, *newsrpc.FindCommentListReq, ...grpc.CallOption) (*newsrpc.FindCommentListResp, error) {
	panic("unexpected call")
}
func (s *stubMessageNewsRPC) FindCommentReplyList(context.Context, *newsrpc.FindCommentReplyListReq, ...grpc.CallOption) (*newsrpc.FindCommentReplyListResp, error) {
	panic("unexpected call")
}
func (s *stubMessageNewsRPC) FindCommentReplyCounts(context.Context, *newsrpc.FindCommentReplyCountsReq, ...grpc.CallOption) (*newsrpc.FindCommentReplyCountsResp, error) {
	panic("unexpected call")
}
func (s *stubMessageNewsRPC) UpdateCommentStatus(context.Context, *newsrpc.UpdateCommentStatusReq, ...grpc.CallOption) (*newsrpc.UpdateCommentStatusResp, error) {
	panic("unexpected call")
}
func (s *stubMessageNewsRPC) LikeComment(context.Context, *newsrpc.LikeCommentReq, ...grpc.CallOption) (*newsrpc.LikeCommentResp, error) {
	panic("unexpected call")
}
func (s *stubMessageNewsRPC) FindUserLikeComment(context.Context, *newsrpc.FindUserLikeCommentReq, ...grpc.CallOption) (*newsrpc.FindLikeCommentResp, error) {
	panic("unexpected call")
}

type stubMessageAccountRPC struct {
	findUserResp    *accountrpc.FindUserListResp
	findVisitorResp *accountrpc.FindVisitorListResp
}

func (s *stubMessageAccountRPC) Login(context.Context, *accountrpc.LoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) Logout(context.Context, *accountrpc.LogoutReq, ...grpc.CallOption) (*accountrpc.LogoutResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) Logoff(context.Context, *accountrpc.LogoffReq, ...grpc.CallOption) (*accountrpc.LogoffResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) Register(context.Context, *accountrpc.RegisterReq, ...grpc.CallOption) (*accountrpc.RegisterResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) ResetPassword(context.Context, *accountrpc.ResetPasswordReq, ...grpc.CallOption) (*accountrpc.ResetPasswordResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) EmailLogin(context.Context, *accountrpc.EmailLoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) PhoneLogin(context.Context, *accountrpc.PhoneLoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) ThirdLogin(context.Context, *accountrpc.ThirdLoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) GetUserInfo(context.Context, *accountrpc.GetUserInfoReq, ...grpc.CallOption) (*accountrpc.GetUserInfoResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) GetUserOauthInfo(context.Context, *accountrpc.GetUserOauthInfoReq, ...grpc.CallOption) (*accountrpc.GetUserOauthInfoResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) UpdateUserInfo(context.Context, *accountrpc.UpdateUserInfoReq, ...grpc.CallOption) (*accountrpc.UpdateUserInfoResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) UpdateUserAvatar(context.Context, *accountrpc.UpdateUserAvatarReq, ...grpc.CallOption) (*accountrpc.UpdateUserAvatarResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) UpdateUserPassword(context.Context, *accountrpc.UpdateUserPasswordReq, ...grpc.CallOption) (*accountrpc.UpdateUserPasswordResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) BindUserEmail(context.Context, *accountrpc.BindUserEmailReq, ...grpc.CallOption) (*accountrpc.BindUserEmailResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) BindUserPhone(context.Context, *accountrpc.BindUserPhoneReq, ...grpc.CallOption) (*accountrpc.BindUserPhoneResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) BindUserOauth(context.Context, *accountrpc.BindUserOauthReq, ...grpc.CallOption) (*accountrpc.BindUserOauthResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) UnbindUserOauth(context.Context, *accountrpc.UnbindUserOauthReq, ...grpc.CallOption) (*accountrpc.UnbindUserOauthResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) AdminUpdateUserStatus(context.Context, *accountrpc.AdminUpdateUserStatusReq, ...grpc.CallOption) (*accountrpc.AdminUpdateUserStatusResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) AdminResetUserPassword(context.Context, *accountrpc.AdminResetUserPasswordReq, ...grpc.CallOption) (*accountrpc.AdminResetUserPasswordResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) FindUserList(context.Context, *accountrpc.FindUserListReq, ...grpc.CallOption) (*accountrpc.FindUserListResp, error) {
	return s.findUserResp, nil
}
func (s *stubMessageAccountRPC) FindUserInfoList(context.Context, *accountrpc.FindUserListReq, ...grpc.CallOption) (*accountrpc.FindUserInfoListResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) FindUserOnlineList(context.Context, *accountrpc.FindUserListReq, ...grpc.CallOption) (*accountrpc.FindUserInfoListResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) AnalysisUser(context.Context, *accountrpc.AnalysisUserReq, ...grpc.CallOption) (*accountrpc.AnalysisUserResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) AnalysisUserAreas(context.Context, *accountrpc.AnalysisUserAreasReq, ...grpc.CallOption) (*accountrpc.AnalysisUserAreasResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) GetClientInfo(context.Context, *accountrpc.GetClientInfoReq, ...grpc.CallOption) (*accountrpc.GetClientInfoResp, error) {
	panic("unexpected call")
}
func (s *stubMessageAccountRPC) FindVisitorList(context.Context, *accountrpc.FindVisitorListReq, ...grpc.CallOption) (*accountrpc.FindVisitorListResp, error) {
	return s.findVisitorResp, nil
}

func TestDeletesMessageBuildsRequest(t *testing.T) {
	rpc := &stubMessageNewsRPC{deleteResp: &newsrpc.DeletesMessageResp{SuccessCount: 2}}
	logic := NewDeletesMessageLogic(context.Background(), &svc.ServiceContext{NewsRpc: rpc})

	resp, err := logic.DeletesMessage(&types.IdsReq{Ids: []int64{1, 2}})
	if err != nil {
		t.Fatalf("DeletesMessage returned error: %v", err)
	}
	if rpc.deleteReq == nil || len(rpc.deleteReq.Ids) != 2 {
		t.Fatalf("unexpected delete request: %#v", rpc.deleteReq)
	}
	if resp.SuccessCount != 2 {
		t.Fatalf("unexpected success count: %d", resp.SuccessCount)
	}
}

func TestUpdateMessageStatusBuildsRequest(t *testing.T) {
	rpc := &stubMessageNewsRPC{updateResp: &newsrpc.UpdateMessageStatusResp{SuccessCount: 1}}
	logic := NewUpdateMessageStatusLogic(context.Background(), &svc.ServiceContext{NewsRpc: rpc})

	resp, err := logic.UpdateMessageStatus(&types.UpdateMessageStatusReq{
		Ids:    []int64{9},
		Status: 2,
	})
	if err != nil {
		t.Fatalf("UpdateMessageStatus returned error: %v", err)
	}
	if rpc.updateReq == nil || rpc.updateReq.Status != 2 || len(rpc.updateReq.Ids) != 1 {
		t.Fatalf("unexpected update request: %#v", rpc.updateReq)
	}
	if resp.SuccessCount != 1 {
		t.Fatalf("unexpected success count: %d", resp.SuccessCount)
	}
}

func TestFindMessageListBuildsQueryAndMapsPage(t *testing.T) {
	newsRPC := &stubMessageNewsRPC{
		findResp: &newsrpc.FindMessageListResp{
			Pagination: &newsrpc.PageResp{Page: 1, PageSize: 10, Total: 1},
			List: []*newsrpc.Message{
				{Id: 1, UserId: "u-1", TerminalId: "t-1", MessageContent: "hello", Status: 1},
			},
		},
	}
	accountRPC := &stubMessageAccountRPC{
		findUserResp:    &accountrpc.FindUserListResp{List: []*accountrpc.User{{UserId: "u-1", Username: "user"}}},
		findVisitorResp: &accountrpc.FindVisitorListResp{List: []*accountrpc.VisitorInfo{{TerminalId: "t-1", Os: "mac"}}},
	}
	logic := NewFindMessageListLogic(context.Background(), &svc.ServiceContext{
		NewsRpc:    newsRPC,
		AccountRpc: accountRPC,
	})

	resp, err := logic.FindMessageList(&types.QueryMessageReq{
		PageQuery: types.PageQuery{Page: 1, PageSize: 10},
		UserId:    "u-1",
		Status:    1,
	})
	if err != nil {
		t.Fatalf("FindMessageList returned error: %v", err)
	}
	if newsRPC.findReq == nil || newsRPC.findReq.UserId != "u-1" || newsRPC.findReq.Status != 1 {
		t.Fatalf("unexpected find request: %#v", newsRPC.findReq)
	}
	list, ok := resp.List.([]*types.MessageBackVO)
	if !ok || len(list) != 1 || list[0].UserInfo == nil || list[0].ClientInfo == nil {
		t.Fatalf("unexpected list payload: %#v", resp.List)
	}
}
