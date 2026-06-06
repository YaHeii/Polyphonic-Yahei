package socialrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/socialrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindFriendListLogic {
	return &FindFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询友链列表
func (l *FindFriendListLogic) FindFriendList(in *socialrpc.FindFriendListReq) (*socialrpc.FindFriendListResp, error) {
	page, size, sorts, conditions, params := buildFriendQuery(in)
	records, total, err := l.svcCtx.TFriendModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	return &socialrpc.FindFriendListResp{
		Pagination: buildPageResp(page, size, total),
		List:       convertFriendListOut(records),
	}, nil
}
