// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package talk

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/common/apiutils"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/socialrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTalkListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取说说列表
func NewFindTalkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTalkListLogic {
	return &FindTalkListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindTalkListLogic) FindTalkList(req *types.QueryTalkReq) (resp *types.PageResp, err error) {
	in := &socialrpc.FindTalkListReq{
		Paginate: &socialrpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		Status: req.Status,
	}

	out, err := l.svcCtx.SocialRpc.FindTalkList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	userMap, err := apiutils.BatchQuery(out.List,
		func(item *socialrpc.Talk) string {
			return item.UserId
		},
		func(ids []string) (map[string]*types.UserInfoVO, error) {
			return apiutils.GetUserInfos(l.ctx, l.svcCtx, ids)
		},
	)
	if err != nil {
		return nil, err
	}

	list := make([]*types.TalkBackVO, 0, len(out.List))
	for _, item := range out.List {
		vo := convertTalkTypes(item)
		vo.UserInfo = userMap[item.UserId]
		list = append(list, vo)
	}

	return &types.PageResp{
		Page:     out.Pagination.Page,
		PageSize: out.Pagination.PageSize,
		Total:    out.Pagination.Total,
		List:     list,
	}, nil
}
