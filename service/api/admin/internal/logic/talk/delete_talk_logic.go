// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package talk

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/socialrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTalkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除说说
func NewDeleteTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTalkLogic {
	return &DeleteTalkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTalkLogic) DeleteTalk(req *types.IdReq) (resp *types.BatchResp, err error) {
	in := &socialrpc.DeletesTalkReq{
		Ids: []int64{req.Id},
	}

	out, err := l.svcCtx.SocialRpc.DeletesTalk(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{SuccessCount: out.SuccessCount}, nil
}
