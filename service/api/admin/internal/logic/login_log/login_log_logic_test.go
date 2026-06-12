package login_log

import (
	"context"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/syslogrpc"
	"google.golang.org/grpc"
)

type stubLoginSyslogRPC struct {
	deleteReq  *syslogrpc.DeletesLoginLogReq
	deleteResp *syslogrpc.DeletesLoginLogResp
	findReq    *syslogrpc.FindLoginLogListReq
	findResp   *syslogrpc.FindLoginLogListResp
}

func (s *stubLoginSyslogRPC) AddLoginLog(context.Context, *syslogrpc.AddLoginLogReq, ...grpc.CallOption) (*syslogrpc.AddLoginLogResp, error) {
	panic("unexpected call")
}
func (s *stubLoginSyslogRPC) AddLogoutLog(context.Context, *syslogrpc.AddLogoutLogReq, ...grpc.CallOption) (*syslogrpc.AddLogoutLogResp, error) {
	panic("unexpected call")
}
func (s *stubLoginSyslogRPC) DeletesLoginLog(_ context.Context, in *syslogrpc.DeletesLoginLogReq, _ ...grpc.CallOption) (*syslogrpc.DeletesLoginLogResp, error) {
	s.deleteReq = in
	return s.deleteResp, nil
}
func (s *stubLoginSyslogRPC) FindLoginLogList(_ context.Context, in *syslogrpc.FindLoginLogListReq, _ ...grpc.CallOption) (*syslogrpc.FindLoginLogListResp, error) {
	s.findReq = in
	return s.findResp, nil
}
func (s *stubLoginSyslogRPC) AddVisitLog(context.Context, *syslogrpc.AddVisitLogReq, ...grpc.CallOption) (*syslogrpc.AddVisitLogResp, error) {
	panic("unexpected call")
}
func (s *stubLoginSyslogRPC) DeletesVisitLog(context.Context, *syslogrpc.DeletesVisitLogReq, ...grpc.CallOption) (*syslogrpc.DeletesVisitLogResp, error) {
	panic("unexpected call")
}
func (s *stubLoginSyslogRPC) FindVisitLogList(context.Context, *syslogrpc.FindVisitLogListReq, ...grpc.CallOption) (*syslogrpc.FindVisitLogListResp, error) {
	panic("unexpected call")
}
func (s *stubLoginSyslogRPC) AddOperationLog(context.Context, *syslogrpc.AddOperationLogReq, ...grpc.CallOption) (*syslogrpc.AddOperationLogResp, error) {
	panic("unexpected call")
}
func (s *stubLoginSyslogRPC) DeletesOperationLog(context.Context, *syslogrpc.DeletesOperationLogReq, ...grpc.CallOption) (*syslogrpc.DeletesOperationLogResp, error) {
	panic("unexpected call")
}
func (s *stubLoginSyslogRPC) FindOperationLogList(context.Context, *syslogrpc.FindOperationLogListReq, ...grpc.CallOption) (*syslogrpc.FindOperationLogListResp, error) {
	panic("unexpected call")
}
func (s *stubLoginSyslogRPC) AddFileLog(context.Context, *syslogrpc.AddFileLogReq, ...grpc.CallOption) (*syslogrpc.AddFileLogResp, error) {
	panic("unexpected call")
}
func (s *stubLoginSyslogRPC) DeletesFileLog(context.Context, *syslogrpc.DeletesFileLogReq, ...grpc.CallOption) (*syslogrpc.DeletesFileLogResp, error) {
	panic("unexpected call")
}
func (s *stubLoginSyslogRPC) FindFileLogList(context.Context, *syslogrpc.FindFileLogListReq, ...grpc.CallOption) (*syslogrpc.FindFileLogListResp, error) {
	panic("unexpected call")
}

type stubLoginAccountRPC struct {
	findUserReq     *accountrpc.FindUserListReq
	findUserResp    *accountrpc.FindUserListResp
	findVisitorReq  *accountrpc.FindVisitorListReq
	findVisitorResp *accountrpc.FindVisitorListResp
}

