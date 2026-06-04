package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户信息
func (l *UpdateUserInfoLogic) UpdateUserInfo(in *accountrpc.UpdateUserInfoReq) (*accountrpc.UpdateUserInfoResp, error) {
	uid, err := getCurrentUserID(l.ctx)
	if err != nil {
		return nil, err
	}
	if in.Nickname == "" {
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "昵称不能为空")
	}

	if err := l.svcCtx.TUserModel.UpdateNicknameInfo(l.ctx, uid, in.Nickname, in.Info); err != nil {
		return nil, err
	}

	return &accountrpc.UpdateUserInfoResp{}, nil
}
