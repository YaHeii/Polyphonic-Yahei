package visit_log

import (
	"context"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/syslogrpc"
	"google.golang.org/grpc"
)

type stubVisitSyslogRPC struct {
	deleteReq  *syslogrpc.DeletesVisitLogReq
	deleteResp *syslogrpc.DeletesVisitLogResp
	findReq    *syslogrpc.FindVisitLogListReq
	findResp   *syslogrpc.FindVisitLogListResp
}

func (s *stubVisitSyslogRPC) AddLoginLog(context.Context, *syslogrpc.AddLoginLogReq, ...grpc.CallOption) (*syslogrpc.AddLoginLogResp, error) {
	panic("unexpected call")
}
func (s *stubVisitSyslogRPC) AddLogoutLog(context.Context, *syslogrpc.AddLogoutLogReq, ...grpc.CallOption) (*syslogrpc.AddLogoutLogResp, error) {
	panic("unexpected call")
}
func (s *stubVisitSyslogRPC) DeletesLoginLog(context.Context, *syslogrpc.DeletesLoginLogReq, ...grpc.CallOption) (*syslogrpc.DeletesLoginLogResp, error) {
	panic("unexpected call")
}
func (s *stubVisitSyslogRPC) FindLoginLogList(context.Context, *syslogrpc.FindLoginLogListReq, ...grpc.CallOption) (*syslogrpc.FindLoginLogListResp, error) {
	panic("unexpected call")
}
func (s *stubVisitSyslogRPC) AddVisitLog(context.Context, *syslogrpc.AddVisitLogReq, ...grpc.CallOption) (*syslogrpc.AddVisitLogResp, error) {
	panic("unexpected call")
}
func (s *stubVisitSyslogRPC) DeletesVisitLog(_ context.Context, in *syslogrpc.DeletesVisitLogReq, _ ...grpc.CallOption) (*syslogrpc.DeletesVisitLogResp, error) {
	s.deleteReq = in
	return s.deleteResp, nil
}
func (s *stubVisitSyslogRPC) FindVisitLogList(_ context.Context, in *syslogrpc.FindVisitLogListReq, _ ...grpc.CallOption) (*syslogrpc.FindVisitLogListResp, error) {
	s.findReq = in
	return s.findResp, nil
}
func (s *stubVisitSyslogRPC) AddOperationLog(context.Context, *syslogrpc.AddOperationLogReq, ...grpc.CallOption) (*syslogrpc.AddOperationLogResp, error) {
	panic("unexpected call")
}
func (s *stubVisitSyslogRPC) DeletesOperationLog(context.Context, *syslogrpc.DeletesOperationLogReq, ...grpc.CallOption) (*syslogrpc.DeletesOperationLogResp, error) {
	panic("unexpected call")
}
func (s *stubVisitSyslogRPC) FindOperationLogList(context.Context, *syslogrpc.FindOperationLogListReq, ...grpc.CallOption) (*syslogrpc.FindOperationLogListResp, error) {
	panic("unexpected call")
}
func (s *stubVisitSyslogRPC) AddFileLog(context.Context, *syslogrpc.AddFileLogReq, ...grpc.CallOption) (*syslogrpc.AddFileLogResp, error) {
	panic("unexpected call")
}
func (s *stubVisitSyslogRPC) DeletesFileLog(context.Context, *syslogrpc.DeletesFileLogReq, ...grpc.CallOption) (*syslogrpc.DeletesFileLogResp, error) {
	panic("unexpected call")
}
func (s *stubVisitSyslogRPC) FindFileLogList(context.Context, *syslogrpc.FindFileLogListReq, ...grpc.CallOption) (*syslogrpc.FindFileLogListResp, error) {
	panic("unexpected call")
}

type stubVisitAccountRPC struct {
	findUserResp    *accountrpc.FindUserListResp
	findVisitorResp *accountrpc.FindVisitorListResp
}

