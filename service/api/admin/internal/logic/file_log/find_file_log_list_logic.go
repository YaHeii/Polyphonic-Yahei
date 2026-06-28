// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package file_log

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/common/apiutils"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/syslogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindFileLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询文件日志
func NewFindFileLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindFileLogListLogic {
	return &FindFileLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindFileLogListLogic) FindFileLogList(req *types.QueryFileLogReq) (resp *types.FileLogPageResp, err error) {
	in := &syslogrpc.FindFileLogListReq{
		Paginate: &syslogrpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		FilePath: req.FilePath,
		FileName: req.FileName,
		FileType: req.FileType,
	}

	out, err := l.svcCtx.SyslogRpc.FindFileLogList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	userMap, err := apiutils.BatchQuery(out.List,
		func(item *syslogrpc.FileLog) string {
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
		func(item *syslogrpc.FileLog) string {
			return item.TerminalId
		},
		func(ids []string) (map[string]*types.ClientInfoVO, error) {
			return apiutils.GetVisitorInfos(l.ctx, l.svcCtx, ids)
		},
	)
	if err != nil {
		return nil, err
	}

	list := make([]*types.FileLogBackVO, 0, len(out.List))
	for _, item := range out.List {
		list = append(list, &types.FileLogBackVO{
			Id:         item.Id,
			UserId:     item.UserId,
			TerminalId: item.TerminalId,
			FilePath:   item.FilePath,
			FileName:   item.FileName,
			FileType:   item.FileType,
			FileSize:   item.FileSize,
			FileMd5:    item.FileMd5,
			FileUrl:    item.FileUrl,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
			UserInfo:   userMap[item.UserId],
			ClientInfo: visitorMap[item.TerminalId],
		})
	}

	return &types.FileLogPageResp{
		PageMeta: types.PageMeta{
			Page:     out.Pagination.Page,
			PageSize: out.Pagination.PageSize,
			Total:    out.Pagination.Total,
		},
		List: list,
	}, nil
}
