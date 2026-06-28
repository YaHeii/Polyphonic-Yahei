// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package operation_log

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/common/apiutils"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/syslogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOperationLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取操作记录列表
func NewFindOperationLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOperationLogListLogic {
	return &FindOperationLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindOperationLogListLogic) FindOperationLogList(req *types.QueryOperationLogReq) (resp *types.OperationLogPageResp, err error) {
	in := &syslogrpc.FindOperationLogListReq{
		Paginate: &syslogrpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
	}

	out, err := l.svcCtx.SyslogRpc.FindOperationLogList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	userMap, err := apiutils.BatchQuery(out.List,
		func(item *syslogrpc.OperationLog) string {
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
		func(item *syslogrpc.OperationLog) string {
			return item.TerminalId
		},
		func(ids []string) (map[string]*types.ClientInfoVO, error) {
			return apiutils.GetVisitorInfos(l.ctx, l.svcCtx, ids)
		},
	)
	if err != nil {
		return nil, err
	}

	list := make([]*types.OperationLogBackVO, 0, len(out.List))
	for _, item := range out.List {
		list = append(list, &types.OperationLogBackVO{
			Id:             item.Id,
			UserId:         item.UserId,
			TerminalId:     item.TerminalId,
			OptModule:      item.OptModule,
			OptDesc:        item.OptDesc,
			RequestUri:     item.RequestUri,
			RequestMethod:  item.RequestMethod,
			RequestData:    item.RequestData,
			ResponseData:   item.ResponseData,
			ResponseStatus: item.ResponseStatus,
			Cost:           item.Cost,
			CreatedAt:      item.CreatedAt,
			UpdatedAt:      item.UpdatedAt,
			UserInfo:       userMap[item.UserId],
			ClientInfo:     visitorMap[item.TerminalId],
		})
	}

	return &types.OperationLogPageResp{
		PageMeta: types.PageMeta{
			Page:     out.Pagination.Page,
			PageSize: out.Pagination.PageSize,
			Total:    out.Pagination.Total,
		},
		List: list,
	}, nil
}
