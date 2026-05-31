// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package auth

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThirdLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 第三方登录
func NewThirdLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThirdLoginLogic {
	return &ThirdLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThirdLoginLogic) ThirdLogin(req *types.ThirdLoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
