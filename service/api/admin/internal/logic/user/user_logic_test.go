package user

import (
	"context"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/common/constant"
	"github.com/YaHeii/Polyphonic-Yahei/common/rediskey"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/captcha"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizheader"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/oauth"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/syslogrpc"
	"google.golang.org/grpc"
)

type stubUserAccountRPC struct {
	accountrpc.AccountRpc
	getUserInfoReq    *accountrpc.GetUserInfoReq
	getUserInfoResp   *accountrpc.GetUserInfoResp
	getOauthReq       *accountrpc.GetUserOauthInfoReq
	getOauthResp      *accountrpc.GetUserOauthInfoResp
	updateAvatarReq   *accountrpc.UpdateUserAvatarReq
	updateInfoReq     *accountrpc.UpdateUserInfoReq
	updatePasswordReq *accountrpc.UpdateUserPasswordReq
	bindEmailReq      *accountrpc.BindUserEmailReq
	bindPhoneReq      *accountrpc.BindUserPhoneReq
	bindOauthReq      *accountrpc.BindUserOauthReq
	unbindOauthReq    *accountrpc.UnbindUserOauthReq
}

func (s *stubUserAccountRPC) GetUserInfo(_ context.Context, in *accountrpc.GetUserInfoReq, _ ...grpc.CallOption) (*accountrpc.GetUserInfoResp, error) {
	s.getUserInfoReq = in
	return s.getUserInfoResp, nil
}

func (s *stubUserAccountRPC) GetUserOauthInfo(_ context.Context, in *accountrpc.GetUserOauthInfoReq, _ ...grpc.CallOption) (*accountrpc.GetUserOauthInfoResp, error) {
	s.getOauthReq = in
	return s.getOauthResp, nil
}

func (s *stubUserAccountRPC) UpdateUserAvatar(_ context.Context, in *accountrpc.UpdateUserAvatarReq, _ ...grpc.CallOption) (*accountrpc.UpdateUserAvatarResp, error) {
	s.updateAvatarReq = in
	return &accountrpc.UpdateUserAvatarResp{}, nil
}

func (s *stubUserAccountRPC) UpdateUserInfo(_ context.Context, in *accountrpc.UpdateUserInfoReq, _ ...grpc.CallOption) (*accountrpc.UpdateUserInfoResp, error) {
	s.updateInfoReq = in
	return &accountrpc.UpdateUserInfoResp{}, nil
}

func (s *stubUserAccountRPC) UpdateUserPassword(_ context.Context, in *accountrpc.UpdateUserPasswordReq, _ ...grpc.CallOption) (*accountrpc.UpdateUserPasswordResp, error) {
	s.updatePasswordReq = in
	return &accountrpc.UpdateUserPasswordResp{}, nil
}

func (s *stubUserAccountRPC) BindUserEmail(_ context.Context, in *accountrpc.BindUserEmailReq, _ ...grpc.CallOption) (*accountrpc.BindUserEmailResp, error) {
	s.bindEmailReq = in
	return &accountrpc.BindUserEmailResp{}, nil
}

func (s *stubUserAccountRPC) BindUserPhone(_ context.Context, in *accountrpc.BindUserPhoneReq, _ ...grpc.CallOption) (*accountrpc.BindUserPhoneResp, error) {
	s.bindPhoneReq = in
	return &accountrpc.BindUserPhoneResp{}, nil
}

func (s *stubUserAccountRPC) BindUserOauth(_ context.Context, in *accountrpc.BindUserOauthReq, _ ...grpc.CallOption) (*accountrpc.BindUserOauthResp, error) {
	s.bindOauthReq = in
	return &accountrpc.BindUserOauthResp{}, nil
}

func (s *stubUserAccountRPC) UnbindUserOauth(_ context.Context, in *accountrpc.UnbindUserOauthReq, _ ...grpc.CallOption) (*accountrpc.UnbindUserOauthResp, error) {
	s.unbindOauthReq = in
	return &accountrpc.UnbindUserOauthResp{}, nil
}

