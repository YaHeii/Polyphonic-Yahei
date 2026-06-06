package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoleListLogic {
	return &FindRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询角色列表
func (l *FindRoleListLogic) FindRoleList(in *permissionrpc.FindRoleListReq) (*permissionrpc.FindRoleListResp, error) {
	page, size, sorts, conditions, params := buildRoleQuery(in)
	records, total, err := l.svcCtx.TRoleModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	return &permissionrpc.FindRoleListResp{
		Pagination: buildPageResp(page, size, total),
		List:       buildRoleTree(records),
	}, nil
}
