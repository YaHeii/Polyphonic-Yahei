package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserInfoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserInfoListLogic {
	return &FindUserInfoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查找用户信息列表
func (l *FindUserInfoListLogic) FindUserInfoList(in *accountrpc.FindUserListReq) (*accountrpc.FindUserInfoListResp, error) {
	// todo: add your logic here and delete this line

	return &accountrpc.FindUserInfoListResp{}, nil
}