type stubUserPermissionRPC struct {
	permissionrpc.PermissionRpc
	apisReq   *permissionrpc.FindUserApisReq
	apisResp  *permissionrpc.FindUserApisResp
	menusReq  *permissionrpc.FindUserMenusReq
	menusResp *permissionrpc.FindUserMenusResp
	rolesReq  *permissionrpc.FindUserRolesReq
	rolesResp *permissionrpc.FindUserRolesResp
}

func (s *stubUserPermissionRPC) FindUserApis(_ context.Context, in *permissionrpc.FindUserApisReq, _ ...grpc.CallOption) (*permissionrpc.FindUserApisResp, error) {
	s.apisReq = in
	return s.apisResp, nil
}

func (s *stubUserPermissionRPC) FindUserMenus(_ context.Context, in *permissionrpc.FindUserMenusReq, _ ...grpc.CallOption) (*permissionrpc.FindUserMenusResp, error) {
	s.menusReq = in
	return s.menusResp, nil
}

func (s *stubUserPermissionRPC) FindUserRoles(_ context.Context, in *permissionrpc.FindUserRolesReq, _ ...grpc.CallOption) (*permissionrpc.FindUserRolesResp, error) {
	s.rolesReq = in
	return s.rolesResp, nil
}

type stubUserSyslogRPC struct {
	syslogrpc.SyslogRpc
	findReq  *syslogrpc.FindLoginLogListReq
	findResp *syslogrpc.FindLoginLogListResp
}

type stubUserOauthProvider struct {
	info *oauth.UserResult
}

func (s *stubUserOauthProvider) GetName() string {
	return "github"
}

func (s *stubUserOauthProvider) GetAuthLoginUrl(state string) string {
	return "https://oauth.example?state=" + state
}

func (s *stubUserOauthProvider) GetAuthUserInfo(string) (*oauth.UserResult, error) {
	return s.info, nil
}

func (s *stubUserSyslogRPC) FindLoginLogList(_ context.Context, in *syslogrpc.FindLoginLogListReq, _ ...grpc.CallOption) (*syslogrpc.FindLoginLogListResp, error) {
	s.findReq = in
	return s.findResp, nil
}

func TestGetUserApisBuildsRequestAndMapsTree(t *testing.T) {
	ctx := context.WithValue(context.Background(), bizheader.HeaderUid, "u-1")
	permissionRPC := &stubUserPermissionRPC{
		apisResp: &permissionrpc.FindUserApisResp{
			List: []*permissionrpc.Api{
				{
					Id:       1,
					Name:     "root",
					Path:     "/root",
					Method:   "GET",
					Children: []*permissionrpc.Api{{Id: 2, ParentId: 1, Name: "child"}},
				},
			},
		},
	}
	logic := NewGetUserApisLogic(ctx, &svc.ServiceContext{PermissionRpc: permissionRPC})

	resp, err := logic.GetUserApis(&types.EmptyReq{})
	if err != nil {
		t.Fatalf("GetUserApis returned error: %v", err)
	}
	if permissionRPC.apisReq == nil || permissionRPC.apisReq.UserId != "u-1" {
		t.Fatalf("unexpected apis request: %#v", permissionRPC.apisReq)
	}
	if len(resp.List) != 1 || len(resp.List[0].Children) != 1 || resp.List[0].Children[0].Id != 2 {
		t.Fatalf("unexpected apis response: %#v", resp)
	}
}

