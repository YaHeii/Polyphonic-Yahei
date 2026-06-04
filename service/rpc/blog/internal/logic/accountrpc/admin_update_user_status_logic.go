package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminUpdateUserStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminUpdateUserStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminUpdateUserStatusLogic {
	return &AdminUpdateUserStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户状态
func (l *AdminUpdateUserStatusLogic) AdminUpdateUserStatus(in *accountrpc.AdminUpdateUserStatusReq) (*accountrpc.AdminUpdateUserStatusResp, error) {
	if in.UserId == "" {
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "用户id不能为空")
	}

	user, err := l.svcCtx.TUserModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, bizerr.NewBizError(bizcode.CodeUserNotExist, "用户不存在")
	}

	user.Status = in.Status

	if err := l.svcCtx.TUserModel.Update(l.ctx, user); err != nil {
		return nil, err
	}

	return &accountrpc.AdminUpdateUserStatusResp{}, nil
}
