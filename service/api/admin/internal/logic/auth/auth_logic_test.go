package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizheader"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/tokenx"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/config"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/syslogrpc"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/core/service"
	"google.golang.org/grpc"
)

type stubAuthAccountRPC struct {
	accountrpc.AccountRpc
	captchaReq      *accountrpc.GenerateCaptchaCodeReq
	captchaResp     *accountrpc.GenerateCaptchaCodeResp
	clientInfoReq   *accountrpc.GetClientInfoReq
	clientInfoResp  *accountrpc.GetClientInfoResp
	oauthReq        *accountrpc.GetOauthAuthorizeUrlReq
	oauthResp       *accountrpc.GetOauthAuthorizeUrlResp
	registerReq     *accountrpc.RegisterReq
	resetReq        *accountrpc.ResetPasswordReq
	sendEmailReq    *accountrpc.SendEmailVerifyCodeReq
	sendPhoneReq    *accountrpc.SendPhoneVerifyCodeReq
	loginReq        *accountrpc.LoginReq
	emailLoginReq   *accountrpc.EmailLoginReq
	phoneLoginReq   *accountrpc.PhoneLoginReq
	thirdLoginReq   *accountrpc.ThirdLoginReq
	logoutReq       *accountrpc.LogoutReq
	logoffReq       *accountrpc.LogoffReq
	loginResp       *accountrpc.LoginResp
	emailLoginResp  *accountrpc.LoginResp
	phoneLoginResp  *accountrpc.LoginResp
	thirdLoginResp  *accountrpc.LoginResp
}

func (s *stubAuthAccountRPC) GenerateCaptchaCode(_ context.Context, in *accountrpc.GenerateCaptchaCodeReq, _ ...grpc.CallOption) (*accountrpc.GenerateCaptchaCodeResp, error) {
	s.captchaReq = in
	return s.captchaResp, nil
}

func (s *stubAuthAccountRPC) GetClientInfo(_ context.Context, in *accountrpc.GetClientInfoReq, _ ...grpc.CallOption) (*accountrpc.GetClientInfoResp, error) {
	s.clientInfoReq = in
	return s.clientInfoResp, nil
}

func (s *stubAuthAccountRPC) GetOauthAuthorizeUrl(_ context.Context, in *accountrpc.GetOauthAuthorizeUrlReq, _ ...grpc.CallOption) (*accountrpc.GetOauthAuthorizeUrlResp, error) {
	s.oauthReq = in
	return s.oauthResp, nil
}

func (s *stubAuthAccountRPC) Register(_ context.Context, in *accountrpc.RegisterReq, _ ...grpc.CallOption) (*accountrpc.RegisterResp, error) {
	s.registerReq = in
	return &accountrpc.RegisterResp{}, nil
}

func (s *stubAuthAccountRPC) ResetPassword(_ context.Context, in *accountrpc.ResetPasswordReq, _ ...grpc.CallOption) (*accountrpc.ResetPasswordResp, error) {
	s.resetReq = in
	return &accountrpc.ResetPasswordResp{}, nil
}

func (s *stubAuthAccountRPC) SendEmailVerifyCode(_ context.Context, in *accountrpc.SendEmailVerifyCodeReq, _ ...grpc.CallOption) (*accountrpc.SendEmailVerifyCodeResp, error) {
	s.sendEmailReq = in
	return &accountrpc.SendEmailVerifyCodeResp{}, nil
}

func (s *stubAuthAccountRPC) SendPhoneVerifyCode(_ context.Context, in *accountrpc.SendPhoneVerifyCodeReq, _ ...grpc.CallOption) (*accountrpc.SendPhoneVerifyCodeResp, error) {
	s.sendPhoneReq = in
	return &accountrpc.SendPhoneVerifyCodeResp{}, nil
}

func (s *stubAuthAccountRPC) Login(_ context.Context, in *accountrpc.LoginReq, _ ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	s.loginReq = in
	return s.loginResp, nil
}

func (s *stubAuthAccountRPC) EmailLogin(_ context.Context, in *accountrpc.EmailLoginReq, _ ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	s.emailLoginReq = in
	return s.emailLoginResp, nil
}

func (s *stubAuthAccountRPC) PhoneLogin(_ context.Context, in *accountrpc.PhoneLoginReq, _ ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	s.phoneLoginReq = in
	return s.phoneLoginResp, nil
}

func (s *stubAuthAccountRPC) ThirdLogin(_ context.Context, in *accountrpc.ThirdLoginReq, _ ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	s.thirdLoginReq = in
	return s.thirdLoginResp, nil
}