func TestGetUserMenusBuildsRequestAndMapsTree(t *testing.T) {
	ctx := context.WithValue(context.Background(), bizheader.HeaderUid, "u-1")
	permissionRPC := &stubUserPermissionRPC{
		menusResp: &permissionrpc.FindUserMenusResp{
			List: []*permissionrpc.Menu{
				{
					Id:        1,
					Path:      "/dashboard",
					Name:      "Dashboard",
					Component: "Layout",
					Meta: &permissionrpc.MenuMeta{
						Title:      "Dashboard",
						Icon:       "home",
						Visible:    true,
						AlwaysShow: true,
					},
					Children: []*permissionrpc.Menu{
						{
							Id:   2,
							Name: "Analysis",
							Meta: &permissionrpc.MenuMeta{Title: "Analysis"},
						},
					},
				},
			},
		},
	}
	logic := NewGetUserMenusLogic(ctx, &svc.ServiceContext{PermissionRpc: permissionRPC})

	resp, err := logic.GetUserMenus(&types.EmptyReq{})
	if err != nil {
		t.Fatalf("GetUserMenus returned error: %v", err)
	}
	if permissionRPC.menusReq == nil || permissionRPC.menusReq.UserId != "u-1" {
		t.Fatalf("unexpected menus request: %#v", permissionRPC.menusReq)
	}
	if len(resp.List) != 1 || !resp.List[0].Meta.Hidden || !resp.List[0].Meta.AlwaysShow || len(resp.List[0].Children) != 1 {
		t.Fatalf("unexpected menus response: %#v", resp)
	}
}

func TestGetUserRolesBuildsRequestAndMapsRoles(t *testing.T) {
	ctx := context.WithValue(context.Background(), bizheader.HeaderUid, "u-1")
	permissionRPC := &stubUserPermissionRPC{
		rolesResp: &permissionrpc.FindUserRolesResp{
			List: []*permissionrpc.Role{
				{Id: 1, ParentId: 2, RoleKey: "admin", RoleLabel: "Admin", RoleComment: "root"},
			},
		},
	}
	logic := NewGetUserRolesLogic(ctx, &svc.ServiceContext{PermissionRpc: permissionRPC})

	resp, err := logic.GetUserRoles(&types.EmptyReq{})
	if err != nil {
		t.Fatalf("GetUserRoles returned error: %v", err)
	}
	if permissionRPC.rolesReq == nil || permissionRPC.rolesReq.UserId != "u-1" {
		t.Fatalf("unexpected roles request: %#v", permissionRPC.rolesReq)
	}
	if len(resp.List) != 1 || resp.List[0].RoleKey != "admin" {
		t.Fatalf("unexpected roles response: %#v", resp)
	}
}

func TestGetUserInfoBuildsRequestsAndAggregatesResponse(t *testing.T) {
	ctx := context.WithValue(context.Background(), bizheader.HeaderUid, "u-1")
	accountRPC := &stubUserAccountRPC{
		getUserInfoResp: &accountrpc.GetUserInfoResp{
			User: &accountrpc.UserInfo{
				UserId:       "u-1",
				Username:     "demo",
				Nickname:     "Demo",
				Avatar:       "avatar",
				Email:        "demo@example.com",
				Phone:        "188",
				RegisterType: "email",
				CreatedAt:    100,
				Info:         `{"intro":"hello","website":"https://example.com"}`,
				Roles: []*accountrpc.UserRoleLabel{
					{RoleKey: "admin", RoleLabel: "Admin"},
				},
			},
		},
		getOauthResp: &accountrpc.GetUserOauthInfoResp{
			List: []*accountrpc.UserOauthInfo{
				{Platform: "github", OpenId: "oid-1", Nickname: "gh", Avatar: "a", CreatedAt: 10},
			},
		},
	}
	permissionRPC := &stubUserPermissionRPC{
		menusResp: &permissionrpc.FindUserMenusResp{
			List: []*permissionrpc.Menu{
				{
					Id:   1,
					Name: "Root",
					Meta: &permissionrpc.MenuMeta{Perm: "system:user:list"},
					Children: []*permissionrpc.Menu{
						{Id: 2, Name: "Child", Meta: &permissionrpc.MenuMeta{Perm: "system:user:edit"}},
					},
				},
			},
		},
	}
	logic := NewGetUserInfoLogic(ctx, &svc.ServiceContext{
		AccountRpc:    accountRPC,
		PermissionRpc: permissionRPC,
	})

	resp, err := logic.GetUserInfo(&types.EmptyReq{})
	if err != nil {
		t.Fatalf("GetUserInfo returned error: %v", err)
	}
	if accountRPC.getUserInfoReq == nil || accountRPC.getUserInfoReq.UserId != "u-1" {
		t.Fatalf("unexpected get user info request: %#v", accountRPC.getUserInfoReq)
	}
	if accountRPC.getOauthReq == nil || accountRPC.getOauthReq.UserId != "u-1" {
		t.Fatalf("unexpected get oauth request: %#v", accountRPC.getOauthReq)
	}
	if permissionRPC.menusReq == nil || permissionRPC.menusReq.UserId != "u-1" {
		t.Fatalf("unexpected get menus request: %#v", permissionRPC.menusReq)
	}
	if resp.UserId != "u-1" || resp.UserInfoExt.Intro != "hello" || len(resp.ThirdParty) != 1 || len(resp.Roles) != 1 || len(resp.Perms) != 2 {
		t.Fatalf("unexpected user info response: %#v", resp)
	}
}

