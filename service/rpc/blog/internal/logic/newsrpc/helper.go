package newsrpclogic

import (
	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/newsrpc"
)

func convertCommentOut(in *model.TComment) (out *newsrpc.Comment) {
	out = &newsrpc.Comment{
		Id:             in.Id,
		UserId:         in.UserId,
		TopicId:        in.TopicId,
		ParentId:       in.ParentId,
		ReplyId:        in.ReplyId,
		ReplyUserId:    in.ReplyUserId,
		CommentContent: in.CommentContent,
		Type:           in.Type,
		Status:         in.Status,
		CreatedAt:      in.CreatedAt.UnixMilli(),
		UpdatedAt:      in.UpdatedAt.UnixMilli(),
		LikeCount:      in.LikeCount,
	}

	return out
}

func convertMessageOut(in *model.TMessage) (out *newsrpc.Message) {
	out = &newsrpc.Message{
		Id:             in.Id,
		UserId:         in.UserId,
		TerminalId:     in.TerminalId,
		MessageContent: in.MessageContent,
		Status:         in.Status,
		CreatedAt:      in.CreatedAt.UnixMilli(),
		UpdatedAt:      in.UpdatedAt.UnixMilli(),
	}

	return out
}

