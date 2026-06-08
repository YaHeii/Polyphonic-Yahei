package website

import (
	"sort"
	"time"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/websiterpc"
)

func mergeVisitTrend(pvTrend, uvTrend []*websiterpc.VisitDailyStatistics) []types.VisitTrendVO {
	trendMap := make(map[string]types.VisitTrendVO)
	dateSet := make(map[string]struct{})

	for _, item := range pvTrend {
		trendMap[item.Date] = types.VisitTrendVO{
			Date:    item.Date,
			PvCount: item.Count,
		}
		dateSet[item.Date] = struct{}{}
	}

	for _, item := range uvTrend {
		trend := trendMap[item.Date]
		trend.Date = item.Date
		trend.UvCount = item.Count
		trendMap[item.Date] = trend
		dateSet[item.Date] = struct{}{}
	}

	dates := make([]string, 0, len(dateSet))
	for date := range dateSet {
		dates = append(dates, date)
	}
	sort.Strings(dates)

	result := make([]types.VisitTrendVO, 0, len(dates))
	for _, date := range dates {
		result = append(result, trendMap[date])
	}

	return result
}

func buildArticleStatistics(list []*articlerpc.ArticleDetails) []*types.ArticleStatisticsVO {
	byDate := make(map[string]int64)
	for _, item := range list {
		date := time.UnixMilli(item.CreatedAt).Format(time.DateOnly)
		byDate[date]++
	}

	result := make([]*types.ArticleStatisticsVO, 0, len(byDate))
	for date, count := range byDate {
		result = append(result, &types.ArticleStatisticsVO{
			Date:  date,
			Count: count,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Date < result[j].Date
	})

	return result
}
