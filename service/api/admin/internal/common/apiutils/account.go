package apiutils

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/jsonconv"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
)

func GetUserInfos(ctx context.Context, svcCtx *svc.ServiceContext, userIDs []string) (map[string]*types.UserInfoVO, error) {
	if len(userIDs) == 0 {
		return map[string]*types.UserInfoVO{}, nil
	}

	out, err := svcCtx.AccountRpc.FindUserList(ctx, &accountrpc.FindUserListReq{
		Paginate: &accountrpc.PageReq{
			Page:     1,
			PageSize: int64(len(userIDs)),
		},
		UserIds: userIDs,
	})
	if err != nil {
		return nil, err
	}

	result := make(map[string]*types.UserInfoVO, len(out.List))
	for _, user := range out.List {
		var ext types.UserInfoExt
		if user.Info != "" {
			if err := jsonconv.JsonToAny(user.Info, &ext); err != nil {
				return nil, err
			}
		}

		result[user.UserId] = &types.UserInfoVO{
			UserId:      user.UserId,
			Username:    user.Username,
			Avatar:      user.Avatar,
			Nickname:    user.Nickname,
			UserInfoExt: ext,
		}
	}

	return result, nil
}

func GetVisitorInfos(ctx context.Context, svcCtx *svc.ServiceContext, terminalIDs []string) (map[string]*types.ClientInfoVO, error) {
	if len(terminalIDs) == 0 {
		return map[string]*types.ClientInfoVO{}, nil
	}

	out, err := svcCtx.AccountRpc.FindVisitorList(ctx, &accountrpc.FindVisitorListReq{
		Paginate: &accountrpc.PageReq{
			Page:     1,
			PageSize: int64(len(terminalIDs)),
		},
		TerminalIds: terminalIDs,
	})
	if err != nil {
		return nil, err
	}

	result := make(map[string]*types.ClientInfoVO, len(out.List))
	for _, visitor := range out.List {
		result[visitor.TerminalId] = &types.ClientInfoVO{
			TerminalId: visitor.TerminalId,
			Os:         visitor.Os,
			Browser:    visitor.Browser,
			IpAddress:  visitor.IpAddress,
			IpSource:   visitor.IpSource,
		}
	}

	return result, nil
}