func (s *stubVisitAccountRPC) Login(context.Context, *accountrpc.LoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) Logout(context.Context, *accountrpc.LogoutReq, ...grpc.CallOption) (*accountrpc.LogoutResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) Logoff(context.Context, *accountrpc.LogoffReq, ...grpc.CallOption) (*accountrpc.LogoffResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) Register(context.Context, *accountrpc.RegisterReq, ...grpc.CallOption) (*accountrpc.RegisterResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) ResetPassword(context.Context, *accountrpc.ResetPasswordReq, ...grpc.CallOption) (*accountrpc.ResetPasswordResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) EmailLogin(context.Context, *accountrpc.EmailLoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) PhoneLogin(context.Context, *accountrpc.PhoneLoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) ThirdLogin(context.Context, *accountrpc.ThirdLoginReq, ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) GetOauthAuthorizeUrl(context.Context, *accountrpc.GetOauthAuthorizeUrlReq, ...grpc.CallOption) (*accountrpc.GetOauthAuthorizeUrlResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) SendEmailVerifyCode(context.Context, *accountrpc.SendEmailVerifyCodeReq, ...grpc.CallOption) (*accountrpc.SendEmailVerifyCodeResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) SendPhoneVerifyCode(context.Context, *accountrpc.SendPhoneVerifyCodeReq, ...grpc.CallOption) (*accountrpc.SendPhoneVerifyCodeResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) GenerateCaptchaCode(context.Context, *accountrpc.GenerateCaptchaCodeReq, ...grpc.CallOption) (*accountrpc.GenerateCaptchaCodeResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) GetUserInfo(context.Context, *accountrpc.GetUserInfoReq, ...grpc.CallOption) (*accountrpc.GetUserInfoResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) GetUserOauthInfo(context.Context, *accountrpc.GetUserOauthInfoReq, ...grpc.CallOption) (*accountrpc.GetUserOauthInfoResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) UpdateUserInfo(context.Context, *accountrpc.UpdateUserInfoReq, ...grpc.CallOption) (*accountrpc.UpdateUserInfoResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) UpdateUserAvatar(context.Context, *accountrpc.UpdateUserAvatarReq, ...grpc.CallOption) (*accountrpc.UpdateUserAvatarResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) UpdateUserPassword(context.Context, *accountrpc.UpdateUserPasswordReq, ...grpc.CallOption) (*accountrpc.UpdateUserPasswordResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) BindUserEmail(context.Context, *accountrpc.BindUserEmailReq, ...grpc.CallOption) (*accountrpc.BindUserEmailResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) BindUserPhone(context.Context, *accountrpc.BindUserPhoneReq, ...grpc.CallOption) (*accountrpc.BindUserPhoneResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) BindUserOauth(context.Context, *accountrpc.BindUserOauthReq, ...grpc.CallOption) (*accountrpc.BindUserOauthResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) UnbindUserOauth(context.Context, *accountrpc.UnbindUserOauthReq, ...grpc.CallOption) (*accountrpc.UnbindUserOauthResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) AdminUpdateUserStatus(context.Context, *accountrpc.AdminUpdateUserStatusReq, ...grpc.CallOption) (*accountrpc.AdminUpdateUserStatusResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) AdminResetUserPassword(context.Context, *accountrpc.AdminResetUserPasswordReq, ...grpc.CallOption) (*accountrpc.AdminResetUserPasswordResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) FindUserList(context.Context, *accountrpc.FindUserListReq, ...grpc.CallOption) (*accountrpc.FindUserListResp, error) {
	return s.findUserResp, nil
}
func (s *stubVisitAccountRPC) FindUserInfoList(context.Context, *accountrpc.FindUserListReq, ...grpc.CallOption) (*accountrpc.FindUserInfoListResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) FindUserOnlineList(context.Context, *accountrpc.FindUserListReq, ...grpc.CallOption) (*accountrpc.FindUserInfoListResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) AnalysisUser(context.Context, *accountrpc.AnalysisUserReq, ...grpc.CallOption) (*accountrpc.AnalysisUserResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) AnalysisUserAreas(context.Context, *accountrpc.AnalysisUserAreasReq, ...grpc.CallOption) (*accountrpc.AnalysisUserAreasResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) GetClientInfo(context.Context, *accountrpc.GetClientInfoReq, ...grpc.CallOption) (*accountrpc.GetClientInfoResp, error) {
	panic("unexpected call")
}
func (s *stubVisitAccountRPC) FindVisitorList(context.Context, *accountrpc.FindVisitorListReq, ...grpc.CallOption) (*accountrpc.FindVisitorListResp, error) {
	return s.findVisitorResp, nil
}

func TestFindVisitLogListBuildsQuery(t *testing.T) {
	syslogRPC := &stubVisitSyslogRPC{
		findResp: &syslogrpc.FindVisitLogListResp{
			Pagination: &syslogrpc.PageResp{Page: 1, PageSize: 10, Total: 1},
			List:       []*syslogrpc.VisitLog{{Id: 1, UserId: "u-1", TerminalId: "t-1", PageName: "home"}},
		},
	}
	accountRPC := &stubVisitAccountRPC{
		findUserResp:    &accountrpc.FindUserListResp{List: []*accountrpc.User{{UserId: "u-1"}}},
		findVisitorResp: &accountrpc.FindVisitorListResp{List: []*accountrpc.VisitorInfo{{TerminalId: "t-1"}}},
	}
	logic := NewFindVisitLogListLogic(context.Background(), &svc.ServiceContext{SyslogRpc: syslogRPC, AccountRpc: accountRPC})

	resp, err := logic.FindVisitLogList(&types.QueryVisitLogReq{
		PageQuery:  types.PageQuery{Page: 1, PageSize: 10},
		UserId:     "u-1",
		TerminalId: "t-1",
		PageName:   "home",
	})
	if err != nil {
		t.Fatalf("FindVisitLogList returned error: %v", err)
	}
	if syslogRPC.findReq == nil || syslogRPC.findReq.UserId != "u-1" || syslogRPC.findReq.TerminalId != "t-1" {
		t.Fatalf("unexpected find request: %#v", syslogRPC.findReq)
	}
	list, ok := resp.List.([]*types.VisitLogBackVO)
	if !ok || len(list) != 1 || list[0].ClientInfo == nil {
		t.Fatalf("unexpected response list: %#v", resp.List)
	}
}
