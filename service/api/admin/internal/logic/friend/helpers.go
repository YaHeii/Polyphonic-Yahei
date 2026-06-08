package friend

import (
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/socialrpc"
)

func convertFriendTypes(out *socialrpc.Friend) *types.FriendBackVO {
	if out == nil {
		return nil
	}

	return &types.FriendBackVO{
		Id:          out.Id,
		LinkName:    out.LinkName,
		LinkAvatar:  out.LinkAvatar,
		LinkAddress: out.LinkAddress,
		LinkIntro:   out.LinkIntro,
		CreatedAt:   out.CreatedAt,
		UpdatedAt:   out.UpdatedAt,
	}
}
