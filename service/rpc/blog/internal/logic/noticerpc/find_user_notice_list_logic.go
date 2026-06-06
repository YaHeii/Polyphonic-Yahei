package noticerpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/common/query"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/noticerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserNoticeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserNoticeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserNoticeListLogic {
	return &FindUserNoticeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户可见通知列表
func (l *FindUserNoticeListLogic) FindUserNoticeList(in *noticerpc.FindUserNoticeListReq) (*noticerpc.FindUserNoticeListResp, error) {
	var opts []query.Option
	if in.Paginate != nil {
		opts = append(opts, query.WithPage(int(in.Paginate.Page)))
		opts = append(opts, query.WithSize(int(in.Paginate.PageSize)))
		opts = append(opts, query.WithSorts(in.Paginate.Sorts...))
	}
	opts = append(opts, query.WithCondition("publish_status = ?", noticeStatusPublished))

	page, size, sorts, conditions, params := query.NewQueryBuilder(opts...).Build()
	records, total, err := l.svcCtx.TSystemNoticeModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	return &noticerpc.FindUserNoticeListResp{
		List: convertNoticeListOut(records),
		Pagination: &noticerpc.PageResp{
			Page:     int64(page),
			PageSize: int64(size),
			Total:    total,
		},
	}, nil
}
