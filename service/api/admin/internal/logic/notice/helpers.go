package notice

import (
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/noticerpc"
)

func convertNoticeOut(in *noticerpc.Notice) *types.NoticeBackVO {
	if in == nil {
		return nil
	}

	return &types.NoticeBackVO{
		Id:            in.Id,
		Title:         in.Title,
		Content:       in.Content,
		Type:          in.Type,
		Level:         in.Level,
		AppName:       in.AppName,
		PublisherId:   in.PublisherId,
		PublishStatus: in.PublishStatus,
		PublishTime:   in.PublishTime,
		RevokeTime:    in.RevokeTime,
		CreatedAt:     in.CreatedAt,
		UpdatedAt:     in.UpdatedAt,
	}
}
