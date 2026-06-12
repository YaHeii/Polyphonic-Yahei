package talk

import (
	"context"
	"errors"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizheader"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/socialrpc"
	"google.golang.org/grpc"
)

type stubTalkSocialRPC struct {
	addTalkReq     *socialrpc.AddTalkReq
	addTalkResp    *socialrpc.AddTalkResp
	addTalkErr     error
	updateTalkReq  *socialrpc.UpdateTalkReq
	updateTalkResp *socialrpc.UpdateTalkResp
	updateTalkErr  error
	deleteTalkReq  *socialrpc.DeletesTalkReq
	deleteTalkResp *socialrpc.DeletesTalkResp
	deleteTalkErr  error
	getTalkReq     *socialrpc.GetTalkReq
	getTalkResp    *socialrpc.GetTalkResp
	getTalkErr     error
	findTalkReq    *socialrpc.FindTalkListReq
	findTalkResp   *socialrpc.FindTalkListResp
	findTalkErr    error
}

type stubTalkAccountRPC struct {
	findUserListReq  *accountrpc.FindUserListReq
	findUserListResp *accountrpc.FindUserListResp
	findUserListErr  error
}

func (s *stubTalkAccountRPC) Login(context.Context, *accountrpc.LoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) Logout(context.Context, *accountrpc.LogoutReq, ...grpc.CallOption) (*accountrpc.LogoutResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) Logoff(context.Context, *accountrpc.LogoffReq, ...grpc.CallOption) (*accountrpc.LogoffResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) Register(context.Context, *accountrpc.RegisterReq, ...grpc.CallOption) (*accountrpc.RegisterResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) ResetPassword(context.Context, *accountrpc.ResetPasswordReq, ...grpc.CallOption) (*accountrpc.ResetPasswordResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) EmailLogin(context.Context, *accountrpc.EmailLoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) PhoneLogin(context.Context, *accountrpc.PhoneLoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) ThirdLogin(context.Context, *accountrpc.ThirdLoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) GetUserInfo(context.Context, *accountrpc.GetUserInfoReq, ...grpc.CallOption) (*accountrpc.GetUserInfoResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) GetUserOauthInfo(context.Context, *accountrpc.GetUserOauthInfoReq, ...grpc.CallOption) (*accountrpc.GetUserOauthInfoResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) UpdateUserInfo(context.Context, *accountrpc.UpdateUserInfoReq, ...grpc.CallOption) (*accountrpc.UpdateUserInfoResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) UpdateUserAvatar(context.Context, *accountrpc.UpdateUserAvatarReq, ...grpc.CallOption) (*accountrpc.UpdateUserAvatarResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) UpdateUserPassword(context.Context, *accountrpc.UpdateUserPasswordReq, ...grpc.CallOption) (*accountrpc.UpdateUserPasswordResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) BindUserEmail(context.Context, *accountrpc.BindUserEmailReq, ...grpc.CallOption) (*accountrpc.BindUserEmailResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) BindUserPhone(context.Context, *accountrpc.BindUserPhoneReq, ...grpc.CallOption) (*accountrpc.BindUserPhoneResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) BindUserOauth(context.Context, *accountrpc.BindUserOauthReq, ...grpc.CallOption) (*accountrpc.BindUserOauthResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) UnbindUserOauth(context.Context, *accountrpc.UnbindUserOauthReq, ...grpc.CallOption) (*accountrpc.UnbindUserOauthResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) AdminUpdateUserStatus(context.Context, *accountrpc.AdminUpdateUserStatusReq, ...grpc.CallOption) (*accountrpc.AdminUpdateUserStatusResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) AdminResetUserPassword(context.Context, *accountrpc.AdminResetUserPasswordReq, ...grpc.CallOption) (*accountrpc.AdminResetUserPasswordResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) FindUserList(_ context.Context, in *accountrpc.FindUserListReq, _ ...grpc.CallOption) (*accountrpc.FindUserListResp, error) {
	s.findUserListReq = in
	return s.findUserListResp, s.findUserListErr
}

