// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package auth

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClientInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取客户端信息
func NewGetClientInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClientInfoLogic {
	return &GetClientInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClientInfoLogic) GetClientInfo(req *types.GetClientInfoReq) (resp *types.GetClientInfoResp, err error) {
	out, err := l.svcCtx.AccountRpc.GetClientInfo(l.ctx, &accountrpc.GetClientInfoReq{})
	if err != nil {
		return nil, err
	}

	visitor := out.GetVisitor()
	return &types.GetClientInfoResp{
		Id:         visitor.GetId(),
		TerminalId: visitor.GetTerminalId(),
		Os:         visitor.GetOs(),
		Browser:    visitor.GetBrowser(),
		IpAddress:  visitor.GetIpAddress(),
		IpSource:   visitor.GetIpSource(),
	}, nil
}
