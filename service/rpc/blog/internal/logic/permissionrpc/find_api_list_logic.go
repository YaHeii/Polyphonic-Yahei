package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindApiListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindApiListLogic {
	return &FindApiListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询接口列表
func (l *FindApiListLogic) FindApiList(in *permissionrpc.FindApiListReq) (*permissionrpc.FindApiListResp, error) {
	page, size, sorts, conditions, params := buildApiQuery(in)
	records, total, err := l.svcCtx.TApiModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	return &permissionrpc.FindApiListResp{
		Pagination: buildPageResp(page, size, total),
		List:       buildApiTree(records),
	}, nil
}
