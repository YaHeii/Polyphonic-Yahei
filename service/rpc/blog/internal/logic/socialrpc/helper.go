package socialrpclogic

import (
	"github.com/YaHeii/Polyphonic-Yahei/common/enums"
	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/common/query"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/socialrpc"
)

func convertAddFriendIn(in *socialrpc.AddFriendReq) *model.TFriend {
	return &model.TFriend{
		Id:          in.Id,
		LinkName:    in.LinkName,
		LinkAvatar:  in.LinkAvatar,
		LinkAddress: in.LinkAddress,
		LinkIntro:   in.LinkIntro,
	}
}

func convertUpdateFriendIn(in *socialrpc.UpdateFriendReq) *model.TFriend {
	return &model.TFriend{
		Id:          in.Id,
		LinkName:    in.LinkName,
		LinkAvatar:  in.LinkAvatar,
		LinkAddress: in.LinkAddress,
		LinkIntro:   in.LinkIntro,
	}
}

func convertFriendOut(record *model.TFriend) *socialrpc.Friend {
	if record == nil {
		return nil
	}

	return &socialrpc.Friend{
		Id:          record.Id,
		LinkName:    record.LinkName,
		LinkAvatar:  record.LinkAvatar,
		LinkAddress: record.LinkAddress,
		LinkIntro:   record.LinkIntro,
		CreatedAt:   record.CreatedAt.UnixMilli(),
		UpdatedAt:   record.UpdatedAt.UnixMilli(),
	}
}

func convertFriendListOut(records []*model.TFriend) []*socialrpc.Friend {
	list := make([]*socialrpc.Friend, 0, len(records))
	for _, record := range records {
		list = append(list, convertFriendOut(record))
	}
	return list
}

func convertAddTalkIn(in *socialrpc.AddTalkReq) *model.TTalk {
	return &model.TTalk{
		Id:      in.Id,
		UserId:  in.UserId,
		Content: in.Content,
		Images:  append([]string(nil), in.ImgList...),
		IsTop:   in.IsTop,
		Status:  in.Status,
	}
}

func convertUpdateTalkIn(in *socialrpc.UpdateTalkReq) *model.TTalk {
	return &model.TTalk{
		Id:      in.Id,
		UserId:  in.UserId,
		Content: in.Content,
		Images:  append([]string(nil), in.ImgList...),
		IsTop:   in.IsTop,
		Status:  in.Status,
	}
}

func convertTalkOut(record *model.TTalk, commentCount int64) *socialrpc.Talk {
	if record == nil {
		return nil
	}

	return &socialrpc.Talk{
		Id:           record.Id,
		UserId:       record.UserId,
		Content:      record.Content,
		ImgList:      []string(record.Images),
		IsTop:        record.IsTop,
		Status:       record.Status,
		CreatedAt:    record.CreatedAt.UnixMilli(),
		UpdatedAt:    record.UpdatedAt.UnixMilli(),
		LikeCount:    record.LikeCount,
		CommentCount: commentCount,
	}
}

func convertTalkListOut(records []*model.TTalk, commentCounts map[int64]int64) []*socialrpc.Talk {
	list := make([]*socialrpc.Talk, 0, len(records))
	for _, record := range records {
		list = append(list, convertTalkOut(record, commentCounts[record.Id]))
	}
	return list
}

func buildFriendQuery(in *socialrpc.FindFriendListReq) (int, int, string, string, []any) {
	var opts []query.Option
	if in.Paginate != nil {
		opts = append(opts, query.WithPage(int(in.Paginate.Page)))
		opts = append(opts, query.WithSize(int(in.Paginate.PageSize)))
		opts = append(opts, query.WithSorts(in.Paginate.Sorts...))
	}
	if in.LinkName != "" {
		opts = append(opts, query.WithCondition("link_name like ?", "%"+in.LinkName+"%"))
	}
	return query.NewQueryBuilder(opts...).Build()
}

func buildTalkQuery(in *socialrpc.FindTalkListReq) (int, int, string, string, []any) {
	var opts []query.Option
	if in.Paginate != nil {
		opts = append(opts, query.WithPage(int(in.Paginate.Page)))
		opts = append(opts, query.WithSize(int(in.Paginate.PageSize)))
		opts = append(opts, query.WithSorts(in.Paginate.Sorts...))
	}
	if in.Status != 0 {
		opts = append(opts, query.WithCondition("status = ?", in.Status))
	}
	return query.NewQueryBuilder(opts...).Build()
}

func buildPageResp(page int, size int, total int64) *socialrpc.PageResp {
	return &socialrpc.PageResp{
		Page:     int64(page),
		PageSize: int64(size),
		Total:    total,
	}
}

func collectTalkIDs(records []*model.TTalk) []int64 {
	ids := make([]int64, 0, len(records))
	for _, record := range records {
		ids = append(ids, record.Id)
	}
	return ids
}

func talkCommentType() int64 {
	return int64(enums.CommentTypeTalk)
}
