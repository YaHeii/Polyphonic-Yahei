package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/common/constant"
	"github.com/YaHeii/Polyphonic-Yahei/common/rediskey"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/captcha"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizheader"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/mail"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/oauth"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/authctx"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/tokenx"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/config"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/syslogrpc"
	"google.golang.org/grpc"
)

type stubAuthAccountRPC struct {
	accountrpc.AccountRpc
	loginReq         *accountrpc.LoginReq
	loginResp        *accountrpc.LoginResp
	loginErr         error
	emailLoginReq    *accountrpc.EmailLoginReq
	emailLoginResp   *accountrpc.LoginResp
	phoneLoginReq    *accountrpc.PhoneLoginReq
	phoneLoginResp   *accountrpc.LoginResp
	thirdLoginReq    *accountrpc.ThirdLoginReq
	thirdLoginResp   *accountrpc.LoginResp
	registerReq      *accountrpc.RegisterReq
	resetPasswordReq *accountrpc.ResetPasswordReq
	logoutReq        *accountrpc.LogoutReq
	logoffReq        *accountrpc.LogoffReq
	clientInfoResp   *accountrpc.GetClientInfoResp
}

func (s *stubAuthAccountRPC) Login(_ context.Context, in *accountrpc.LoginReq, _ ...grpc.CallOption) (*accountrpc.LoginResp, error) {
	s.loginReq = in
	return s.loginResp, s.loginErr
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

func (s *stubAuthAccountRPC) Register(_ context.Context, in *accountrpc.RegisterReq, _ ...grpc.CallOption) (*accountrpc.RegisterResp, error) {
	s.registerReq = in
	return &accountrpc.RegisterResp{}, nil
}

func (s *stubAuthAccountRPC) ResetPassword(_ context.Context, in *accountrpc.ResetPasswordReq, _ ...grpc.CallOption) (*accountrpc.ResetPasswordResp, error) {
	s.resetPasswordReq = in
	return &accountrpc.ResetPasswordResp{}, nil
}

func (s *stubAuthAccountRPC) Logout(_ context.Context, in *accountrpc.LogoutReq, _ ...grpc.CallOption) (*accountrpc.LogoutResp, error) {
	s.logoutReq = in
	return &accountrpc.LogoutResp{}, nil
}

func (s *stubAuthAccountRPC) Logoff(_ context.Context, in *accountrpc.LogoffReq, _ ...grpc.CallOption) (*accountrpc.LogoffResp, error) {
	s.logoffReq = in
	return &accountrpc.LogoffResp{}, nil
}

func (s *stubAuthAccountRPC) GetClientInfo(_ context.Context, in *accountrpc.GetClientInfoReq, _ ...grpc.CallOption) (*accountrpc.GetClientInfoResp, error) {
	return s.clientInfoResp, nil
}

type stubAuthSyslogRPC struct {
	syslogrpc.SyslogRpc
	addLoginReq  *syslogrpc.AddLoginLogReq
	addLogoutReq *syslogrpc.AddLogoutLogReq
}

func (s *stubAuthSyslogRPC) AddLoginLog(_ context.Context, in *syslogrpc.AddLoginLogReq, _ ...grpc.CallOption) (*syslogrpc.AddLoginLogResp, error) {
	s.addLoginReq = in
	return &syslogrpc.AddLoginLogResp{}, nil
}

func (s *stubAuthSyslogRPC) AddLogoutLog(_ context.Context, in *syslogrpc.AddLogoutLogReq, _ ...grpc.CallOption) (*syslogrpc.AddLogoutLogResp, error) {
	s.addLogoutReq = in
	return &syslogrpc.AddLogoutLogResp{}, nil
}

type stubAuthMailer struct {
	msg *mail.EmailMessage
}

func (s *stubAuthMailer) DeliveryEmail(msg *mail.EmailMessage) error {
	s.msg = msg
	return nil
}

type stubAuthOauthProvider struct {
	url  string
	info *oauth.UserResult
	err  error
}

type stubAuthTokenStore struct {
	data map[string]string
}

func (s *stubAuthOauthProvider) GetName() string {
	return "github"
}

func (s *stubAuthOauthProvider) GetAuthLoginUrl(state string) string {
	return s.url + state
}

func (s *stubAuthOauthProvider) GetAuthUserInfo(code string) (*oauth.UserResult, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.info, nil
}

func (s *stubAuthTokenStore) Set(key string, value string, expireSeconds int) error {
	if s.data == nil {
		s.data = make(map[string]string)
	}
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

func newAuthTokenManager(t *testing.T) *tokenx.JwtTokenManager {
	t.Helper()
	return tokenx.NewJwtTokenManager(&stubAuthTokenStore{}, "secret", "admin-api", 900, 2592000)
}

func newAuthServiceContext(t *testing.T) (*svc.ServiceContext, *stubAuthAccountRPC, *stubAuthSyslogRPC, *captcha.CaptchaHolder) {
	t.Helper()
	accountRPC := &stubAuthAccountRPC{
		loginResp: &accountrpc.LoginResp{
			User:      &accountrpc.UserInfo{UserId: "u-1"},
			LoginType: "password",
		},
		emailLoginResp: &accountrpc.LoginResp{
			User:      &accountrpc.UserInfo{UserId: "u-2"},
			LoginType: "email",
		},
		phoneLoginResp: &accountrpc.LoginResp{
			User:      &accountrpc.UserInfo{UserId: "u-3"},
			LoginType: "phone",
		},
		thirdLoginResp: &accountrpc.LoginResp{
			User:      &accountrpc.UserInfo{UserId: "u-4"},
			LoginType: "oauth",
		},
		clientInfoResp: &accountrpc.GetClientInfoResp{
			Visitor: &accountrpc.VisitorInfo{
				Id:         1,
				TerminalId: "t-1",
				Os:         "linux",
				Browser:    "chrome",
				IpAddress:  "127.0.0.1",
				IpSource:   "localhost",
			},
		},
	}
	syslogRPC := &stubAuthSyslogRPC{}
	holder := captcha.NewCaptchaHolder()
	cfg := config.Config{}
	cfg.Name = "admin-api"
	svcCtx := &svc.ServiceContext{
		Config:          cfg,
		AccountRpc:      accountRPC,
		SyslogRpc:       syslogRPC,
		CaptchaHolder:   holder,
		JwtTokenManager: newAuthTokenManager(t),
		OauthProviders: map[string]oauth.Oauth{
			"admin-web:github": &stubAuthOauthProvider{
				url:  "https://oauth.example/login?state=",
				info: &oauth.UserResult{OpenId: "oid-1", NickName: "gh", Avatar: "avatar"},
			},
		},
		EmailDeliver: &stubAuthMailer{},
	}
	return svcCtx, accountRPC, syslogRPC, holder
}

func seedCaptcha(t *testing.T, holder *captcha.CaptchaHolder, key string, code string) {
	t.Helper()
	got, err := holder.GetCodeCaptcha(key)
	if err != nil {
		t.Fatalf("seed captcha failed: %v", err)
	}
	if code != "" && got != code {
		t.Fatalf("unexpected captcha code: want %q got %q", code, got)
	}
}

func TestLoginReturnsBearerTokenAndWritesLoginLog(t *testing.T) {
	svcCtx, accountRPC, syslogRPC, holder := newAuthServiceContext(t)
	key := "captcha-login"
	code, err := holder.GetCodeCaptcha(key)
	if err != nil {
		t.Fatalf("GetCodeCaptcha returned error: %v", err)
	}

	logic := NewLoginLogic(context.Background(), svcCtx)
	resp, err := logic.Login(&types.LoginReq{
		Username:    "demo",
		Password:    "secret",
		CaptchaKey:  key,
		CaptchaCode: code,
	})
	if err != nil {
		t.Fatalf("Login returned error: %v", err)
	}
	if accountRPC.loginReq == nil || accountRPC.loginReq.Username != "demo" {
		t.Fatalf("unexpected login request: %#v", accountRPC.loginReq)
	}
	if syslogRPC.addLoginReq == nil || syslogRPC.addLoginReq.UserId != "u-1" {
		t.Fatalf("unexpected login log request: %#v", syslogRPC.addLoginReq)
	}
	if resp.Token == nil || resp.Token.TokenType != tokenx.TokenTypeBearer || resp.Token.AccessToken == "" || resp.Token.RefreshToken == "" {
		t.Fatalf("unexpected token bundle: %#v", resp)
	}
	if resp.Token.ExpiresIn != 900 {
		t.Fatalf("unexpected access ttl: %#v", resp.Token)
	}
}

func TestRefreshTokenRotatesAndReturnsUserID(t *testing.T) {
	svcCtx, _, _, _ := newAuthServiceContext(t)
	token, err := svcCtx.JwtTokenManager.GenerateToken("u-1")
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}

	logic := NewRefreshTokenLogic(context.Background(), svcCtx)
	resp, err := logic.RefreshToken(&types.RefreshTokenReq{RefreshToken: token.RefreshToken})
	if err != nil {
		t.Fatalf("RefreshToken returned error: %v", err)
	}
	if resp.UserId != "u-1" || resp.Token == nil {
		t.Fatalf("unexpected refresh response: %#v", resp)
	}
	if resp.Token.AccessToken == token.AccessToken || resp.Token.RefreshToken == token.RefreshToken {
		t.Fatalf("expected rotated tokens, got seed=%#v refreshed=%#v", token, resp.Token)
	}
	if _, err := svcCtx.JwtTokenManager.RefreshToken(token.RefreshToken); !errors.Is(err, tokenx.ErrTokenInvalid) && !errors.Is(err, tokenx.ErrTokenExpired) {
		t.Fatalf("expected old refresh token invalid or expired, got %v", err)
	}
}

func TestLogoutRevokesRefreshTokenUsingJWTUserID(t *testing.T) {
	svcCtx, accountRPC, syslogRPC, _ := newAuthServiceContext(t)
	token, err := svcCtx.JwtTokenManager.GenerateToken("u-1")
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}
	ctx := context.WithValue(context.Background(), authctx.UserIDKey, "u-1")

	logic := NewLogoutLogic(ctx, svcCtx)
	if _, err := logic.Logout(&types.EmptyReq{}); err != nil {
		t.Fatalf("Logout returned error: %v", err)
	}
	if accountRPC.logoutReq == nil || accountRPC.logoutReq.UserId != "u-1" {
		t.Fatalf("unexpected logout request: %#v", accountRPC.logoutReq)
	}
	if syslogRPC.addLogoutReq == nil || syslogRPC.addLogoutReq.UserId != "u-1" {
		t.Fatalf("unexpected logout log request: %#v", syslogRPC.addLogoutReq)
	}
	if _, err := svcCtx.JwtTokenManager.RefreshToken(token.RefreshToken); !errors.Is(err, tokenx.ErrTokenExpired) {
		t.Fatalf("expected revoked refresh token expired, got %v", err)
	}
}

