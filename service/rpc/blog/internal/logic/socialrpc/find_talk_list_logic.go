package socialrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/socialrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTalkListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindTalkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTalkListLogic {
	return &FindTalkListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询说说列表
func (l *FindTalkListLogic) FindTalkList(in *socialrpc.FindTalkListReq) (*socialrpc.FindTalkListResp, error) {
	page, size, sorts, conditions, params := buildTalkQuery(in)
	records, total, err := l.svcCtx.TTalkModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	commentCounts, err := l.svcCtx.TCommentModel.CountGroupByTopicIDs(l.ctx, collectTalkIDs(records), talkCommentType())
	if err != nil {
		return nil, err
	}

	return &socialrpc.FindTalkListResp{
		Pagination: buildPageResp(page, size, total),
		List:       convertTalkListOut(records, commentCounts),
	}, nil
}