func (s *stubTalkAccountRPC) FindUserInfoList(context.Context, *accountrpc.FindUserListReq, ...grpc.CallOption) (*accountrpc.FindUserInfoListResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) FindUserOnlineList(context.Context, *accountrpc.FindUserListReq, ...grpc.CallOption) (*accountrpc.FindUserInfoListResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) AnalysisUser(context.Context, *accountrpc.AnalysisUserReq, ...grpc.CallOption) (*accountrpc.AnalysisUserResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) AnalysisUserAreas(context.Context, *accountrpc.AnalysisUserAreasReq, ...grpc.CallOption) (*accountrpc.AnalysisUserAreasResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) GetClientInfo(context.Context, *accountrpc.GetClientInfoReq, ...grpc.CallOption) (*accountrpc.GetClientInfoResp, error) {
	panic("unexpected call")
}

func (s *stubTalkAccountRPC) FindVisitorList(context.Context, *accountrpc.FindVisitorListReq, ...grpc.CallOption) (*accountrpc.FindVisitorListResp, error) {
	panic("unexpected call")
}

func (s *stubTalkSocialRPC) AddFriend(context.Context, *socialrpc.AddFriendReq, ...grpc.CallOption) (*socialrpc.AddFriendResp, error) {
	panic("unexpected call")
}

func (s *stubTalkSocialRPC) UpdateFriend(context.Context, *socialrpc.UpdateFriendReq, ...grpc.CallOption) (*socialrpc.UpdateFriendResp, error) {
	panic("unexpected call")
}

func (s *stubTalkSocialRPC) DeletesFriend(context.Context, *socialrpc.DeletesFriendReq, ...grpc.CallOption) (*socialrpc.DeletesFriendResp, error) {
	panic("unexpected call")
}

func (s *stubTalkSocialRPC) FindFriendList(context.Context, *socialrpc.FindFriendListReq, ...grpc.CallOption) (*socialrpc.FindFriendListResp, error) {
	panic("unexpected call")
}

func (s *stubTalkSocialRPC) AddTalk(_ context.Context, in *socialrpc.AddTalkReq, _ ...grpc.CallOption) (*socialrpc.AddTalkResp, error) {
	s.addTalkReq = in
	return s.addTalkResp, s.addTalkErr
}

func (s *stubTalkSocialRPC) UpdateTalk(_ context.Context, in *socialrpc.UpdateTalkReq, _ ...grpc.CallOption) (*socialrpc.UpdateTalkResp, error) {
	s.updateTalkReq = in
	return s.updateTalkResp, s.updateTalkErr
}

func (s *stubTalkSocialRPC) DeletesTalk(_ context.Context, in *socialrpc.DeletesTalkReq, _ ...grpc.CallOption) (*socialrpc.DeletesTalkResp, error) {
	s.deleteTalkReq = in
	return s.deleteTalkResp, s.deleteTalkErr
}

func (s *stubTalkSocialRPC) GetTalk(_ context.Context, in *socialrpc.GetTalkReq, _ ...grpc.CallOption) (*socialrpc.GetTalkResp, error) {
	s.getTalkReq = in
	return s.getTalkResp, s.getTalkErr
}

func (s *stubTalkSocialRPC) FindTalkList(_ context.Context, in *socialrpc.FindTalkListReq, _ ...grpc.CallOption) (*socialrpc.FindTalkListResp, error) {
	s.findTalkReq = in
	return s.findTalkResp, s.findTalkErr
}

func (s *stubTalkSocialRPC) LikeTalk(context.Context, *socialrpc.LikeTalkReq, ...grpc.CallOption) (*socialrpc.LikeTalkResp, error) {
	panic("unexpected call")
}

func (s *stubTalkSocialRPC) FindUserLikeTalk(context.Context, *socialrpc.FindUserLikeTalkReq, ...grpc.CallOption) (*socialrpc.FindUserLikeTalkResp, error) {
	panic("unexpected call")
}