func TestLogoffRevokesRefreshTokenUsingJWTUserID(t *testing.T) {
	svcCtx, accountRPC, syslogRPC, _ := newAuthServiceContext(t)
	token, err := svcCtx.JwtTokenManager.GenerateToken("u-1")
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}
	ctx := context.WithValue(context.Background(), authctx.UserIDKey, "u-1")

	logic := NewLogoffLogic(ctx, svcCtx)
	if _, err := logic.Logoff(&types.EmptyReq{}); err != nil {
		t.Fatalf("Logoff returned error: %v", err)
	}
	if accountRPC.logoffReq == nil || accountRPC.logoffReq.UserId != "u-1" {
		t.Fatalf("unexpected logoff request: %#v", accountRPC.logoffReq)
	}
	if syslogRPC.addLogoutReq == nil || syslogRPC.addLogoutReq.UserId != "u-1" {
		t.Fatalf("unexpected logout log request: %#v", syslogRPC.addLogoutReq)
	}
	if _, err := svcCtx.JwtTokenManager.RefreshToken(token.RefreshToken); !errors.Is(err, tokenx.ErrTokenExpired) {
		t.Fatalf("expected revoked refresh token expired, got %v", err)
	}
}

func TestRegisterAndResetPasswordUseVerificationCodeOnlyAtAPIBoundary(t *testing.T) {
	svcCtx, accountRPC, _, holder := newAuthServiceContext(t)
	registerKey := rediskey.GetCaptchaKey(constant.CodeTypeRegister, "demo@example.com")
	resetKey := rediskey.GetCaptchaKey(constant.CodeTypeResetPwd, "demo@example.com")
	registerCode, err := holder.GetCodeCaptcha(registerKey)
	if err != nil {
		t.Fatalf("GetCodeCaptcha returned error: %v", err)
	}
	resetCode, err := holder.GetCodeCaptcha(resetKey)
	if err != nil {
		t.Fatalf("GetCodeCaptcha returned error: %v", err)
	}

	registerLogic := NewRegisterLogic(context.Background(), svcCtx)
	if _, err := registerLogic.Register(&types.RegisterReq{
		Username:        "demo",
		Password:        "secret",
		ConfirmPassword: "secret",
		Email:           "demo@example.com",
		VerifyCode:      registerCode,
	}); err != nil {
		t.Fatalf("Register returned error: %v", err)
	}
	if accountRPC.registerReq == nil || accountRPC.registerReq.Email != "demo@example.com" || accountRPC.registerReq.Username != "demo" {
		t.Fatalf("unexpected register request: %#v", accountRPC.registerReq)
	}

	resetLogic := NewResetPasswordLogic(context.Background(), svcCtx)
	if _, err := resetLogic.ResetPassword(&types.ResetPasswordReq{
		Password:        "new-secret",
		ConfirmPassword: "new-secret",
		Email:           "demo@example.com",
		VerifyCode:      resetCode,
	}); err != nil {
		t.Fatalf("ResetPassword returned error: %v", err)
	}
	if accountRPC.resetPasswordReq == nil || accountRPC.resetPasswordReq.Email != "demo@example.com" || accountRPC.resetPasswordReq.Password != "new-secret" {
		t.Fatalf("unexpected reset request: %#v", accountRPC.resetPasswordReq)
	}
}

