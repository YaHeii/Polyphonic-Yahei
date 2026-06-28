// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package login_log

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/common/apiutils"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/syslogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindLoginLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询登录日志
func NewFindLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindLoginLogListLogic {
	return &FindLoginLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindLoginLogListLogic) FindLoginLogList(req *types.QueryLoginLogReq) (resp *types.LoginLogPageResp, err error) {
	in := &syslogrpc.FindLoginLogListReq{
		Paginate: &syslogrpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		UserId: req.UserId,
	}

	out, err := l.svcCtx.SyslogRpc.FindLoginLogList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	userMap, err := apiutils.BatchQuery(out.List,
		func(item *syslogrpc.LoginLog) string {
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
		func(item *syslogrpc.LoginLog) string {
			return item.TerminalId
		},
		func(ids []string) (map[string]*types.ClientInfoVO, error) {
			return apiutils.GetVisitorInfos(l.ctx, l.svcCtx, ids)
		},
	)
	if err != nil {
		return nil, err
	}

	list := make([]*types.LoginLogBackVO, 0, len(out.List))
	for _, item := range out.List {
		list = append(list, &types.LoginLogBackVO{
			Id:         item.Id,
			UserId:     item.UserId,
			TerminalId: item.TerminalId,
			LoginType:  item.LoginType,
			AppName:    item.AppName,
			LoginAt:    item.LoginAt,
			LogoutAt:   item.LogoutAt,
			UserInfo:   userMap[item.UserId],
			ClientInfo: visitorMap[item.TerminalId],
		})
	}

	return &types.LoginLogPageResp{
		PageMeta: types.PageMeta{
			Page:     out.Pagination.Page,
			PageSize: out.Pagination.PageSize,
			Total:    out.Pagination.Total,
		},
		List: list,
	}, nil
}
