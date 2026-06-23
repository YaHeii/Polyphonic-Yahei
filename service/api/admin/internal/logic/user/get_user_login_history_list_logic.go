// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/authctx"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/syslogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLoginHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询用户登录历史
func NewGetUserLoginHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLoginHistoryListLogic {
	return &GetUserLoginHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLoginHistoryListLogic) GetUserLoginHistoryList(req *types.QueryUserLoginHistoryReq) (resp *types.PageResp, err error) {
	out, err := l.svcCtx.SyslogRpc.FindLoginLogList(l.ctx, &syslogrpc.FindLoginLogListReq{
		Paginate: &syslogrpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		UserId: authctx.CurrentUserID(l.ctx),
	})
	if err != nil {
		return nil, err
	}

	list := make([]*types.UserLoginHistory, 0, len(out.List))
	for _, item := range out.List {
		list = append(list, &types.UserLoginHistory{
			Id:         item.Id,
			UserId:     item.UserId,
			TerminalId: item.TerminalId,
			LoginType:  item.LoginType,
			AppName:    item.AppName,
			LoginAt:    item.LoginAt,
			LogoutAt:   item.LogoutAt,
		})
	}

	return &types.PageResp{
		Page:     out.Pagination.Page,
		PageSize: out.Pagination.PageSize,
		Total:    out.Pagination.Total,
		List:     list,
	}, nil
}
