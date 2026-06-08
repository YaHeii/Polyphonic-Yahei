// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户头像
func NewUpdateUserAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserAvatarLogic {
	return &UpdateUserAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserAvatarLogic) UpdateUserAvatar(req *types.UpdateUserAvatarReq) (resp *types.EmptyResp, err error) {
	_, err = l.svcCtx.AccountRpc.UpdateUserAvatar(l.ctx, &accountrpc.UpdateUserAvatarReq{
		Avatar: req.Avatar,
	})
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
