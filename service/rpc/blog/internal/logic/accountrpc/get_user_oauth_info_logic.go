package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserOauthInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserOauthInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserOauthInfoLogic {
	return &GetUserOauthInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户第三平台信息
func (l *GetUserOauthInfoLogic) GetUserOauthInfo(in *accountrpc.GetUserOauthInfoReq) (*accountrpc.GetUserOauthInfoResp, error) {
	if in.UserId == "" {
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "用户id不能为空")
	}

	list, err := l.svcCtx.TUserOauthModel.FindByUserID(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	resp := make([]*accountrpc.UserOauthInfo, 0, len(list))
	for _, item := range list {
		resp = append(resp, convertUserOauthInfoOut(item))
	}

	return &accountrpc.GetUserOauthInfoResp{List: resp}, nil
}
