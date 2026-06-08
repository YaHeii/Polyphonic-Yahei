// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package api

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/permissionrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除api路由
func NewDeletesApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesApiLogic {
	return &DeletesApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletesApiLogic) DeletesApi(req *types.IdsReq) (resp *types.BatchResp, err error) {
	in := &permissionrpc.DeletesApiReq{
		Ids: req.Ids,
	}

	out, err := l.svcCtx.PermissionRpc.DeletesApi(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}
	return resp, nil
}
