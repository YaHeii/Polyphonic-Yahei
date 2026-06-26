package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type LogoffLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoffLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoffLogic {
	return &LogoffLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注销
func (l *LogoffLogic) Logoff(in *accountrpc.LogoffReq) (*accountrpc.LogoffResp, error) {
	if in.UserId == "" {
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "用户id不能为空")
	}

	user, err := l.svcCtx.TUserModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, bizerr.NewBizError(bizcode.CodeUserNotExist, "用户不存在")
	}

	if err := l.logoffUser(user); err != nil {
		return nil, err
	}

	return &accountrpc.LogoffResp{}, nil
}

func (l *LogoffLogic) logoffUser(user *model.TUser) error {
	err := l.svcCtx.SqlConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		conn := sqlx.NewSqlConnFromSession(session)

		txUserOauthModel := model.NewTUserOauthModel(conn, buildRegisterModelCacheConf(l.svcCtx.Config.RedisConf))
		oauthList, err := txUserOauthModel.FindByUserID(ctx, user.UserId)
		if err != nil {
			return err
		}
		for _, oauth := range oauthList {
			if err := txUserOauthModel.Delete(ctx, oauth.Id); err != nil {
				return err
			}
		}

		txUserModel := model.NewTUserModel(conn, buildRegisterModelCacheConf(l.svcCtx.Config.RedisConf))
		return txUserModel.Delete(ctx, user.Id)
	})
	if err != nil {
		return err
	}

	return l.svcCtx.OnlineUserService.Logout(l.ctx, user.UserId)
}
