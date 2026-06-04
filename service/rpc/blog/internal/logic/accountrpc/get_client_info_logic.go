package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/ipx"
	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClientInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetClientInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClientInfoLogic {
	return &GetClientInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取客户端信息
func (l *GetClientInfoLogic) GetClientInfo(in *accountrpc.GetClientInfoReq) (*accountrpc.GetClientInfoResp, error) {
	terminalID := resolveTerminalID(l.ctx)
	if terminalID == "" {
		return &accountrpc.GetClientInfoResp{}, nil
	}

	visitor, err := l.svcCtx.TVisitorModel.FindOneByTerminalId(l.ctx, terminalID)
	if err == nil && visitor != nil {
		return &accountrpc.GetClientInfoResp{
			Visitor: convertVisitorInfoOut(visitor),
		}, nil
	}

	ip := getRemoteIPFromContext(l.ctx)
	visitor = &model.TVisitor{
		TerminalId: terminalID,
		Os:         resolveClientOS(l.ctx),
		Browser:    resolveClientBrowser(l.ctx),
		IpAddress:  ip,
		IpSource:   ipx.GetIpSourceByBaidu(ip),
	}
	if _, err := l.svcCtx.TVisitorModel.Insert(l.ctx, visitor); err != nil {
		return nil, err
	}

	visitor, err = l.svcCtx.TVisitorModel.FindOneByTerminalId(l.ctx, terminalID)
	if err != nil {
		return nil, err
	}

	return &accountrpc.GetClientInfoResp{
		Visitor: convertVisitorInfoOut(visitor),
	}, nil
}
