package accountrpclogic

import (
	"context"
	"errors"
	"strings"

	"github.com/YaHeii/Polyphonic-Yahei/common/enums"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/service/model"
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
	openID := strings.TrimSpace(in.OpenId)
	if openID == "" {
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "open_id 不能为空")
	}

	userOauth, err := l.svcCtx.TUserOauthModel.FindOneByOpenIdPlatform(l.ctx, openID, in.Platform)
	if errors.Is(err, model.ErrNotFound) {
		return nil, bizerr.NewBizError(bizcode.CodeUserNotExist, "第三方账号未绑定邮箱账号")
	}
	if err != nil {
		return nil, err
	}

	user, err := l.svcCtx.TUserModel.FindOneByUserId(l.ctx, userOauth.UserId)
	if err != nil {
		return nil, bizerr.NewBizError(bizcode.CodeUserNotExist, "用户不存在")
	}

	return onLogin(l.ctx, l.svcCtx, user, enums.LoginTypeOauth)
}
