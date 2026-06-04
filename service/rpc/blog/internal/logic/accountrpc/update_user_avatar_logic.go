package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserAvatarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserAvatarLogic {
	return &UpdateUserAvatarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户头像
func (l *UpdateUserAvatarLogic) UpdateUserAvatar(in *accountrpc.UpdateUserAvatarReq) (*accountrpc.UpdateUserAvatarResp, error) {
	uid, err := getCurrentUserID(l.ctx)
	if err != nil {
		return nil, err
	}
	if in.Avatar == "" {
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "头像不能为空")
	}

	if err := l.svcCtx.TUserModel.UpdateAvatarByUserID(l.ctx, uid, in.Avatar); err != nil {
		return nil, err
	}

	return &accountrpc.UpdateUserAvatarResp{}, nil
}
