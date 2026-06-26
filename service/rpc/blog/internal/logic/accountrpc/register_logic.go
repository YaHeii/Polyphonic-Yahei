package accountrpclogic

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/YaHeii/Polyphonic-Yahei/common/enums"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/cryptox"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/ipx"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/patternx"
	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注册
func (l *RegisterLogic) Register(in *accountrpc.RegisterReq) (*accountrpc.RegisterResp, error) {
	// 验证邮箱格式
	if !patternx.IsValidEmail(in.Email) {
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "邮箱格式不正确")
	}
	// 查找name是否已经被注册
	exist, err := l.svcCtx.TUserModel.FindOneByUsername(l.ctx, in.Username)
	if err == nil && exist != nil {
		return nil, bizerr.NewBizError(bizcode.CodeUserAlreadyExist, "用户名已被注册")
	}
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, err
	}
	// 验证邮箱是否已被注册
	exist, err = l.svcCtx.TUserModel.FindOneByEmail(l.ctx, in.Email)
	if err == nil && exist != nil {
		return nil, bizerr.NewBizError(bizcode.CodeUserAlreadyExist, "邮箱已被注册")
	}
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, err
	}
	user, err := l.register(in)
	if err != nil {
		return nil, err
	}

	roles, err := l.svcCtx.TRoleModel.FindRolesByUserID(l.ctx, user.UserId)
	if err != nil {
		return nil, err
	}

	return &accountrpc.RegisterResp{
		User: convertUserInfoOut(user, roles),
	}, nil
}

func (l *RegisterLogic) register(in *accountrpc.RegisterReq) (*model.TUser, error) {
	ip := getRemoteIPFromContext(l.ctx)
	user := &model.TUser{
		UserId:       strings.ReplaceAll(uuid.NewString(), "-", ""),
		Username:     in.Username,
		Password:     cryptox.BcryptHash(in.Password),
		Nickname:     in.Email,
		Avatar:       "https://mms1.baidu.com/it/u=2815887849,1501151317&fm=253&app=138&f=JPEG",
		Email:        in.Email,
		Phone:        "",
		Info:         "",
		Status:       enums.UserStatusNormal,
		RegisterType: enums.LoginTypeEmail,
		IpAddress:    ip,
		IpSource:     ipx.GetIpSourceByBaidu(ip),
	}

	err := l.svcCtx.SqlConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		conn := sqlx.NewSqlConnFromSession(session)

		txUserModel := model.NewTUserModel(conn, buildRegisterModelCacheConf(l.svcCtx.Config.RedisConf))
		visitorRole, err := model.NewTRoleModel(conn, buildRegisterModelCacheConf(l.svcCtx.Config.RedisConf)).FindByRoleKey(ctx, "visitor")
		if err != nil {
			return err
		}
		user.RoleId = visitorRole.Id

		if _, err := txUserModel.Insert(ctx, user); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return l.svcCtx.TUserModel.FindOneByEmail(l.ctx, in.Email)
}
