package auth

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/authctx"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/tokenx"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
)

func onLogin(_ context.Context, svcCtx *svc.ServiceContext, login *accountrpc.LoginResp) (*types.LoginResp, error) {
	tk, err := svcCtx.JwtTokenManager.GenerateToken(login.GetUser().GetUserId())
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		UserId: login.GetUser().GetUserId(),
		Scope:  svcCtx.Config.Name,
		Token:  convertToken(tk),
	}, nil
}

func convertToken(tk *tokenx.Token) *types.Token {
	if tk == nil {
		return nil
	}

	return &types.Token{
		TokenType:        tk.TokenType,
		AccessToken:      tk.AccessToken,
		ExpiresIn:        tk.ExpiresIn,
		RefreshToken:     tk.RefreshToken,
		RefreshExpiresIn: tk.RefreshExpiresIn,
		RefreshExpiresAt: tk.RefreshExpiresAt,
	}
}

func currentUserID(ctx context.Context) string {
	return authctx.CurrentUserID(ctx)
}
