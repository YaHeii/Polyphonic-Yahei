package articlerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesTagLogic {
	return &DeletesTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除标签
func (l *DeletesTagLogic) DeletesTag(in *articlerpc.DeletesTagReq) (*articlerpc.DeletesTagResp, error) {
	rows, err := l.svcCtx.TTagModel.Deletes(l.ctx, "id = any(?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &articlerpc.DeletesTagResp{
		SuccessCount: rows,
	}, nil
}