func (s *stubAuthAccountRPC) Logout(_ context.Context, in *accountrpc.LogoutReq, _ ...grpc.CallOption) (*accountrpc.LogoutResp, error) {
	s.logoutReq = in
	return &accountrpc.LogoutResp{}, nil
}

func (s *stubAuthAccountRPC) Logoff(_ context.Context, in *accountrpc.LogoffReq, _ ...grpc.CallOption) (*accountrpc.LogoffResp, error) {
	s.logoffReq = in
	return &accountrpc.LogoffResp{}, nil
}

type stubAuthSyslogRPC struct {
	syslogrpc.SyslogRpc
	loginReq  *syslogrpc.AddLoginLogReq
	logoutReq *syslogrpc.AddLogoutLogReq
}

func (s *stubAuthSyslogRPC) AddLoginLog(_ context.Context, in *syslogrpc.AddLoginLogReq, _ ...grpc.CallOption) (*syslogrpc.AddLoginLogResp, error) {
	s.loginReq = in
	return &syslogrpc.AddLoginLogResp{}, nil
}

func (s *stubAuthSyslogRPC) AddLogoutLog(_ context.Context, in *syslogrpc.AddLogoutLogReq, _ ...grpc.CallOption) (*syslogrpc.AddLogoutLogResp, error) {
	s.logoutReq = in
	return &syslogrpc.AddLogoutLogResp{}, nil
}

type stubAuthTokenStore struct {
	data        map[string]string
	setCalls    []setCall
	deleteCalls []string
	setErr      error
}

type setCall struct {
	key    string
	value  string
	expire int
}

type revokeCall struct {
	uid       string
	isRefresh bool
}

func (s *stubAuthTokenStore) Set(key string, value string, expireSeconds int) error {
	if s.setErr != nil {
		return s.setErr
	}
	if s.data == nil {
		s.data = make(map[string]string)
	}
	s.setCalls = append(s.setCalls, setCall{key: key, value: value, expire: expireSeconds})
	s.data[key] = value
	return nil
}

func (s *stubAuthTokenStore) Get(key string) (string, error) {
	if s.data == nil {
		return "", nil
	}
	return s.data[key], nil
}

func (s *stubAuthTokenStore) Delete(key string) error {
	s.deleteCalls = append(s.deleteCalls, key)
	if s.data != nil {
		delete(s.data, key)
	}
	return nil
}

func (s *stubAuthTokenStore) Exists(key string) (bool, error) {
	if s.data == nil {
		return false, nil
	}
	_, ok := s.data[key]
	return ok, nil
}

func (s *stubAuthTokenStore) SetExpire(string, int) error {
	return nil
}

func TestAuthUtilityRPCMapping(t *testing.T) {
	accountRPC := &stubAuthAccountRPC{
		captchaResp: &accountrpc.GenerateCaptchaCodeResp{
			CaptchaKey:    "key",
			CaptchaBase64: "base64",
			CaptchaCode:   "1234",
		},
		clientInfoResp: &accountrpc.GetClientInfoResp{
			Visitor: &accountrpc.VisitorInfo{
				Id:         1,
				TerminalId: "t-1",
				Os:         "mac",
				Browser:    "chrome",
				IpAddress:  "127.0.0.1",
				IpSource:   "local",
			},
		},
		oauthResp: &accountrpc.GetOauthAuthorizeUrlResp{AuthorizeUrl: "https://oauth.example"},
	}
	ctx := context.Background()
	svcCtx := &svc.ServiceContext{AccountRpc: accountRPC}

	captchaResp, err := NewGetCaptchaCodeLogic(ctx, svcCtx).GetCaptchaCode(&types.GetCaptchaCodeReq{Width: 120, Height: 40})
	if err != nil {
		t.Fatalf("GetCaptchaCode returned error: %v", err)
	}
	if accountRPC.captchaReq == nil || accountRPC.captchaReq.Width != 120 || accountRPC.captchaReq.Height != 40 {
		t.Fatalf("unexpected captcha request: %#v", accountRPC.captchaReq)
	}
	if captchaResp.CaptchaKey != "key" || captchaResp.CaptchaBase64 != "base64" || captchaResp.CaptchaCode != "1234" {
		t.Fatalf("unexpected captcha response: %#v", captchaResp)
	}

	clientResp, err := NewGetClientInfoLogic(ctx, svcCtx).GetClientInfo(&types.GetClientInfoReq{})
	if err != nil {
		t.Fatalf("GetClientInfo returned error: %v", err)
	}
	if accountRPC.clientInfoReq == nil {
		t.Fatal("expected client info request to be sent")
	}
	if clientResp.TerminalId != "t-1" || clientResp.Os != "mac" || clientResp.IpSource != "local" {
		t.Fatalf("unexpected client info response: %#v", clientResp)
	}

	oauthResp, err := NewGetOauthAuthorizeUrlLogic(ctx, svcCtx).GetOauthAuthorizeUrl(&types.GetOauthAuthorizeUrlReq{
		Platform: "github",
		State:    "state-1",
	})
	if err != nil {
		t.Fatalf("GetOauthAuthorizeUrl returned error: %v", err)
	}
	if accountRPC.oauthReq == nil || accountRPC.oauthReq.Platform != "github" || accountRPC.oauthReq.State != "state-1" {
		t.Fatalf("unexpected oauth request: %#v", accountRPC.oauthReq)
	}
	if oauthResp.AuthorizeUrl != "https://oauth.example" {
		t.Fatalf("unexpected oauth response: %#v", oauthResp)
	}
}

