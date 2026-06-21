package articlerpclogic

import (
	"context"
	"sort"

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
func (l *ArticleHelperLogic) findTagGroupArticle(list []*model.TArticle) map[int64][]string {
	atm := make(map[int64][]string, len(list))
	for _, article := range list {
		if len(article.Tags) == 0 {
			continue
		}

		tagList := make([]string, 0, len(article.Tags))
		for _, tagName := range article.Tags {
			if tagName == "" {
				continue
			}
			tagList = append(tagList, tagName)
		}
		if len(tagList) > 0 {
			atm[article.Id] = tagList
		}
	}

	return atm
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

func (l *ArticleHelperLogic) ensureTags(names []string) ([]string, error) {
	if len(names) == 0 {
		return []string{}, nil
	}

	result := make([]string, 0, len(names))
	seen := make(map[string]struct{}, len(names))
	for _, name := range names {
		if name == "" {
			continue
		}
		if _, exists := seen[name]; exists {
			continue
		}
		seen[name] = struct{}{}
		result = append(result, name)
	}

	return result, nil
}

func (l *ArticleHelperLogic) convertArticleQuery(in *articlerpc.FindArticleListReq) (page int, size int, sorts string, conditions string, params []any) {
	var opts []query.Option
	if in.Paginate != nil {
		opts = append(opts, query.WithPage(int(in.Paginate.Page)))
		opts = append(opts, query.WithSize(int(in.Paginate.PageSize)))
		opts = append(opts, query.WithSorts(in.Paginate.Sorts...))
	}

	if len(in.Ids) > 0 {
		opts = append(opts, query.WithCondition("id = any(?)", in.Ids))
	}

	if in.IsTop != nil {
		opts = append(opts, query.WithCondition("is_top = ?", *in.IsTop))
	}

	if in.IsDelete != nil {
		opts = append(opts, query.WithCondition("is_delete = ?", *in.IsDelete))
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
		opts = append(opts, query.WithCondition("tags @> ?", []string{in.TagName}))
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

	atm := l.findTagGroupArticle(records)

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
			IsTop:          entity.IsTop,
			IsDelete:       entity.IsDelete,
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
			for _, tagName := range v {
				tagList = append(tagList, &articlerpc.ArticleTag{
					TagName: tagName,
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

func (l *ArticleHelperLogic) FindAllTags() ([]*articlerpc.TagDetails, error) {
	articles, err := l.svcCtx.TArticleModel.FindALL(l.ctx, "")
	if err != nil {
		return nil, err
	}

	tagCounts := make(map[string]int64)
	for _, article := range articles {
		seen := make(map[string]struct{}, len(article.Tags))
		for _, tagName := range article.Tags {
			if tagName == "" {
				continue
			}
			if _, exists := seen[tagName]; exists {
				continue
			}
			seen[tagName] = struct{}{}
			tagCounts[tagName]++
		}
	}

	if len(tagCounts) == 0 {
		return []*articlerpc.TagDetails{}, nil
	}

	tagNames := make([]string, 0, len(tagCounts))
	for tagName := range tagCounts {
		tagNames = append(tagNames, tagName)
	}
	sort.Strings(tagNames)

	list := make([]*articlerpc.TagDetails, 0, len(tagNames))
	for _, tagName := range tagNames {
		list = append(list, &articlerpc.TagDetails{
			TagName:      tagName,
			ArticleCount: tagCounts[tagName],
		})
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
