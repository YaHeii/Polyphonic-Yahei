// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package visit_log

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/common/apiutils"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/syslogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindVisitLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取操作记录列表
func NewFindVisitLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindVisitLogListLogic {
	return &FindVisitLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindVisitLogListLogic) FindVisitLogList(req *types.QueryVisitLogReq) (resp *types.PageResp, err error) {
	in := &syslogrpc.FindVisitLogListReq{
		Paginate: &syslogrpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		UserId:     req.UserId,
		TerminalId: req.TerminalId,
		PageName:   req.PageName,
	}

	out, err := l.svcCtx.SyslogRpc.FindVisitLogList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	userMap, err := apiutils.BatchQuery(out.List,
		func(item *syslogrpc.VisitLog) string {
			return item.UserId
		},
		func(ids []string) (map[string]*types.UserInfoVO, error) {
			return apiutils.GetUserInfos(l.ctx, l.svcCtx, ids)
		},
	)
	if err != nil {
		return nil, err
	}

	visitorMap, err := apiutils.BatchQuery(out.List,
		func(item *syslogrpc.VisitLog) string {
			return item.TerminalId
		},
		func(ids []string) (map[string]*types.ClientInfoVO, error) {
			return apiutils.GetVisitorInfos(l.ctx, l.svcCtx, ids)
		},
	)
	if err != nil {
		return nil, err
	}

	list := make([]*types.VisitLogBackVO, 0, len(out.List))
	for _, item := range out.List {
		list = append(list, &types.VisitLogBackVO{
			Id:         item.Id,
			UserId:     item.UserId,
			TerminalId: item.TerminalId,
			PageName:   item.PageName,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
			UserInfo:   userMap[item.UserId],
			ClientInfo: visitorMap[item.TerminalId],
		})
	}

	return &types.PageResp{
		Page:     out.Pagination.Page,
		PageSize: out.Pagination.PageSize,
		Total:    out.Pagination.Total,
		List:     list,
	}, nil
}