func (s *stubLoginAccountRPC) Login(context.Context, *accountrpc.LoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) Logout(context.Context, *accountrpc.LogoutReq, ...grpc.CallOption) (*accountrpc.LogoutResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) Logoff(context.Context, *accountrpc.LogoffReq, ...grpc.CallOption) (*accountrpc.LogoffResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) Register(context.Context, *accountrpc.RegisterReq, ...grpc.CallOption) (*accountrpc.RegisterResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) ResetPassword(context.Context, *accountrpc.ResetPasswordReq, ...grpc.CallOption) (*accountrpc.ResetPasswordResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) EmailLogin(context.Context, *accountrpc.EmailLoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) PhoneLogin(context.Context, *accountrpc.PhoneLoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) ThirdLogin(context.Context, *accountrpc.ThirdLoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) GetUserInfo(context.Context, *accountrpc.GetUserInfoReq, ...grpc.CallOption) (*accountrpc.GetUserInfoResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) GetUserOauthInfo(context.Context, *accountrpc.GetUserOauthInfoReq, ...grpc.CallOption) (*accountrpc.GetUserOauthInfoResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) UpdateUserInfo(context.Context, *accountrpc.UpdateUserInfoReq, ...grpc.CallOption) (*accountrpc.UpdateUserInfoResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) UpdateUserAvatar(context.Context, *accountrpc.UpdateUserAvatarReq, ...grpc.CallOption) (*accountrpc.UpdateUserAvatarResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) UpdateUserPassword(context.Context, *accountrpc.UpdateUserPasswordReq, ...grpc.CallOption) (*accountrpc.UpdateUserPasswordResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) BindUserEmail(context.Context, *accountrpc.BindUserEmailReq, ...grpc.CallOption) (*accountrpc.BindUserEmailResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) BindUserPhone(context.Context, *accountrpc.BindUserPhoneReq, ...grpc.CallOption) (*accountrpc.BindUserPhoneResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) BindUserOauth(context.Context, *accountrpc.BindUserOauthReq, ...grpc.CallOption) (*accountrpc.BindUserOauthResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) UnbindUserOauth(context.Context, *accountrpc.UnbindUserOauthReq, ...grpc.CallOption) (*accountrpc.UnbindUserOauthResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) AdminUpdateUserStatus(context.Context, *accountrpc.AdminUpdateUserStatusReq, ...grpc.CallOption) (*accountrpc.AdminUpdateUserStatusResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) AdminResetUserPassword(context.Context, *accountrpc.AdminResetUserPasswordReq, ...grpc.CallOption) (*accountrpc.AdminResetUserPasswordResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) FindUserList(_ context.Context, in *accountrpc.FindUserListReq, _ ...grpc.CallOption) (*accountrpc.FindUserListResp, error) {
	s.findUserReq = in
	return s.findUserResp, nil
}
func (s *stubLoginAccountRPC) FindUserInfoList(context.Context, *accountrpc.FindUserListReq, ...grpc.CallOption) (*accountrpc.FindUserInfoListResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) FindUserOnlineList(context.Context, *accountrpc.FindUserListReq, ...grpc.CallOption) (*accountrpc.FindUserInfoListResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) AnalysisUser(context.Context, *accountrpc.AnalysisUserReq, ...grpc.CallOption) (*accountrpc.AnalysisUserResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) AnalysisUserAreas(context.Context, *accountrpc.AnalysisUserAreasReq, ...grpc.CallOption) (*accountrpc.AnalysisUserAreasResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) GetClientInfo(context.Context, *accountrpc.GetClientInfoReq, ...grpc.CallOption) (*accountrpc.GetClientInfoResp, error) {
	panic("unexpected call")
}
func (s *stubLoginAccountRPC) FindVisitorList(_ context.Context, in *accountrpc.FindVisitorListReq, _ ...grpc.CallOption) (*accountrpc.FindVisitorListResp, error) {
	s.findVisitorReq = in
	return s.findVisitorResp, nil
}

func TestDeletesLoginLogBuildsRequest(t *testing.T) {
	rpc := &stubLoginSyslogRPC{deleteResp: &syslogrpc.DeletesLoginLogResp{SuccessCount: 2}}
	logic := NewDeletesLoginLogLogic(context.Background(), &svc.ServiceContext{SyslogRpc: rpc})

	resp, err := logic.DeletesLoginLog(&types.IdsReq{Ids: []int64{1, 2}})
	if err != nil {
		t.Fatalf("DeletesLoginLog returned error: %v", err)
	}
	if rpc.deleteReq == nil || len(rpc.deleteReq.Ids) != 2 {
		t.Fatalf("unexpected delete request: %#v", rpc.deleteReq)
	}
	if resp.SuccessCount != 2 {
		t.Fatalf("unexpected success count: %d", resp.SuccessCount)
	}
}

func TestFindLoginLogListBuildsQueryAndMapsPage(t *testing.T) {
	syslogRPC := &stubLoginSyslogRPC{
		findResp: &syslogrpc.FindLoginLogListResp{
			Pagination: &syslogrpc.PageResp{Page: 1, PageSize: 10, Total: 1},
			List: []*syslogrpc.LoginLog{
				{Id: 1, UserId: "u-1", TerminalId: "t-1", LoginType: "password"},
			},
		},
	}
	accountRPC := &stubLoginAccountRPC{
		findUserResp: &accountrpc.FindUserListResp{
			List: []*accountrpc.User{{UserId: "u-1", Username: "user"}},
		},
		findVisitorResp: &accountrpc.FindVisitorListResp{
			List: []*accountrpc.VisitorInfo{{TerminalId: "t-1", Os: "mac"}},
		},
	}
	logic := NewFindLoginLogListLogic(context.Background(), &svc.ServiceContext{
		SyslogRpc:  syslogRPC,
		AccountRpc: accountRPC,
	})

	resp, err := logic.FindLoginLogList(&types.QueryLoginLogReq{
		PageQuery: types.PageQuery{Page: 1, PageSize: 10},
		UserId:    "u-1",
	})
	if err != nil {
		t.Fatalf("FindLoginLogList returned error: %v", err)
	}
	if syslogRPC.findReq == nil || syslogRPC.findReq.UserId != "u-1" {
		t.Fatalf("unexpected syslog request: %#v", syslogRPC.findReq)
	}
	list, ok := resp.List.([]*types.LoginLogBackVO)
	if !ok || len(list) != 1 || list[0].UserInfo == nil || list[0].ClientInfo == nil {
		t.Fatalf("unexpected response list: %#v", resp.List)
	}
}
