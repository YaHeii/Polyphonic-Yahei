package visitor

import (
	"context"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
	"google.golang.org/grpc"
)

type stubVisitorAccountRPC struct {
	findReq  *accountrpc.FindVisitorListReq
	findResp *accountrpc.FindVisitorListResp
}

func (s *stubVisitorAccountRPC) Login(context.Context, *accountrpc.LoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) Logout(context.Context, *accountrpc.LogoutReq, ...grpc.CallOption) (*accountrpc.LogoutResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) Logoff(context.Context, *accountrpc.LogoffReq, ...grpc.CallOption) (*accountrpc.LogoffResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) Register(context.Context, *accountrpc.RegisterReq, ...grpc.CallOption) (*accountrpc.RegisterResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) ResetPassword(context.Context, *accountrpc.ResetPasswordReq, ...grpc.CallOption) (*accountrpc.ResetPasswordResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) EmailLogin(context.Context, *accountrpc.EmailLoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) PhoneLogin(context.Context, *accountrpc.PhoneLoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) ThirdLogin(context.Context, *accountrpc.ThirdLoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) GetUserInfo(context.Context, *accountrpc.GetUserInfoReq, ...grpc.CallOption) (*accountrpc.GetUserInfoResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) GetUserOauthInfo(context.Context, *accountrpc.GetUserOauthInfoReq, ...grpc.CallOption) (*accountrpc.GetUserOauthInfoResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) UpdateUserInfo(context.Context, *accountrpc.UpdateUserInfoReq, ...grpc.CallOption) (*accountrpc.UpdateUserInfoResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) UpdateUserAvatar(context.Context, *accountrpc.UpdateUserAvatarReq, ...grpc.CallOption) (*accountrpc.UpdateUserAvatarResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) UpdateUserPassword(context.Context, *accountrpc.UpdateUserPasswordReq, ...grpc.CallOption) (*accountrpc.UpdateUserPasswordResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) BindUserEmail(context.Context, *accountrpc.BindUserEmailReq, ...grpc.CallOption) (*accountrpc.BindUserEmailResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) BindUserPhone(context.Context, *accountrpc.BindUserPhoneReq, ...grpc.CallOption) (*accountrpc.BindUserPhoneResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) BindUserOauth(context.Context, *accountrpc.BindUserOauthReq, ...grpc.CallOption) (*accountrpc.BindUserOauthResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) UnbindUserOauth(context.Context, *accountrpc.UnbindUserOauthReq, ...grpc.CallOption) (*accountrpc.UnbindUserOauthResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) AdminUpdateUserStatus(context.Context, *accountrpc.AdminUpdateUserStatusReq, ...grpc.CallOption) (*accountrpc.AdminUpdateUserStatusResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) AdminResetUserPassword(context.Context, *accountrpc.AdminResetUserPasswordReq, ...grpc.CallOption) (*accountrpc.AdminResetUserPasswordResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) FindUserList(context.Context, *accountrpc.FindUserListReq, ...grpc.CallOption) (*accountrpc.FindUserListResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) FindUserInfoList(context.Context, *accountrpc.FindUserListReq, ...grpc.CallOption) (*accountrpc.FindUserInfoListResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) FindUserOnlineList(context.Context, *accountrpc.FindUserListReq, ...grpc.CallOption) (*accountrpc.FindUserInfoListResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) AnalysisUser(context.Context, *accountrpc.AnalysisUserReq, ...grpc.CallOption) (*accountrpc.AnalysisUserResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) AnalysisUserAreas(context.Context, *accountrpc.AnalysisUserAreasReq, ...grpc.CallOption) (*accountrpc.AnalysisUserAreasResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) GetClientInfo(context.Context, *accountrpc.GetClientInfoReq, ...grpc.CallOption) (*accountrpc.GetClientInfoResp, error) {
	panic("unexpected call")
}
func (s *stubVisitorAccountRPC) FindVisitorList(_ context.Context, in *accountrpc.FindVisitorListReq, _ ...grpc.CallOption) (*accountrpc.FindVisitorListResp, error) {
	s.findReq = in
	return s.findResp, nil
}

func TestFindVisitorListBuildsQueryAndMapsPage(t *testing.T) {
	rpc := &stubVisitorAccountRPC{
		findResp: &accountrpc.FindVisitorListResp{
			Pagination: &accountrpc.PageResp{Page: 2, PageSize: 5, Total: 1},
			List: []*accountrpc.VisitorInfo{
				{Id: 1, TerminalId: "t-1", Os: "mac", IpSource: "shanghai"},
			},
		},
	}
	logic := NewFindVisitorListLogic(context.Background(), &svc.ServiceContext{AccountRpc: rpc})

	resp, err := logic.FindVisitorList(&types.QueryVisitorReq{
		PageQuery:  types.PageQuery{Page: 2, PageSize: 5},
		TerminalId: "t-1",
		IpSource:   "shanghai",
	})
	if err != nil {
		t.Fatalf("FindVisitorList returned error: %v", err)
	}
	if rpc.findReq == nil || rpc.findReq.TerminalId != "t-1" || rpc.findReq.IpSource != "shanghai" {
		t.Fatalf("unexpected request: %#v", rpc.findReq)
	}
	list, ok := resp.List.([]*types.VisitorBackVO)
	if !ok || len(list) != 1 || list[0].TerminalId != "t-1" {
		t.Fatalf("unexpected list payload: %#v", resp.List)
	}
}
