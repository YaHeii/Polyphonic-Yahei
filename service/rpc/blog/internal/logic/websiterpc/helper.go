package websiterpclogic

import (
	"context"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/YaHeii/Polyphonic-Yahei/common/rediskey"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/cryptox"
	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/common/rpcutils"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/websiterpc"
)

const visitDateLayout = "2006-01-02"

func currentVisitDate() string {
	return time.Now().Format(visitDateLayout)
}

func previousVisitDate() string {
	return time.Now().AddDate(0, 0, -1).Format(visitDateLayout)
}

func resolveVisitorID(ctx context.Context) string {
	if terminalID, _ := rpcutils.GetTerminalIdFromCtx(ctx); strings.TrimSpace(terminalID) != "" {
		return strings.TrimSpace(terminalID)
	}

	if userID, _ := rpcutils.GetUserIdFromCtx(ctx); strings.TrimSpace(userID) != "" {
		return "user:" + strings.TrimSpace(userID)
	}

	remoteIP, _ := rpcutils.GetRemoteIPFromCtx(ctx)
	remoteAgent, _ := rpcutils.GetRemoteAgentFromCtx(ctx)
	remoteIP = strings.TrimSpace(remoteIP)
	remoteAgent = strings.TrimSpace(remoteAgent)
	if remoteIP != "" || remoteAgent != "" {
		return cryptox.Md5v(remoteIP, remoteAgent)
	}

	return "anonymous"
}

func markDailyVisitor(ctx context.Context, rdb *redis.Client, date string, visitorID string) (bool, error) {
	key := rediskey.GetDailyUserVisitKey(date)
	added, err := rdb.SAdd(ctx, key, visitorID).Result()
	if err != nil {
		return false, err
	}

	expireAt := truncateDate(time.Now()).AddDate(0, 0, 2)
	if err := rdb.ExpireAt(ctx, key, expireAt).Err(); err != nil {
		return false, err
	}

	return added > 0, nil
}

func calculateGrowthRate(current int64, previous int64) float64 {
	if previous == 0 {
		if current == 0 {
			return 0
		}
		return 100
	}
	return float64(current-previous) / float64(previous) * 100
}

func normalizeVisitRange(startDate string, endDate string) (time.Time, time.Time, error) {
	now := time.Now()
	start := now.AddDate(0, 0, -6)
	end := now

	if strings.TrimSpace(startDate) != "" {
		parsed, err := time.ParseInLocation(visitDateLayout, strings.TrimSpace(startDate), time.Local)
		if err != nil {
			return time.Time{}, time.Time{}, err
		}
		start = parsed
	}

	if strings.TrimSpace(endDate) != "" {
		parsed, err := time.ParseInLocation(visitDateLayout, strings.TrimSpace(endDate), time.Local)
		if err != nil {
			return time.Time{}, time.Time{}, err
		}
		end = parsed
	}

	start = truncateDate(start)
	end = truncateDate(end)
	if start.After(end) {
		start, end = end, start
	}

	return start, end, nil
}

func buildVisitTrend(records []*model.TVisitDailyStats, start time.Time, end time.Time) []*websiterpc.VisitDailyStatistics {
	counts := make(map[string]int64, len(records))
	for _, record := range records {
		counts[record.Date] = record.ViewCount
	}

	list := make([]*websiterpc.VisitDailyStatistics, 0, int(end.Sub(start).Hours()/24)+1)
	for day := start; !day.After(end); day = day.AddDate(0, 0, 1) {
		date := day.Format(visitDateLayout)
		list = append(list, &websiterpc.VisitDailyStatistics{
			Date:  date,
			Count: counts[date],
		})
	}
	return list
}

func truncateDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}
