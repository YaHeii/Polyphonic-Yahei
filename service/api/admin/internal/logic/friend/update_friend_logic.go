// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package friend

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/socialrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新友链
func NewUpdateFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFriendLogic {
	return &UpdateFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFriendLogic) UpdateFriend(req *types.NewFriendReq) (resp *types.FriendBackVO, err error) {
	in := &socialrpc.UpdateFriendReq{
		Id:          req.Id,
		LinkName:    req.LinkName,
		LinkAvatar:  req.LinkAvatar,
		LinkAddress: req.LinkAddress,
		LinkIntro:   req.LinkIntro,
	}

	out, err := l.svcCtx.SocialRpc.UpdateFriend(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convertFriendTypes(out.Friend), nil
}
