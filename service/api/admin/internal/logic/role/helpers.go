package role

import (
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/permissionrpc"
)

func convertRoleTypes(in *permissionrpc.Role) *types.RoleBackVO {
	if in == nil {
		return nil
	}

	return &types.RoleBackVO{
		Id:          in.Id,
		RoleKey:     in.RoleKey,
		RoleComment: in.RoleComment,
		Status:      in.Status,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
}
