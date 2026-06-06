package socialrpclogic

import (
	"context"

	"github.com/spf13/cast"

	"github.com/YaHeii/Polyphonic-Yahei/common/rediskey"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/socialrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLikeTalkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLikeTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLikeTalkLogic {
	return &FindUserLikeTalkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户点赞的说说
func (l *FindUserLikeTalkLogic) FindUserLikeTalk(in *socialrpc.FindUserLikeTalkReq) (*socialrpc.FindUserLikeTalkResp, error) {
	values, err := l.svcCtx.Redis.SMembers(l.ctx, rediskey.GetUserLikeTalkKey(in.UserId)).Result()
	if err != nil {
		return nil, err
	}

	ids := make([]int64, 0, len(values))
	for _, value := range values {
		if id := cast.ToInt64(value); id != 0 {
			ids = append(ids, id)
		}
	}

	return &socialrpc.FindUserLikeTalkResp{LikeTalkList: ids}, nil
}
