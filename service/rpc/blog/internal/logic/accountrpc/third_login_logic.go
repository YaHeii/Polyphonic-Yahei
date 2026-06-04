package accountrpclogic

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/YaHeii/Polyphonic-Yahei/common/enums"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/oauth"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/cryptox"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/ipx"
	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/common/rpcutils"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ThirdLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThirdLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThirdLoginLogic {
	return &ThirdLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 第三方登录
func (l *ThirdLoginLogic) ThirdLogin(in *accountrpc.ThirdLoginReq) (*accountrpc.LoginResp, error) {
	app, err := rpcutils.GetAppNameFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	auth, err := GetPlatformOauth(l.ctx, l.svcCtx, app, in.Platform)
	if err != nil {
		return nil, err
	}

	info, err := auth.GetAuthUserInfo(in.Code)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(info.OpenId) == "" {
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "open_id 不能为空")
	}

	userOauth, err := l.svcCtx.TUserOauthModel.FindOneByOpenIdPlatform(l.ctx, info.OpenId, in.Platform)
	switch {
	case err == nil:
	case errors.Is(err, model.ErrNotFound):
		userOauth, err = l.oauthRegister(in.Platform, info)
		if err != nil {
			return nil, err
		}
	default:
		return nil, err
	}

	user, err := l.svcCtx.TUserModel.FindOneByUserId(l.ctx, userOauth.UserId)
	if err != nil {
		return nil, bizerr.NewBizError(bizcode.CodeUserNotExist, "用户不存在")
	}

	return onLogin(l.ctx, l.svcCtx, user, enums.LoginTypeOauth)
}

func (l *ThirdLoginLogic) oauthRegister(platform string, info *oauth.UserResult) (*model.TUserOauth, error) {
	ip := getRemoteIPFromContext(l.ctx)
	user := &model.TUser{
		UserId:       strings.ReplaceAll(uuid.NewString(), "-", ""),
		Username:     l.generateOauthUsername(info),
		Password:     cryptox.BcryptHash(defaultOauthPassword(info)),
		Nickname:     defaultOauthNickname(info),
		Avatar:       info.Avatar,
		Email:        strings.TrimSpace(info.Email),
		Phone:        strings.TrimSpace(info.Mobile),
		Info:         "",
		Status:       enums.UserStatusNormal,
		RegisterType: platform,
		IpAddress:    ip,
		IpSource:     ipx.GetIpSourceByBaidu(ip),
	}

	userOauth := &model.TUserOauth{
		UserId:   user.UserId,
		Platform: platform,
		OpenId:   strings.TrimSpace(info.OpenId),
		Nickname: defaultOauthNickname(info),
		Avatar:   info.Avatar,
	}

	err := l.svcCtx.SqlConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		conn := sqlx.NewSqlConnFromSession(session)

		txUserModel := model.NewTUserModel(conn, buildRegisterModelCacheConf(l.svcCtx.Config.RedisConf))
		if _, err := txUserModel.Insert(ctx, user); err != nil {
			return err
		}

		roles, err := model.NewTRoleModel(conn, buildRegisterModelCacheConf(l.svcCtx.Config.RedisConf)).FindDefaultRoles(ctx)
		if err != nil {
			return err
		}

		txUserRoleModel := model.NewTUserRoleModel(conn)
		for _, role := range roles {
			if _, err := txUserRoleModel.Insert(ctx, &model.TUserRole{
				UserId: user.UserId,
				RoleId: role.Id,
			}); err != nil {
				return err
			}
		}

		txUserOauthModel := model.NewTUserOauthModel(conn, buildRegisterModelCacheConf(l.svcCtx.Config.RedisConf))
		_, err = txUserOauthModel.Insert(ctx, userOauth)
		return err
	})
	if err != nil {
		return nil, err
	}

	return l.svcCtx.TUserOauthModel.FindOneByOpenIdPlatform(l.ctx, userOauth.OpenId, platform)
}

func (l *ThirdLoginLogic) generateOauthUsername(info *oauth.UserResult) string {
	base := strings.TrimSpace(info.Name)
	if base == "" {
		base = strings.TrimSpace(info.EnName)
	}
	if base == "" {
		base = "oauth_" + info.OpenId
	}

	base = strings.ReplaceAll(base, " ", "_")
	if len(base) > 24 {
		base = base[:24]
	}

	username := base
	for i := 0; ; i++ {
		user, err := l.svcCtx.TUserModel.FindOneByUsername(l.ctx, username)
		if errors.Is(err, model.ErrNotFound) || user == nil {
			return username
		}

		suffix := cryptox.Md5v(info.OpenId, platformKey(info)+fmt.Sprintf("%d", i))
		username = base + "_" + suffix[:8]
	}
}

func defaultOauthNickname(info *oauth.UserResult) string {
	if nickname := strings.TrimSpace(info.NickName); nickname != "" {
		return nickname
	}
	if name := strings.TrimSpace(info.Name); name != "" {
		return name
	}
	if enName := strings.TrimSpace(info.EnName); enName != "" {
		return enName
	}
	return "oauth_user"
}

func defaultOauthPassword(info *oauth.UserResult) string {
	if source := strings.TrimSpace(info.EnName); source != "" {
		return source
	}
	if source := strings.TrimSpace(info.OpenId); source != "" {
		return source
	}
	return uuid.NewString()
}

func platformKey(info *oauth.UserResult) string {
	return strings.TrimSpace(info.Name) + strings.TrimSpace(info.EnName)
}
