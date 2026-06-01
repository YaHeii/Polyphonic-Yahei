package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindVisitorListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindVisitorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindVisitorListLogic {
	return &FindVisitorListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询游客信息
func (l *FindVisitorListLogic) FindVisitorList(in *accountrpc.FindVisitorListReq) (*accountrpc.FindVisitorListResp, error) {
	// todo: add your logic here and delete this line

	return &accountrpc.FindVisitorListResp{}, nil
}
