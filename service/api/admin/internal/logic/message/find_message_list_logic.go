// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package message

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/common/apiutils"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/newsrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMessageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取留言列表
func NewFindMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMessageListLogic {
	return &FindMessageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindMessageListLogic) FindMessageList(req *types.QueryMessageReq) (resp *types.MessagePageResp, err error) {
	in := &newsrpc.FindMessageListReq{
		Paginate: &newsrpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		UserId: req.UserId,
		Status: req.Status,
	}

	out, err := l.svcCtx.NewsRpc.FindMessageList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	userMap, err := apiutils.BatchQuery(out.List,
		func(item *newsrpc.Message) string {
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
		func(item *newsrpc.Message) string {
			return item.TerminalId
		},
		func(ids []string) (map[string]*types.ClientInfoVO, error) {
			return apiutils.GetVisitorInfos(l.ctx, l.svcCtx, ids)
		},
	)
	if err != nil {
		return nil, err
	}

	list := make([]*types.MessageBackVO, 0, len(out.List))
	for _, item := range out.List {
		list = append(list, &types.MessageBackVO{
			Id:             item.Id,
			UserId:         item.UserId,
			TerminalId:     item.TerminalId,
			MessageContent: item.MessageContent,
			Status:         item.Status,
			CreatedAt:      item.CreatedAt,
			UpdatedAt:      item.UpdatedAt,
			UserInfo:       userMap[item.UserId],
			ClientInfo:     visitorMap[item.TerminalId],
		})
	}

	return &types.MessagePageResp{
		PageMeta: types.PageMeta{
			Page:     out.Pagination.Page,
			PageSize: out.Pagination.PageSize,
			Total:    out.Pagination.Total,
		},
		List: list,
	}, nil
}
