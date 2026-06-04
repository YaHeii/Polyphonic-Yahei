package articlerpclogic

import (
	"context"

	"github.com/lib/pq"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/YaHeii/Polyphonic-Yahei/common/rediskey"
	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/common/query"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"
)

type ArticleHelperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleHelperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleHelperLogic {
	return &ArticleHelperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func convertCategoryIn(in *articlerpc.AddCategoryReq) (out *model.TCategory) {
	out = &model.TCategory{
		Id:           in.Id,
		CategoryName: in.CategoryName,
	}

	return out
}

func convertTagIn(in *articlerpc.AddTagReq) (out *model.TTag) {
	out = &model.TTag{
		Id:      in.Id,
		TagName: in.TagName,
	}

	return out
}

func (l *ArticleHelperLogic) findArticleCountGroupCategory(list []*model.TCategory) (acm map[int64]int, err error) {
	var ids []int64
	for _, v := range list {
		ids = append(ids, v.Id)
	}
	result, err := l.svcCtx.TArticleModel.CountGroupByCategoryIDs(l.ctx, ids)
	if err != nil {
		return nil, err
	}

	acm = make(map[int64]int)
	for categoryID, articleCount := range result {
		acm[categoryID] = int(articleCount)
	}

	return acm, nil
}

// 查询标签下的文章数量
func (l *ArticleHelperLogic) findArticleCountGroupTag(list []*model.TTag) (acm map[int64]int, err error) {
	var names []string
	tagNameIDMap := make(map[string]int64, len(list))
	for _, v := range list {
		names = append(names, v.TagName)
		tagNameIDMap[v.TagName] = v.Id
	}
	result, err := l.svcCtx.TArticleModel.CountGroupByTagNames(l.ctx, names)
	if err != nil {
		return nil, err
	}

	acm = make(map[int64]int)
	for tagName, articleCount := range result {
		if tagID, ok := tagNameIDMap[tagName]; ok {
			acm[tagID] = int(articleCount)
		}
	}

	return acm, nil
}

// 查询文章列表对应的分类
func (l *ArticleHelperLogic) findCategoryGroupArticle(list []*model.TArticle) (acm map[int64]*model.TCategory, err error) {
	categorySet := make(map[int64]struct{})
	for _, v := range list {
		categorySet[v.CategoryId] = struct{}{}
	}

	categoryIDs := make([]int64, 0, len(categorySet))
	for id := range categorySet {
		categoryIDs = append(categoryIDs, id)
	}
	if len(categoryIDs) == 0 {
		return map[int64]*model.TCategory{}, nil
	}

	cs, err := l.svcCtx.TCategoryModel.FindByIds(l.ctx, categoryIDs)
	if err != nil {
		return nil, err
	}

	categoryMap := make(map[int64]*model.TCategory, len(cs))
	for _, category := range cs {
		categoryMap[category.Id] = category
	}

	acm = make(map[int64]*model.TCategory)
	for _, v := range list {
		if category, ok := categoryMap[v.CategoryId]; ok {
			acm[v.Id] = category
		}
	}

	return acm, nil
}

// 查询文章列表对应的标签
func (l *ArticleHelperLogic) findTagGroupArticle(list []*model.TArticle) (atm map[int64][]*model.TTag, err error) {
	tagNameSet := make(map[string]struct{})
	for _, v := range list {
		for _, tagName := range v.Tags {
			if tagName != "" {
				tagNameSet[tagName] = struct{}{}
			}
		}
	}
	if len(tagNameSet) == 0 {
		return map[int64][]*model.TTag{}, nil
	}

	tagNames := make([]string, 0, len(tagNameSet))
	for tagName := range tagNameSet {
		tagNames = append(tagNames, tagName)
	}

	ts, err := l.svcCtx.TTagModel.FindByNames(l.ctx, tagNames)
	if err != nil {
		return nil, err
	}

	tagMap := make(map[string]*model.TTag, len(ts))
	for _, tag := range ts {
		tagMap[tag.TagName] = tag
	}

	atm = make(map[int64][]*model.TTag)
	for _, article := range list {
		for _, tagName := range article.Tags {
			if tag, ok := tagMap[tagName]; ok {
				atm[article.Id] = append(atm[article.Id], tag)
			}
		}
	}

	return atm, nil
}

// 查询或添加文字分类
func (l *ArticleHelperLogic) findOrAddCategory(name string) (int64, error) {
	if name == "" {
		return 0, nil
	}

	category, err := l.svcCtx.TCategoryModel.FindOneByCategoryName(l.ctx, name)
	if err != nil {
		insert := &model.TCategory{
			CategoryName: name,
		}
		_, err := l.svcCtx.TCategoryModel.Insert(l.ctx, insert)
		if err != nil {
			return 0, err
		}
		return insert.Id, nil
	}

	return category.Id, nil
}

// 查询或添加标签
func (l *ArticleHelperLogic) findOrAddTag(name string) (int64, error) {
	if name == "" {
		return 0, nil
	}

	tag, err := l.svcCtx.TTagModel.FindOneByTagName(l.ctx, name)
	if err != nil {
		insert := &model.TTag{
			TagName: name,
		}
		_, err := l.svcCtx.TTagModel.Insert(l.ctx, insert)
		if err != nil {
			return 0, err
		}
		return insert.Id, nil
	}

	return tag.Id, nil
}

func (l *ArticleHelperLogic) convertArticleQuery(in *articlerpc.FindArticleListReq) (page int, size int, sorts string, conditions string, params []any) {
	var opts []query.Option
	if in.Paginate != nil {
		opts = append(opts, query.WithPage(int(in.Paginate.Page)))
		opts = append(opts, query.WithSize(int(in.Paginate.PageSize)))
		opts = append(opts, query.WithSorts(in.Paginate.Sorts...))
	}

	if len(in.Ids) > 0 {
		opts = append(opts, query.WithCondition("id in (?)", in.Ids))
	}

	if in.IsTop >= 0 {
		opts = append(opts, query.WithCondition("is_top = ?", in.IsTop))
	}

	if in.IsDelete >= 0 {
		opts = append(opts, query.WithCondition("is_delete = ?", in.IsDelete))
	}

	if in.Status != 0 {
		opts = append(opts, query.WithCondition("status = ?", in.Status))
	}

	if in.ArticleType != 0 {
		opts = append(opts, query.WithCondition("article_type = ?", in.ArticleType))
	}

	if in.ArticleTitle != "" {
		opts = append(opts, query.WithCondition("article_title like ?", "%"+in.ArticleTitle+"%"))
	}

	if in.CategoryName != "" {
		category, err := l.svcCtx.TCategoryModel.FindOneByCategoryName(l.ctx, in.CategoryName)
		if err == nil {
			opts = append(opts, query.WithCondition("category_id = ?", category.Id))
		}
	}

	if in.TagName != "" {
		opts = append(opts, query.WithCondition("tags @> ?", pq.StringArray{in.TagName}))
	}

	return query.NewQueryBuilder(opts...).Build()
}

func (l *ArticleHelperLogic) convertArticlePreviewOut(record *model.TArticle) (out *articlerpc.ArticlePreview) {
	out = &articlerpc.ArticlePreview{
		Id:           record.Id,
		ArticleCover: record.ArticleCover,
		ArticleTitle: record.ArticleTitle,
		CreatedAt:    record.CreatedAt.UnixMilli(),
		LikeCount:    record.LikeCount,
		ViewCount:    l.GetArticleViewCount(record.Id),
	}
	return out
}

func (l *ArticleHelperLogic) convertArticle(records []*model.TArticle) (out []*articlerpc.ArticleDetails, err error) {
	acm, err := l.findCategoryGroupArticle(records)
	if err != nil {
		return nil, err
	}

	atm, err := l.findTagGroupArticle(records)
	if err != nil {
		return nil, err
	}

	var list []*articlerpc.ArticleDetails
	for _, entity := range records {
		m := &articlerpc.ArticleDetails{
			Id:             entity.Id,
			UserId:         entity.UserId,
			CategoryId:     entity.CategoryId,
			ArticleCover:   entity.ArticleCover,
			ArticleTitle:   entity.ArticleTitle,
			ArticleContent: entity.ArticleContent,
			ArticleType:    entity.ArticleType,
			OriginalUrl:    entity.OriginalUrl,
			IsTop:          boolToInt64(entity.IsTop),
			IsDelete:       boolToInt64(entity.IsDelete),
			Status:         entity.Status,
			CreatedAt:      entity.CreatedAt.UnixMilli(),
			UpdatedAt:      entity.UpdatedAt.UnixMilli(),
			LikeCount:      entity.LikeCount,
			ViewCount:      l.GetArticleViewCount(entity.Id),
			Category:       nil,
			TagList:        nil,
		}

		if v, ok := acm[entity.Id]; ok {
			m.Category = &articlerpc.ArticleCategory{
				Id:           v.Id,
				CategoryName: v.CategoryName,
			}
		}

		if v, ok := atm[entity.Id]; ok {
			tagList := make([]*articlerpc.ArticleTag, 0, len(v))
			for _, tag := range v {
				tagList = append(tagList, &articlerpc.ArticleTag{
					Id:      tag.Id,
					TagName: tag.TagName,
				})
			}
			m.TagList = tagList
		}

		list = append(list, m)
	}

	return list, nil
}

func (l *ArticleHelperLogic) convertCategory(records []*model.TCategory) (out []*articlerpc.CategoryDetails, err error) {
	acm, err := l.findArticleCountGroupCategory(records)
	if err != nil {
		return nil, err
	}

	var list []*articlerpc.CategoryDetails
	for _, entity := range records {

		m := &articlerpc.CategoryDetails{
			Id:           entity.Id,
			CategoryName: entity.CategoryName,
			ArticleCount: 0,
			CreatedAt:    entity.CreatedAt.UnixMilli(),
			UpdatedAt:    entity.UpdatedAt.UnixMilli(),
		}

		if v, ok := acm[entity.Id]; ok {
			m.ArticleCount = int64(v)
		}

		list = append(list, m)
	}

	return list, nil
}

func (l *ArticleHelperLogic) convertTag(records []*model.TTag) (out []*articlerpc.TagDetails, err error) {
	acm, err := l.findArticleCountGroupTag(records)
	if err != nil {
		return nil, err
	}

	var list []*articlerpc.TagDetails
	for _, entity := range records {
		m := &articlerpc.TagDetails{
			Id:           entity.Id,
			TagName:      entity.TagName,
			ArticleCount: 0,
			CreatedAt:    entity.CreatedAt.UnixMilli(),
			UpdatedAt:    entity.UpdatedAt.UnixMilli(),
		}

		if v, ok := acm[entity.Id]; ok {
			m.ArticleCount = int64(v)
		}

		list = append(list, m)
	}

	return list, nil
}

func (l *ArticleHelperLogic) GetArticleViewCount(articleId int64) (count int64) {
	id := cast.ToString(articleId)
	key := rediskey.GetArticleViewCountKey()
	result, err := l.svcCtx.Redis.ZScore(l.ctx, key, id).Result()
	if err != nil {
		// 未找到，从数据库查找，并且设置
		article, err := l.svcCtx.TArticleModel.FindById(l.ctx, articleId)
		if err != nil {
			return 0
		}

		_ = l.svcCtx.Redis.ZIncrBy(l.ctx, key, float64(article.ViewCount), id).Err()
		return article.ViewCount
	}

	return int64(result)
}

// 获取浏览人数最高的文章列表
func (l *ArticleHelperLogic) GetViewTopArticleList(count int64) (list []*model.TArticle, err error) {
	key := rediskey.GetArticleViewCountKey()
	ids, err := l.svcCtx.Redis.ZRevRange(l.ctx, key, 0, count).Result()
	if err != nil {
		return nil, err
	}

	var idList []int64
	for _, v := range ids {
		idList = append(idList, cast.ToInt64(v))
	}
	if len(idList) == 0 {
		return []*model.TArticle{}, nil
	}

	list, err = l.svcCtx.TArticleModel.FindByIds(l.ctx, idList)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// 获取每日文章生产数量
func (l *ArticleHelperLogic) GetArticleDailyStatistics() (out map[string]int64, err error) {
	return l.svcCtx.TArticleModel.GetDailyStatistics(l.ctx)
}

func boolToInt64(v bool) int64 {
	if v {
		return 1
	}
	return 0
}