func TestThirdLoginUsesAppNameFromContext(t *testing.T) {
	svcCtx, accountRPC, syslogRPC, _ := newAuthServiceContext(t)
	ctx := context.WithValue(context.Background(), authctx.UserIDKey, "ignored")
	ctx = context.WithValue(ctx, bizheader.HeaderAppName, "admin-web")

	logic := NewThirdLoginLogic(ctx, svcCtx)
	resp, err := logic.ThirdLogin(&types.ThirdLoginReq{Platform: "github", Code: "oauth-code"})
	if err != nil {
		t.Fatalf("ThirdLogin returned error: %v", err)
	}
	if accountRPC.thirdLoginReq == nil || accountRPC.thirdLoginReq.OpenId != "oid-1" {
		t.Fatalf("unexpected third login request: %#v", accountRPC.thirdLoginReq)
	}
	if syslogRPC.addLoginReq == nil || syslogRPC.addLoginReq.UserId != "u-4" {
		t.Fatalf("unexpected login log request: %#v", syslogRPC.addLoginReq)
	}
	if resp.Token == nil || resp.Token.TokenType != tokenx.TokenTypeBearer {
		t.Fatalf("unexpected third login response: %#v", resp)
	}
}

func TestGetOauthAuthorizeURLUsesRequestMetaAppName(t *testing.T) {
	svcCtx, _, _, _ := newAuthServiceContext(t)
	ctx := context.WithValue(context.Background(), bizheader.HeaderAppName, "admin-web")

	logic := NewGetOauthAuthorizeUrlLogic(ctx, svcCtx)
	resp, err := logic.GetOauthAuthorizeUrl(&types.GetOauthAuthorizeUrlReq{
		Platform: "github",
		State:    "abc",
	})
	if err != nil {
		t.Fatalf("GetOauthAuthorizeUrl returned error: %v", err)
	}
	if resp.AuthorizeUrl != "https://oauth.example/login?state=abc" {
		t.Fatalf("unexpected authorize url: %#v", resp)
	}
}

func TestGetClientInfoMapsRPCResponse(t *testing.T) {
	svcCtx, _, _, _ := newAuthServiceContext(t)

	logic := NewGetClientInfoLogic(context.Background(), svcCtx)
	resp, err := logic.GetClientInfo(&types.GetClientInfoReq{})
	if err != nil {
		t.Fatalf("GetClientInfo returned error: %v", err)
	}
	if resp.Id != 1 || resp.TerminalId != "t-1" || resp.IpAddress != "127.0.0.1" {
		t.Fatalf("unexpected client info: %#v", resp)
	}
}