func TestAuthUtilityWriteRPCMapping(t *testing.T) {
	accountRPC := &stubAuthAccountRPC{}
	ctx := context.Background()
	svcCtx := &svc.ServiceContext{AccountRpc: accountRPC}

	registerResp, err := NewRegisterLogic(ctx, svcCtx).Register(&types.RegisterReq{
		Username:   "demo",
		Password:   "secret",
		Email:      "demo@example.com",
		VerifyCode: "123456",
	})
	if err != nil {
		t.Fatalf("Register returned error: %v", err)
	}
	if registerResp == nil {
		t.Fatal("expected register response")
	}
	if accountRPC.registerReq == nil || accountRPC.registerReq.Username != "demo" || accountRPC.registerReq.VerifyCode != "123456" {
		t.Fatalf("unexpected register request: %#v", accountRPC.registerReq)
	}

	resetResp, err := NewResetPasswordLogic(ctx, svcCtx).ResetPassword(&types.ResetPasswordReq{
		Email:      "demo@example.com",
		Password:   "secret",
		VerifyCode: "654321",
	})
	if err != nil {
		t.Fatalf("ResetPassword returned error: %v", err)
	}
	if resetResp == nil {
		t.Fatal("expected reset password response")
	}
	if accountRPC.resetReq == nil || accountRPC.resetReq.Email != "demo@example.com" || accountRPC.resetReq.VerifyCode != "654321" {
		t.Fatalf("unexpected reset request: %#v", accountRPC.resetReq)
	}

	emailResp, err := NewSendEmailVerifyCodeLogic(ctx, svcCtx).SendEmailVerifyCode(&types.SendEmailVerifyCodeReq{
		Email: "demo@example.com",
		Type:  "register",
	})
	if err != nil {
		t.Fatalf("SendEmailVerifyCode returned error: %v", err)
	}
	if emailResp == nil {
		t.Fatal("expected send email verify code response")
	}
	if accountRPC.sendEmailReq == nil || accountRPC.sendEmailReq.Email != "demo@example.com" || accountRPC.sendEmailReq.Type != "register" {
		t.Fatalf("unexpected send email request: %#v", accountRPC.sendEmailReq)
	}

	phoneResp, err := NewSendPhoneVerifyCodeLogic(ctx, svcCtx).SendPhoneVerifyCode(&types.SendPhoneVerifyCodeReq{
		Phone: "18800000000",
		Type:  "reset_password",
	})
	if err != nil {
		t.Fatalf("SendPhoneVerifyCode returned error: %v", err)
	}
	if phoneResp == nil {
		t.Fatal("expected send phone verify code response")
	}
	if accountRPC.sendPhoneReq == nil || accountRPC.sendPhoneReq.Phone != "18800000000" || accountRPC.sendPhoneReq.Type != "reset_password" {
		t.Fatalf("unexpected send phone request: %#v", accountRPC.sendPhoneReq)
	}
}

