package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnbindUserOauthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnbindUserOauthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnbindUserOauthLogic {
	return &UnbindUserOauthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 解绑第三方账号
func (l *UnbindUserOauthLogic) UnbindUserOauth(in *accountrpc.UnbindUserOauthReq) (*accountrpc.UnbindUserOauthResp, error) {
	user, err := getCurrentUser(l.ctx, l.svcCtx)
	if err != nil {
		return nil, err
	}
	if in.Platform == "" {
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "平台不能为空")
	}

	oauth, err := l.svcCtx.TUserOauthModel.FindOneByUserIdPlatform(l.ctx, user.UserId, in.Platform)
	if err != nil {
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "绑定关系不存在")
	}
	if err := l.svcCtx.TUserOauthModel.Delete(l.ctx, oauth.Id); err != nil {
		return nil, err
	}

	return &accountrpc.UnbindUserOauthResp{}, nil
}
