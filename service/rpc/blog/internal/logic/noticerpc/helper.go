package noticerpclogic

import (
	"database/sql"
	"time"

	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/noticerpc"
)

const (
	noticeStatusDraft     int64 = 1
	noticeStatusPublished int64 = 2
	noticeStatusRevoked   int64 = 3
)

func convertNoticeOut(record *model.TSystemNotice) *noticerpc.Notice {
	if record == nil {
		return nil
	}

	return &noticerpc.Notice{
		Id:            record.Id,
		Title:         record.Title,
		Content:       record.Content,
		Type:          record.Type,
		Level:         record.Level,
		AppName:       record.AppName,
		PublisherId:   record.PublisherId,
		PublishStatus: record.PublishStatus,
		PublishTime:   nullTimeUnixMilli(record.PublishTime),
		RevokeTime:    nullTimeUnixMilli(record.RevokeTime),
		CreatedAt:     record.CreatedAt.UnixMilli(),
		UpdatedAt:     record.UpdatedAt.UnixMilli(),
	}
}

func convertNoticeListOut(records []*model.TSystemNotice) []*noticerpc.Notice {
	list := make([]*noticerpc.Notice, 0, len(records))
	for _, record := range records {
		list = append(list, convertNoticeOut(record))
	}
	return list
}

func nullTimeUnixMilli(v sql.NullTime) int64 {
	if !v.Valid {
		return 0
	}
	return v.Time.UnixMilli()
}

func buildNoticeStatusTimes(status int64, currentPublishTime sql.NullTime, currentRevokeTime sql.NullTime) (sql.NullTime, sql.NullTime) {
	now := time.Now()
	switch status {
	case noticeStatusPublished:
		if !currentPublishTime.Valid {
			currentPublishTime = sql.NullTime{Time: now, Valid: true}
		}
	case noticeStatusRevoked:
		currentRevokeTime = sql.NullTime{Time: now, Valid: true}
	}

	return currentPublishTime, currentRevokeTime
}
