package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *GetUserInfoLogic) GetUserInfo(in *accountrpc.GetUserInfoReq) (*accountrpc.GetUserInfoResp, error) {
	uid := in.UserId

	ui, err := l.svcCtx.TUserModel.FindOneByUserId(l.ctx, uid)
	if err != nil {
		return nil, err
	}

	// 查找用户角色
	rList, err := getUserRoles(l.ctx, l.svcCtx, uid)
	if err != nil {
		return nil, err
	}

	return &accountrpc.GetUserInfoResp{
		User: convertUserInfoOut(ui, rList),
	}, nil
}
