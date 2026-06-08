// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package comment

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/common/apiutils"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/newsrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCommentBackListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询评论列表(后台)
func NewFindCommentBackListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCommentBackListLogic {
	return &FindCommentBackListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindCommentBackListLogic) FindCommentBackList(req *types.QueryCommentReq) (resp *types.PageResp, err error) {
	out, err := l.svcCtx.NewsRpc.FindCommentList(l.ctx, &newsrpc.FindCommentListReq{
		Paginate: &newsrpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		UserId: req.UserId,
		Status: req.Status,
		Type:   req.Type,
	})
	if err != nil {
		return nil, err
	}

	userMap, err := apiutils.BatchQueryMulti(out.List,
		func(item *newsrpc.Comment) []string {
			return []string{item.UserId, item.ReplyUserId}
		},
		func(ids []string) (map[string]*types.UserInfoVO, error) {
			return apiutils.GetUserInfos(l.ctx, l.svcCtx, ids)
		},
	)
	if err != nil {
		return nil, err
	}

	visitorMap, err := apiutils.BatchQuery(out.List,
		func(item *newsrpc.Comment) string {
			return item.TerminalId
		},
		func(ids []string) (map[string]*types.ClientInfoVO, error) {
			return apiutils.GetVisitorInfos(l.ctx, l.svcCtx, ids)
		},
	)
	if err != nil {
		return nil, err
	}

	topicIDs := make([]int64, 0, len(out.List))
	seenTopicIDs := make(map[int64]struct{}, len(out.List))
	for _, item := range out.List {
		if item.TopicId == 0 {
			continue
		}
		if _, ok := seenTopicIDs[item.TopicId]; ok {
			continue
		}
		seenTopicIDs[item.TopicId] = struct{}{}
		topicIDs = append(topicIDs, item.TopicId)
	}

	topicMap := make(map[int64]*articlerpc.ArticlePreview, len(topicIDs))
	if len(topicIDs) > 0 {
		topics, err := l.svcCtx.ArticleRpc.FindArticlePreviewList(l.ctx, &articlerpc.FindArticleListReq{
			Ids: topicIDs,
		})
		if err != nil {
			return nil, err
		}
		for _, item := range topics.List {
			topicMap[item.Id] = item
		}
	}

	list := make([]*types.CommentBackVO, 0, len(out.List))
	for _, item := range out.List {
		comment := &types.CommentBackVO{
			Id:             item.Id,
			UserId:         item.UserId,
			TerminalId:     item.TerminalId,
			Type:           item.Type,
			ReplyUserId:    item.ReplyUserId,
			CommentContent: item.CommentContent,
			Status:         item.Status,
			CreatedAt:      item.CreatedAt,
			ClientInfo:     visitorMap[item.TerminalId],
			UserInfo:       userMap[item.UserId],
			ReplyUserInfo:  userMap[item.ReplyUserId],
		}
		if topic := topicMap[item.TopicId]; topic != nil {
			comment.TopicTitle = topic.ArticleTitle
		}
		list = append(list, comment)
	}

	return &types.PageResp{
		Page:     out.Pagination.Page,
		PageSize: out.Pagination.PageSize,
		Total:    out.Pagination.Total,
		List:     list,
	}, nil
}
