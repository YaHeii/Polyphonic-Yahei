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

type AddApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建api路由
func NewAddApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddApiLogic {
	return &AddApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddApiLogic) AddApi(req *types.NewApiReq) (resp *types.ApiBackVO, err error) {
	in := &permissionrpc.AddApiReq{
		Id:        req.Id,
		ParentId:  req.ParentId,
		Path:      req.Path,
		Name:      req.Name,
		Method:    req.Method,
		Traceable: req.Traceable,
		Status:    req.Status,
		Children:  nil,
	}

	out, err := l.svcCtx.PermissionRpc.AddApi(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convertApiTypes(out.Api), nil
}
