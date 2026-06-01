package websiterpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/websiterpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVisitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddVisitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVisitLogic {
	return &AddVisitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加用户访问记录
func (l *AddVisitLogic) AddVisit(in *websiterpc.AddVisitReq) (*websiterpc.AddVisitResp, error) {
	// todo: add your logic here and delete this line

	return &websiterpc.AddVisitResp{}, nil
}
