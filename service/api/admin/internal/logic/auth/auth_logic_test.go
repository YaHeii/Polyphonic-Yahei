package auth

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/common/constant"
	"github.com/YaHeii/Polyphonic-Yahei/common/rediskey"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/captcha"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizheader"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/mail"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/oauth"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/tokenx"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/config"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/syslogrpc"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"google.golang.org/grpc"
)

type stubAuthAccountRPC struct {
	accountrpc.AccountRpc
	clientInfoReq  *accountrpc.GetClientInfoReq
	clientInfoResp *accountrpc.GetClientInfoResp
	registerReq    *accountrpc.RegisterReq
	resetReq       *accountrpc.ResetPasswordReq
	loginReq       *accountrpc.LoginReq
	emailLoginReq  *accountrpc.EmailLoginReq
	phoneLoginReq  *accountrpc.PhoneLoginReq
	thirdLoginReq  *accountrpc.ThirdLoginReq
	logoutReq      *accountrpc.LogoutReq
	logoffReq      *accountrpc.LogoffReq
	loginResp      *accountrpc.LoginResp
	emailLoginResp *accountrpc.LoginResp
	phoneLoginResp *accountrpc.LoginResp
	thirdLoginResp *accountrpc.LoginResp
}

func (s *stubAuthAccountRPC) GetClientInfo(_ context.Context, in *accountrpc.GetClientInfoReq, _ ...grpc.CallOption) (*accountrpc.GetClientInfoResp, error) {
	s.clientInfoReq = in
	return s.clientInfoResp, nil
}

func (s *stubAuthAccountRPC) Register(_ context.Context, in *accountrpc.RegisterReq, _ ...grpc.CallOption) (*accountrpc.RegisterResp, error) {
	s.registerReq = in
	return &accountrpc.RegisterResp{}, nil
}

