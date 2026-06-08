// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package talk

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizheader"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/socialrpc"
	"github.com/spf13/cast"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTalkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建说说
func NewAddTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTalkLogic {
	return &AddTalkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTalkLogic) AddTalk(req *types.NewTalkReq) (resp *types.TalkBackVO, err error) {
	in := &socialrpc.AddTalkReq{
		Id:      req.Id,
		UserId:  cast.ToString(l.ctx.Value(bizheader.HeaderUid)),
		Content: req.Content,
		ImgList: req.ImgList,
		IsTop:   req.IsTop,
		Status:  req.Status,
	}

	out, err := l.svcCtx.SocialRpc.AddTalk(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convertTalkTypes(out.Talk), nil
}
