package account

import (
	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/jsonconv"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
)

func convertUserInfoTypes(in *accountrpc.UserInfo) *types.UserInfoDetail {

	var info types.UserInfoExt
	jsonconv.JsonToAny(in.Info, &info)

	roles := make([]*types.UserRoleLabel, 0)
	for _, v := range in.Roles {
		m := &types.UserRoleLabel{
			RoleId:    v.RoleId,
			RoleKey:   v.RoleKey,
			RoleLabel: v.RoleLabel,
		}

		roles = append(roles, m)
	}

	out := &types.UserInfoDetail{
		UserId:       in.UserId,
		Username:     in.Username,
		Nickname:     in.Nickname,
		Avatar:       in.Avatar,
		Email:        in.Email,
		Phone:        in.Phone,
		Status:       in.Status,
		RegisterType: in.RegisterType,
		IpAddress:    in.IpAddress,
		IpSource:     in.IpSource,
		CreatedAt:    in.CreatedAt,
		UpdatedAt:    in.UpdatedAt,
		UserInfoExt:  info,
		RoleLabels:   roles,
	}

	return out
}
