package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserOnlineListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserOnlineListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserOnlineListLogic {
	return &FindUserOnlineListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查找在线用户列表
func (l *FindUserOnlineListLogic) FindUserOnlineList(in *accountrpc.FindUserListReq) (*accountrpc.FindUserInfoListResp, error) {
	page, pageSize, _ := getPageParams(in.Paginate)

	total, err := l.svcCtx.OnlineUserService.GetOnlineUserCount(l.ctx)
	if err != nil {
		return nil, err
	}

	uids, err := l.svcCtx.OnlineUserService.GetOnlineUsers(l.ctx, page, pageSize)
	if err != nil {
		return nil, err
	}

	list := make([]*accountrpc.UserInfo, 0, len(uids))
	for _, uid := range uids {
		user, err := l.svcCtx.TUserModel.FindOneByUserId(l.ctx, uid)
		if err != nil {
			continue
		}

		roles, err := getUserRoles(l.ctx, l.svcCtx, uid)
		if err != nil {
			return nil, err
		}

		list = append(list, convertUserInfoOut(user, roles))
	}

	return &accountrpc.FindUserInfoListResp{
		Pagination: buildPageResp(page, pageSize, total),
		List:       list,
	}, nil
}