func TestGetUserLoginHistoryListBuildsRequestAndMapsPage(t *testing.T) {
	ctx := context.WithValue(context.Background(), bizheader.HeaderUid, "u-1")
	syslogRPC := &stubUserSyslogRPC{
		findResp: &syslogrpc.FindLoginLogListResp{
			Pagination: &syslogrpc.PageResp{Page: 1, PageSize: 10, Total: 1},
			List: []*syslogrpc.LoginLog{
				{Id: 1, UserId: "u-1", TerminalId: "t-1", LoginType: "password", AppName: "admin", LoginAt: 10, LogoutAt: 20},
			},
		},
	}
	logic := NewGetUserLoginHistoryListLogic(ctx, &svc.ServiceContext{SyslogRpc: syslogRPC})

	resp, err := logic.GetUserLoginHistoryList(&types.QueryUserLoginHistoryReq{
		PageQuery: types.PageQuery{Page: 1, PageSize: 10, Sorts: []string{"login_at desc"}},
	})
	if err != nil {
		t.Fatalf("GetUserLoginHistoryList returned error: %v", err)
	}
	if syslogRPC.findReq == nil || syslogRPC.findReq.UserId != "u-1" {
		t.Fatalf("unexpected login history request: %#v", syslogRPC.findReq)
	}
	list, ok := resp.List.([]*types.UserLoginHistory)
	if !ok || len(list) != 1 || list[0].Id != 1 {
		t.Fatalf("unexpected login history response: %#v", resp)
	}
}

func TestUserUpdateOperationsBuildRequests(t *testing.T) {
	accountRPC := &stubUserAccountRPC{}
	ctx := context.Background()

	avatarLogic := NewUpdateUserAvatarLogic(ctx, &svc.ServiceContext{AccountRpc: accountRPC})
	if _, err := avatarLogic.UpdateUserAvatar(&types.UpdateUserAvatarReq{Avatar: "avatar"}); err != nil {
		t.Fatalf("UpdateUserAvatar returned error: %v", err)
	}
	if accountRPC.updateAvatarReq == nil || accountRPC.updateAvatarReq.Avatar != "avatar" {
		t.Fatalf("unexpected avatar request: %#v", accountRPC.updateAvatarReq)
	}

	infoLogic := NewUpdateUserInfoLogic(ctx, &svc.ServiceContext{AccountRpc: accountRPC})
	if _, err := infoLogic.UpdateUserInfo(&types.UpdateUserInfoReq{
		Nickname: "demo",
		UserInfoExt: types.UserInfoExt{
			Intro:   "hello",
			Website: "https://example.com",
		},
	}); err != nil {
		t.Fatalf("UpdateUserInfo returned error: %v", err)
	}
	if accountRPC.updateInfoReq == nil || accountRPC.updateInfoReq.Nickname != "demo" || accountRPC.updateInfoReq.Info == "" {
		t.Fatalf("unexpected user info request: %#v", accountRPC.updateInfoReq)
	}

	passwordLogic := NewUpdateUserPasswordLogic(ctx, &svc.ServiceContext{AccountRpc: accountRPC})
	if _, err := passwordLogic.UpdateUserPassword(&types.UpdateUserPasswordReq{
		OldPassword: "old",
		NewPassword: "new",
	}); err != nil {
		t.Fatalf("UpdateUserPassword returned error: %v", err)
	}
	if accountRPC.updatePasswordReq == nil || accountRPC.updatePasswordReq.OldPassword != "old" || accountRPC.updatePasswordReq.NewPassword != "new" {
		t.Fatalf("unexpected password request: %#v", accountRPC.updatePasswordReq)
	}
}