func TestAuthUtilityTokenMapping(t *testing.T) {
	store := &stubAuthTokenStore{}
	manager := tokenx.NewJwtTokenManager(store, "secret", "admin", 7200, 86400)
	seed, err := manager.GenerateToken("u-1")
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}
	logic := NewRefreshTokenLogic(context.Background(), &svc.ServiceContext{
		Config:          config.Config{RestConf: rest.RestConf{ServiceConf: service.ServiceConf{Name: "admin"}}},
		JwtTokenManager: manager,
	})

	resp, err := logic.RefreshToken(&types.RefreshTokenReq{
		UserId:       "u-1",
		RefreshToken: seed.RefreshToken,
	})
	if err != nil {
		t.Fatalf("RefreshToken returned error: %v", err)
	}
	if resp.UserId != "u-1" || resp.Scope != "admin" || resp.Token == nil {
		t.Fatalf("unexpected refresh response: %#v", resp)
	}
	if resp.Token.AccessToken == "" || resp.Token.RefreshToken == "" {
		t.Fatalf("unexpected refresh response: %#v", resp)
	}
	if store.data[tokenx.JwtAccessKey("u-1")] != resp.Token.AccessToken || store.data[tokenx.JwtRefreshKey("u-1")] != resp.Token.RefreshToken {
		t.Fatalf("unexpected token store state: %#v", store.data)
	}
}

func TestAuthLoginLifecycleMapping(t *testing.T) {
	accountRPC := &stubAuthAccountRPC{
		loginResp:      newLoginRPCResp("u-1", "password"),
		emailLoginResp: newLoginRPCResp("u-2", "email"),
		phoneLoginResp: newLoginRPCResp("u-3", "phone"),
		thirdLoginResp: newLoginRPCResp("u-4", "github"),
	}
	syslogRPC := &stubAuthSyslogRPC{}
	store := &stubAuthTokenStore{}
	manager := tokenx.NewJwtTokenManager(store, "secret", "admin", 7200, 86400)
	svcCtx := &svc.ServiceContext{
		Config:          config.Config{RestConf: rest.RestConf{ServiceConf: service.ServiceConf{Name: "admin"}}},
		AccountRpc:      accountRPC,
		SyslogRpc:       syslogRPC,
		JwtTokenManager: manager,
	}
	ctx := context.Background()

	loginResp, err := NewLoginLogic(ctx, svcCtx).Login(&types.LoginReq{Username: "demo", Password: "secret"})
	if err != nil {
		t.Fatalf("Login returned error: %v", err)
	}
	if accountRPC.loginReq == nil || accountRPC.loginReq.Username != "demo" {
		t.Fatalf("unexpected login request: %#v", accountRPC.loginReq)
	}
	if syslogRPC.loginReq == nil || syslogRPC.loginReq.UserId != "u-1" || syslogRPC.loginReq.LoginType != "password" {
		t.Fatalf("unexpected login log request: %#v", syslogRPC.loginReq)
	}
	if loginResp.UserId != "u-1" || loginResp.Scope != "admin" || loginResp.Token == nil {
		t.Fatalf("unexpected login response: %#v", loginResp)
	}

	_, err = NewEmailLoginLogic(ctx, svcCtx).EmailLogin(&types.EmailLoginReq{Email: "demo@example.com", Password: "secret"})
	if err != nil {
		t.Fatalf("EmailLogin returned error: %v", err)
	}
	if accountRPC.emailLoginReq == nil || accountRPC.emailLoginReq.Email != "demo@example.com" {
		t.Fatalf("unexpected email login request: %#v", accountRPC.emailLoginReq)
	}
	if syslogRPC.loginReq == nil || syslogRPC.loginReq.UserId != "u-2" || syslogRPC.loginReq.LoginType != "email" {
		t.Fatalf("unexpected email login log request: %#v", syslogRPC.loginReq)
	}

	_, err = NewPhoneLoginLogic(ctx, svcCtx).PhoneLogin(&types.PhoneLoginReq{Phone: "18800000000", VerifyCode: "123456"})
	if err != nil {
		t.Fatalf("PhoneLogin returned error: %v", err)
	}
	if accountRPC.phoneLoginReq == nil || accountRPC.phoneLoginReq.Phone != "18800000000" || accountRPC.phoneLoginReq.VerifyCode != "123456" {
		t.Fatalf("unexpected phone login request: %#v", accountRPC.phoneLoginReq)
	}
	if syslogRPC.loginReq == nil || syslogRPC.loginReq.UserId != "u-3" || syslogRPC.loginReq.LoginType != "phone" {
		t.Fatalf("unexpected phone login log request: %#v", syslogRPC.loginReq)
	}

	_, err = NewThirdLoginLogic(ctx, svcCtx).ThirdLogin(&types.ThirdLoginReq{Platform: "github", Code: "oauth-code"})
	if err != nil {
		t.Fatalf("ThirdLogin returned error: %v", err)
	}
	if accountRPC.thirdLoginReq == nil || accountRPC.thirdLoginReq.Platform != "github" || accountRPC.thirdLoginReq.Code != "oauth-code" {
		t.Fatalf("unexpected third login request: %#v", accountRPC.thirdLoginReq)
	}
	if syslogRPC.loginReq == nil || syslogRPC.loginReq.UserId != "u-4" || syslogRPC.loginReq.LoginType != "github" {
		t.Fatalf("unexpected third login log request: %#v", syslogRPC.loginReq)
	}
}