func TestAddTalkBuildsRequest(t *testing.T) {
	rpc := &stubTalkSocialRPC{
		addTalkResp: &socialrpc.AddTalkResp{
			Talk: &socialrpc.Talk{Id: 1, UserId: "u-1", Content: "hello"},
		},
	}
	ctx := context.WithValue(context.Background(), bizheader.HeaderUid, "u-1")
	logic := NewAddTalkLogic(ctx, &svc.ServiceContext{SocialRpc: rpc})

	resp, err := logic.AddTalk(&types.NewTalkReq{
		Id:      1,
		Content: "hello",
		ImgList: []string{"a"},
		IsTop:   true,
		Status:  2,
	})
	if err != nil {
		t.Fatalf("AddTalk returned error: %v", err)
	}

	if rpc.addTalkReq == nil || rpc.addTalkReq.UserId != "u-1" || rpc.addTalkReq.Content != "hello" {
		t.Fatalf("unexpected add request: %#v", rpc.addTalkReq)
	}

	if resp.Id != 1 || resp.UserId != "u-1" {
		t.Fatalf("unexpected response: %#v", resp)
	}
}

func TestDeleteTalkPropagatesRPCError(t *testing.T) {
	wantErr := errors.New("delete failed")
	logic := NewDeleteTalkLogic(context.Background(), &svc.ServiceContext{
		SocialRpc: &stubTalkSocialRPC{deleteTalkErr: wantErr},
	})

	_, err := logic.DeleteTalk(&types.IdReq{Id: 3})
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected %v, got %v", wantErr, err)
	}
}

func TestGetTalkBuildsRequest(t *testing.T) {
	rpc := &stubTalkSocialRPC{
		getTalkResp: &socialrpc.GetTalkResp{
			Talk: &socialrpc.Talk{Id: 9, Content: "detail"},
		},
	}
	logic := NewGetTalkLogic(context.Background(), &svc.ServiceContext{SocialRpc: rpc})

	resp, err := logic.GetTalk(&types.IdReq{Id: 9})
	if err != nil {
		t.Fatalf("GetTalk returned error: %v", err)
	}

	if rpc.getTalkReq == nil || rpc.getTalkReq.Id != 9 {
		t.Fatalf("unexpected get request: %#v", rpc.getTalkReq)
	}

	if resp.Id != 9 || resp.Content != "detail" {
		t.Fatalf("unexpected response: %#v", resp)
	}
}

func TestFindTalkListBuildsQuery(t *testing.T) {
	rpc := &stubTalkSocialRPC{
		findTalkResp: &socialrpc.FindTalkListResp{
			Pagination: &socialrpc.PageResp{Page: 2, PageSize: 5, Total: 1},
			List: []*socialrpc.Talk{
				{Id: 1, UserId: "u-1", Content: "hello"},
			},
		},
	}
	accountRPC := &stubTalkAccountRPC{
		findUserListResp: &accountrpc.FindUserListResp{
			List: []*accountrpc.User{{UserId: "u-1", Username: "user"}},
		},
	}
	logic := NewFindTalkListLogic(context.Background(), &svc.ServiceContext{
		SocialRpc:  rpc,
		AccountRpc: accountRPC,
	})

	resp, err := logic.FindTalkList(&types.QueryTalkReq{
		PageQuery: types.PageQuery{
			Page:     2,
			PageSize: 5,
			Sorts:    []string{"created_at desc"},
		},
		Status: 2,
	})
	if err != nil {
		t.Fatalf("FindTalkList returned error: %v", err)
	}

	if rpc.findTalkReq == nil || rpc.findTalkReq.Paginate == nil || rpc.findTalkReq.Status != 2 {
		t.Fatalf("unexpected find request: %#v", rpc.findTalkReq)
	}

	if accountRPC.findUserListReq == nil || len(accountRPC.findUserListReq.UserIds) != 1 || accountRPC.findUserListReq.UserIds[0] != "u-1" {
		t.Fatalf("unexpected account request: %#v", accountRPC.findUserListReq)
	}

	list, ok := resp.List.([]*types.TalkBackVO)
	if !ok || len(list) != 1 || list[0].Id != 1 || list[0].UserInfo == nil || list[0].UserInfo.UserId != "u-1" {
		t.Fatalf("unexpected list payload: %#v", resp.List)
	}
}
