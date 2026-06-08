package talk

import (
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/socialrpc"
)

func convertTalkTypes(in *socialrpc.Talk) *types.TalkBackVO {
	if in == nil {
		return nil
	}

	return &types.TalkBackVO{
		Id:           in.Id,
		UserId:       in.UserId,
		Content:      in.Content,
		ImgList:      in.ImgList,
		IsTop:        in.IsTop,
		Status:       in.Status,
		LikeCount:    in.LikeCount,
		CommentCount: in.CommentCount,
		CreatedAt:    in.CreatedAt,
		UpdatedAt:    in.UpdatedAt,
	}
}