func TestAuthLogoutLifecycleMapping(t *testing.T) {
	accountRPC := &stubAuthAccountRPC{}
	syslogRPC := &stubAuthSyslogRPC{}
	store := &stubAuthTokenStore{
		data: map[string]string{
			tokenx.JwtAccessKey("u-9"):  "access",
			tokenx.JwtRefreshKey("u-9"): "refresh",
		},
	}
	manager := tokenx.NewJwtTokenManager(store, "secret", "admin", 7200, 86400)
	ctx := context.WithValue(context.Background(), bizheader.HeaderUid, "u-9")
	svcCtx := &svc.ServiceContext{
		AccountRpc:      accountRPC,
		SyslogRpc:       syslogRPC,
		JwtTokenManager: manager,
	}

	logoutResp, err := NewLogoutLogic(ctx, svcCtx).Logout(&types.EmptyReq{})
	if err != nil {
		t.Fatalf("Logout returned error: %v", err)
	}
	if logoutResp == nil {
		t.Fatal("expected logout response")
	}
	if accountRPC.logoutReq == nil || accountRPC.logoutReq.UserId != "u-9" {
		t.Fatalf("unexpected logout request: %#v", accountRPC.logoutReq)
	}
	if syslogRPC.logoutReq == nil || syslogRPC.logoutReq.UserId != "u-9" || syslogRPC.logoutReq.LogoutAt <= 0 {
		t.Fatalf("unexpected logout log request: %#v", syslogRPC.logoutReq)
	}
	if len(store.deleteCalls) != 1 || store.deleteCalls[0] != tokenx.JwtAccessKey("u-9") {
		t.Fatalf("unexpected delete calls: %#v", store.deleteCalls)
	}

	store.deleteCalls = nil
	store.data[tokenx.JwtAccessKey("u-9")] = "access"
	store.data[tokenx.JwtRefreshKey("u-9")] = "refresh"
	logoffResp, err := NewLogoffLogic(ctx, svcCtx).Logoff(&types.EmptyReq{})
	if err != nil {
		t.Fatalf("Logoff returned error: %v", err)
	}
	if logoffResp == nil {
		t.Fatal("expected logoff response")
	}
	if accountRPC.logoffReq == nil || accountRPC.logoffReq.UserId != "u-9" {
		t.Fatalf("unexpected logoff request: %#v", accountRPC.logoffReq)
	}
	if syslogRPC.logoutReq == nil || syslogRPC.logoutReq.UserId != "u-9" || syslogRPC.logoutReq.LogoutAt <= 0 {
		t.Fatalf("unexpected logoff log request: %#v", syslogRPC.logoutReq)
	}
	if len(store.deleteCalls) != 2 {
		t.Fatalf("unexpected delete calls: %#v", store.deleteCalls)
	}
	if store.deleteCalls[0] != tokenx.JwtAccessKey("u-9") || store.deleteCalls[1] != tokenx.JwtRefreshKey("u-9") {
		t.Fatalf("unexpected delete order: %#v", store.deleteCalls)
	}
}

func TestOnLoginPropagatesTokenErrors(t *testing.T) {
	manager := tokenx.NewJwtTokenManager(&stubAuthTokenStore{setErr: errors.New("store failed")}, "secret", "admin", 7200, 86400)
	_, err := onLogin(context.Background(), &svc.ServiceContext{
		Config:          config.Config{RestConf: rest.RestConf{ServiceConf: service.ServiceConf{Name: "admin"}}},
		JwtTokenManager: manager,
	}, newLoginRPCResp("u-err", "password"))
	if err == nil {
		t.Fatal("expected onLogin error")
	}
}

func newLoginRPCResp(userID, loginType string) *accountrpc.LoginResp {
	return &accountrpc.LoginResp{
		User: &accountrpc.UserInfo{UserId: userID},
		LoginType: loginType,
	}
}
