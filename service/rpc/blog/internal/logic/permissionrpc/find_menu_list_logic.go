package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMenuListLogic {
	return &FindMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询菜单列表
func (l *FindMenuListLogic) FindMenuList(in *permissionrpc.FindMenuListReq) (*permissionrpc.FindMenuListResp, error) {
	page, size, sorts, conditions, params := buildMenuQuery(in)
	records, total, err := l.svcCtx.TMenuModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	return &permissionrpc.FindMenuListResp{
		Pagination: buildPageResp(page, size, total),
		List:       buildMenuTree(records),
	}, nil
}
