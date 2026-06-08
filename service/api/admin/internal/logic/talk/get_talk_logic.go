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

type GetTalkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询说说
func NewGetTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTalkLogic {
	return &GetTalkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTalkLogic) GetTalk(req *types.IdReq) (resp *types.TalkBackVO, err error) {
	in := &socialrpc.GetTalkReq{Id: req.Id}

	out, err := l.svcCtx.SocialRpc.GetTalk(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convertTalkTypes(out.Talk), nil
}