func (s *stubAuthAccountRPC) ResetPassword(_ context.Context, in *accountrpc.ResetPasswordReq, _ ...grpc.CallOption) (*accountrpc.ResetPasswordResp, error) {
	s.resetReq = in
	return &accountrpc.ResetPasswordResp{}, nil
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

type stubEmailDeliver struct {
	message *mail.EmailMessage
}

func (s *stubEmailDeliver) DeliveryEmail(message *mail.EmailMessage) error {
	s.message = message
	return nil
}

type stubOauthProvider struct {
	info *oauth.UserResult
}

func (s *stubOauthProvider) GetName() string {
	return "github"
}

func (s *stubOauthProvider) GetAuthLoginUrl(state string) string {
	return "https://oauth.example?state=" + state
}

func (s *stubOauthProvider) GetAuthUserInfo(string) (*oauth.UserResult, error) {
	return s.info, nil
}

func TestAuthUtilityLocalMapping(t *testing.T) {
	accountRPC := &stubAuthAccountRPC{
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
	}
	svcCtx := &svc.ServiceContext{
		AccountRpc:     accountRPC,
		CaptchaHolder:  captcha.NewCaptchaHolder(),
		OauthProviders: map[string]oauth.Oauth{"admin-web:github": &stubOauthProvider{}},
	}
	ctx := context.Background()

	captchaResp, err := NewGetCaptchaCodeLogic(ctx, svcCtx).GetCaptchaCode(&types.GetCaptchaCodeReq{Width: 120, Height: 40})
	if err != nil {
		t.Fatalf("GetCaptchaCode returned error: %v", err)
	}
	if captchaResp.CaptchaKey == "" || captchaResp.CaptchaBase64 == "" || captchaResp.CaptchaCode == "" {
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
	if !strings.Contains(oauthResp.AuthorizeUrl, "state-1") {
		t.Fatalf("unexpected oauth response: %#v", oauthResp)
	}
}

func TestAuthWriteOperationsVerifyLocallyBeforeRPC(t *testing.T) {
	accountRPC := &stubAuthAccountRPC{}
	holder := captcha.NewCaptchaHolder()
	emailDeliver := &stubEmailDeliver{}
	svcCtx := &svc.ServiceContext{AccountRpc: accountRPC, CaptchaHolder: holder, EmailDeliver: emailDeliver}
	ctx := context.Background()

	registerCode, _ := holder.GetCodeCaptcha(rediskey.GetCaptchaKey(constant.CodeTypeRegister, "demo@example.com"))
	if _, err := NewRegisterLogic(ctx, svcCtx).Register(&types.RegisterReq{
		Username:   "demo",
		Password:   "secret",
		Email:      "demo@example.com",
		VerifyCode: registerCode,
	}); err != nil {
		t.Fatalf("Register returned error: %v", err)
	}
	if accountRPC.registerReq == nil || accountRPC.registerReq.Username != "demo" || accountRPC.registerReq.Email != "demo@example.com" {
		t.Fatalf("unexpected register request: %#v", accountRPC.registerReq)
	}

	resetCode, _ := holder.GetCodeCaptcha(rediskey.GetCaptchaKey(constant.CodeTypeResetPwd, "demo@example.com"))
	if _, err := NewResetPasswordLogic(ctx, svcCtx).ResetPassword(&types.ResetPasswordReq{
		Email:      "demo@example.com",
		Password:   "secret",
		VerifyCode: resetCode,
	}); err != nil {
		t.Fatalf("ResetPassword returned error: %v", err)
	}
	if accountRPC.resetReq == nil || accountRPC.resetReq.Email != "demo@example.com" || accountRPC.resetReq.Password != "secret" {
		t.Fatalf("unexpected reset request: %#v", accountRPC.resetReq)
	}

	if _, err := NewSendEmailVerifyCodeLogic(ctx, svcCtx).SendEmailVerifyCode(&types.SendEmailVerifyCodeReq{
		Email: "demo@example.com",
		Type:  constant.CodeTypeRegister,
	}); err != nil {
		t.Fatalf("SendEmailVerifyCode returned error: %v", err)
	}
	if emailDeliver.message == nil || len(emailDeliver.message.To) != 1 || emailDeliver.message.To[0] != "demo@example.com" {
		t.Fatalf("unexpected sent email: %#v", emailDeliver.message)
	}

	if _, err := NewSendPhoneVerifyCodeLogic(ctx, svcCtx).SendPhoneVerifyCode(&types.SendPhoneVerifyCodeReq{
		Phone: "18800000000",
		Type:  constant.CodeTypePhoneLogin,
	}); err != nil {
		t.Fatalf("SendPhoneVerifyCode returned error: %v", err)
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
}

func TestAuthLoginLifecycleMapping(t *testing.T) {
	accountRPC := &stubAuthAccountRPC{
		loginResp:      newLoginRPCResp("u-1", "password"),
		emailLoginResp: newLoginRPCResp("u-2", "email"),
		phoneLoginResp: newLoginRPCResp("u-3", "phone"),
		thirdLoginResp: newLoginRPCResp("u-4", "github"),
	}
	syslogRPC := &stubAuthSyslogRPC{}
	manager := tokenx.NewJwtTokenManager(&stubAuthTokenStore{}, "secret", "admin", 7200, 86400)
	holder := captcha.NewCaptchaHolder()
	svcCtx := &svc.ServiceContext{
		Config:          config.Config{RestConf: rest.RestConf{ServiceConf: service.ServiceConf{Name: "admin"}}},
		AccountRpc:      accountRPC,
		SyslogRpc:       syslogRPC,
		JwtTokenManager: manager,
		CaptchaHolder:   holder,
		OauthProviders: map[string]oauth.Oauth{"admin-web:github": &stubOauthProvider{
			info: &oauth.UserResult{OpenId: "open-1", NickName: "gh", Avatar: "avatar"},
		}},
	}
	ctx := context.Background()

	loginKey, loginCode, err := addCaptchaForTest(holder)
	if err != nil {
		t.Fatalf("captcha setup failed: %v", err)
	}
	loginResp, err := NewLoginLogic(ctx, svcCtx).Login(&types.LoginReq{Username: "demo", Password: "secret", CaptchaKey: loginKey, CaptchaCode: loginCode})
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

	emailKey, emailCode, err := addCaptchaForTest(holder)
	if err != nil {
		t.Fatalf("captcha setup failed: %v", err)
	}
	if _, err := NewEmailLoginLogic(ctx, svcCtx).EmailLogin(&types.EmailLoginReq{Email: "demo@example.com", Password: "secret", CaptchaKey: emailKey, CaptchaCode: emailCode}); err != nil {
		t.Fatalf("EmailLogin returned error: %v", err)
	}
	if accountRPC.emailLoginReq == nil || accountRPC.emailLoginReq.Email != "demo@example.com" {
		t.Fatalf("unexpected email login request: %#v", accountRPC.emailLoginReq)
	}

	phoneCode, _ := holder.GetCodeCaptcha(rediskey.GetCaptchaKey(constant.CodeTypePhoneLogin, "18800000000"))
	if _, err := NewPhoneLoginLogic(ctx, svcCtx).PhoneLogin(&types.PhoneLoginReq{Phone: "18800000000", VerifyCode: phoneCode}); err != nil {
		t.Fatalf("PhoneLogin returned error: %v", err)
	}
	if accountRPC.phoneLoginReq == nil || accountRPC.phoneLoginReq.Phone != "18800000000" {
		t.Fatalf("unexpected phone login request: %#v", accountRPC.phoneLoginReq)
	}

	if _, err := NewThirdLoginLogic(ctx, svcCtx).ThirdLogin(&types.ThirdLoginReq{Platform: "github", Code: "oauth-code"}); err != nil {
		t.Fatalf("ThirdLogin returned error: %v", err)
	}
	if accountRPC.thirdLoginReq == nil || accountRPC.thirdLoginReq.Platform != "github" || accountRPC.thirdLoginReq.OpenId != "open-1" {
		t.Fatalf("unexpected third login request: %#v", accountRPC.thirdLoginReq)
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

	if _, err := NewLogoutLogic(ctx, svcCtx).Logout(&types.EmptyReq{}); err != nil {
		t.Fatalf("Logout returned error: %v", err)
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
	if _, err := NewLogoffLogic(ctx, svcCtx).Logoff(&types.EmptyReq{}); err != nil {
		t.Fatalf("Logoff returned error: %v", err)
	}
	if accountRPC.logoffReq == nil || accountRPC.logoffReq.UserId != "u-9" {
		t.Fatalf("unexpected logoff request: %#v", accountRPC.logoffReq)
	}
	if len(store.deleteCalls) != 2 {
		t.Fatalf("unexpected delete calls: %#v", store.deleteCalls)
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
		User:      &accountrpc.UserInfo{UserId: userID},
		LoginType: loginType,
	}
}

func addCaptchaForTest(holder *captcha.CaptchaHolder) (string, string, error) {
	key, _, code, err := holder.GetMathImageCaptcha(40, 120)
	return key, code, err
}
