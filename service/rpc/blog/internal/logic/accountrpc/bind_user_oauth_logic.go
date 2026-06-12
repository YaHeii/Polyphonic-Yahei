package accountrpclogic

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/common/rpcutils"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type BindUserOauthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindUserOauthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindUserOauthLogic {
	return &BindUserOauthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户第三方账号
func (l *BindUserOauthLogic) BindUserOauth(in *accountrpc.BindUserOauthReq) (*accountrpc.BindUserOauthResp, error) {
	// 查找当前用户是否存在
	userId, err := rpcutils.GetUserIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	openID := strings.TrimSpace(in.OpenId)
	if openID == "" {
		return nil, fmt.Errorf("open_id is empty")
	}

	current, err := l.svcCtx.TUserOauthModel.FindOneByUserIdPlatform(l.ctx, userId, in.Platform)
	if err == nil && current != nil {
		return nil, fmt.Errorf("platform %s is already bound", in.Platform)
	}
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, err
	}

	oa, err := l.svcCtx.TUserOauthModel.FindOneByOpenIdPlatform(l.ctx, openID, in.Platform)
	if err == nil && oa != nil {
		return nil, fmt.Errorf("open_id %s is already exist", openID)
	}
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, err
	}

	// 绑定第三方账号
	_, err = l.svcCtx.TUserOauthModel.Insert(l.ctx, &model.TUserOauth{
		Id:       0,
		UserId:   userId,
		Platform: in.Platform,
		OpenId:   openID,
		Nickname: in.Nickname,
		Avatar:   in.Avatar,
	})

	return &accountrpc.BindUserOauthResp{}, nil
}
