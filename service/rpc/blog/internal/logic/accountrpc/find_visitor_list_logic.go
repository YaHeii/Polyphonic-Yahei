package accountrpclogic

import (
	"context"
	"strconv"

	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindVisitorListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindVisitorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindVisitorListLogic {
	return &FindVisitorListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询游客信息
func (l *FindVisitorListLogic) FindVisitorList(in *accountrpc.FindVisitorListReq) (*accountrpc.FindVisitorListResp, error) {
	page, pageSize, offset := getPageParams(in.Paginate)

	conditions := make([]string, 0, 4)
	args := make([]any, 0, 4)
	conditions, args = appendStringEqualCondition(conditions, args, "terminal_id", in.TerminalId)
	conditions, args = appendStringInCondition(conditions, args, "terminal_id", in.TerminalIds)
	conditions, args = appendStringLikeCondition(conditions, args, "ip_source", in.IpSource)

	whereClause := joinWhereClause(conditions)

	var totalRow struct {
		Count int64 `db:"count"`
	}
	if err := l.svcCtx.SqlConn.QueryRowCtx(l.ctx, &totalRow, `select count(*) as count from "public"."t_visitor"`+whereClause, args...); err != nil {
		return nil, err
	}

	queryArgs := append(append([]any{}, args...), pageSize, offset)
	query := `select id, terminal_id, os, browser, ip_address, ip_source, created_at, updated_at
from "public"."t_visitor"` + whereClause + `
order by id desc
limit $` + strconv.Itoa(len(args)+1) + ` offset $` + strconv.Itoa(len(args)+2)

	var visitors []*model.TVisitor
	if err := l.svcCtx.SqlConn.QueryRowsCtx(l.ctx, &visitors, query, queryArgs...); err != nil {
		return nil, err
	}

	list := make([]*accountrpc.VisitorInfo, 0, len(visitors))
	for _, visitor := range visitors {
		list = append(list, convertVisitorInfoOut(visitor))
	}

	return &accountrpc.FindVisitorListResp{
		Pagination: buildPageResp(page, pageSize, totalRow.Count),
		List:       list,
	}, nil
}
