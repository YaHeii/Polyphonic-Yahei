// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package website

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSystemStateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取服务器信息
func NewGetSystemStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSystemStateLogic {
	return &GetSystemStateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSystemStateLogic) GetSystemState(req *types.EmptyReq) (resp *types.Server, err error) {
	// todo: add your logic here and delete this line

	return
}