func TestUserBindOperationsBuildRequests(t *testing.T) {
	accountRPC := &stubUserAccountRPC{}
	holder := captcha.NewCaptchaHolder()
	ctx := context.Background()
	svcCtx := &svc.ServiceContext{
		AccountRpc:     accountRPC,
		CaptchaHolder:  holder,
		OauthProviders: map[string]oauth.Oauth{"admin-web:github": &stubUserOauthProvider{info: &oauth.UserResult{OpenId: "open-1", NickName: "gh", Avatar: "avatar"}}},
	}

	emailCode, _ := holder.GetCodeCaptcha(rediskey.GetCaptchaKey(constant.CodeTypeBindEmail, "demo@example.com"))
	emailLogic := NewUpdateUserBindEmailLogic(ctx, svcCtx)
	if _, err := emailLogic.UpdateUserBindEmail(&types.UpdateUserBindEmailReq{Email: "demo@example.com", VerifyCode: emailCode}); err != nil {
		t.Fatalf("UpdateUserBindEmail returned error: %v", err)
	}
	if accountRPC.bindEmailReq == nil || accountRPC.bindEmailReq.Email != "demo@example.com" {
		t.Fatalf("unexpected bind email request: %#v", accountRPC.bindEmailReq)
	}

	phoneCode, _ := holder.GetCodeCaptcha(rediskey.GetCaptchaKey(constant.CodeTypeBindPhone, "18800000000"))
	phoneLogic := NewUpdateUserBindPhoneLogic(ctx, svcCtx)
	if _, err := phoneLogic.UpdateUserBindPhone(&types.UpdateUserBindPhoneReq{Phone: "18800000000", VerifyCode: phoneCode}); err != nil {
		t.Fatalf("UpdateUserBindPhone returned error: %v", err)
	}
	if accountRPC.bindPhoneReq == nil || accountRPC.bindPhoneReq.Phone != "18800000000" {
		t.Fatalf("unexpected bind phone request: %#v", accountRPC.bindPhoneReq)
	}

	oauthLogic := NewUpdateUserBindThirdPartyLogic(ctx, svcCtx)
	if _, err := oauthLogic.UpdateUserBindThirdParty(&types.UpdateUserBindThirdPartyReq{Platform: "github", Code: "code", State: "state"}); err != nil {
		t.Fatalf("UpdateUserBindThirdParty returned error: %v", err)
	}
	if accountRPC.bindOauthReq == nil || accountRPC.bindOauthReq.Platform != "github" || accountRPC.bindOauthReq.OpenId != "open-1" || accountRPC.bindOauthReq.Nickname != "gh" {
		t.Fatalf("unexpected bind oauth request: %#v", accountRPC.bindOauthReq)
	}

	unbindLogic := NewDeleteUserBindThirdPartyLogic(ctx, &svc.ServiceContext{AccountRpc: accountRPC})
	if _, err := unbindLogic.DeleteUserBindThirdParty(&types.DeleteUserBindThirdPartyReq{Platform: "github"}); err != nil {
		t.Fatalf("DeleteUserBindThirdParty returned error: %v", err)
	}
	if accountRPC.unbindOauthReq == nil || accountRPC.unbindOauthReq.Platform != "github" {
		t.Fatalf("unexpected unbind oauth request: %#v", accountRPC.unbindOauthReq)
	}
}
