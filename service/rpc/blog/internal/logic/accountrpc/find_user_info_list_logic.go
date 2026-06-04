package accountrpclogic

import (
	"context"
	"strconv"

	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserInfoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserInfoListLogic {
	return &FindUserInfoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查找用户信息列表
func (l *FindUserInfoListLogic) FindUserInfoList(in *accountrpc.FindUserListReq) (*accountrpc.FindUserInfoListResp, error) {
	page, pageSize, offset := getPageParams(in.Paginate)

	conditions := make([]string, 0, 6)
	args := make([]any, 0, 6)
	conditions, args = appendStringLikeCondition(conditions, args, "username", in.Username)
	conditions, args = appendStringLikeCondition(conditions, args, "nickname", in.Nickname)
	conditions, args = appendStringLikeCondition(conditions, args, "email", in.Email)
	conditions, args = appendStringLikeCondition(conditions, args, "phone", in.Phone)
	if in.Status != 0 {
		conditions, args = appendInt64EqualCondition(conditions, args, "status", in.Status)
	}
	conditions, args = appendStringInCondition(conditions, args, "user_id", in.UserIds)

	whereClause := joinWhereClause(conditions)

	var totalRow struct {
		Count int64 `db:"count"`
	}
	if err := l.svcCtx.SqlConn.QueryRowCtx(l.ctx, &totalRow, `select count(*) as count from "public"."t_user"`+whereClause, args...); err != nil {
		return nil, err
	}

	queryArgs := append(append([]any{}, args...), pageSize, offset)
	query := `select id, user_id, username, password, nickname, avatar, email, phone, info, status, register_type, ip_address, ip_source, created_at, updated_at
from "public"."t_user"` + whereClause + `
order by id desc
limit $` + strconv.Itoa(len(args)+1) + ` offset $` + strconv.Itoa(len(args)+2)

	var users []*model.TUser
	if err := l.svcCtx.SqlConn.QueryRowsCtx(l.ctx, &users, query, queryArgs...); err != nil {
		return nil, err
	}

	list := make([]*accountrpc.UserInfo, 0, len(users))
	for _, user := range users {
		roles, err := getUserRoles(l.ctx, l.svcCtx, user.UserId)
		if err != nil {
			return nil, err
		}

		list = append(list, convertUserInfoOut(user, roles))
	}

	return &accountrpc.FindUserInfoListResp{
		Pagination: buildPageResp(page, pageSize, totalRow.Count),
		List:       list,
	}, nil
}
