package syslogrpclogic

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/common/query"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/common/rpcutils"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/syslogrpc"
)

func convertAddLoginLogIn(ctx context.Context, in *syslogrpc.AddLoginLogReq) *model.TLoginLog {
	userID := in.UserId
	if userID == "" {
		userID = metadataUserID(ctx)
	}

	return &model.TLoginLog{
		UserId:     userID,
		TerminalId: metadataTerminalID(ctx),
		LoginType:  in.LoginType,
		AppName:    metadataAppName(ctx),
		LoginAt:    time.Now(),
	}
}

func convertLoginLogOut(record *model.TLoginLog) *syslogrpc.LoginLog {
	if record == nil {
		return nil
	}

	return &syslogrpc.LoginLog{
		Id:         record.Id,
		UserId:     record.UserId,
		TerminalId: record.TerminalId,
		LoginType:  record.LoginType,
		AppName:    record.AppName,
		LoginAt:    unixMilli(record.LoginAt),
		LogoutAt:   nullTimeUnixMilli(record.LogoutAt),
	}
}

func convertLoginLogListOut(records []*model.TLoginLog) []*syslogrpc.LoginLog {
	list := make([]*syslogrpc.LoginLog, 0, len(records))
	for _, record := range records {
		list = append(list, convertLoginLogOut(record))
	}
	return list
}

func convertAddVisitLogIn(ctx context.Context, in *syslogrpc.AddVisitLogReq) *model.TVisitLog {
	return &model.TVisitLog{
		UserId:     metadataUserID(ctx),
		TerminalId: metadataTerminalID(ctx),
		PageName:   in.PageName,
	}
}

func convertVisitLogOut(record *model.TVisitLog) *syslogrpc.VisitLog {
	if record == nil {
		return nil
	}

	return &syslogrpc.VisitLog{
		Id:         record.Id,
		UserId:     record.UserId,
		TerminalId: record.TerminalId,
		PageName:   record.PageName,
		CreatedAt:  unixMilli(record.CreatedAt),
		UpdatedAt:  unixMilli(record.UpdatedAt),
	}
}

func convertVisitLogListOut(records []*model.TVisitLog) []*syslogrpc.VisitLog {
	list := make([]*syslogrpc.VisitLog, 0, len(records))
	for _, record := range records {
		list = append(list, convertVisitLogOut(record))
	}
	return list
}

func convertAddOperationLogIn(ctx context.Context, in *syslogrpc.AddOperationLogReq) *model.TOperationLog {
	userID := in.UserId
	if userID == "" {
		userID = metadataUserID(ctx)
	}

	terminalID := in.TerminalId
	if terminalID == "" {
		terminalID = metadataTerminalID(ctx)
	}

	return &model.TOperationLog{
		UserId:         userID,
		TerminalId:     terminalID,
		OptModule:      in.OptModule,
		OptDesc:        in.OptDesc,
		RequestUri:     in.RequestUri,
		RequestMethod:  in.RequestMethod,
		RequestData:    in.RequestData,
		ResponseData:   in.ResponseData,
		ResponseStatus: in.ResponseStatus,
		Cost:           in.Cost,
	}
}

func convertOperationLogOut(record *model.TOperationLog) *syslogrpc.OperationLog {
	if record == nil {
		return nil
	}

	return &syslogrpc.OperationLog{
		Id:             record.Id,
		UserId:         record.UserId,
		TerminalId:     record.TerminalId,
		OptModule:      record.OptModule,
		OptDesc:        record.OptDesc,
		RequestUri:     record.RequestUri,
		RequestMethod:  record.RequestMethod,
		RequestData:    record.RequestData,
		ResponseData:   record.ResponseData,
		ResponseStatus: record.ResponseStatus,
		Cost:           record.Cost,
		CreatedAt:      unixMilli(record.CreatedAt),
		UpdatedAt:      unixMilli(record.UpdatedAt),
	}
}

func convertOperationLogListOut(records []*model.TOperationLog) []*syslogrpc.OperationLog {
	list := make([]*syslogrpc.OperationLog, 0, len(records))
	for _, record := range records {
		list = append(list, convertOperationLogOut(record))
	}
	return list
}

func convertAddFileLogIn(ctx context.Context, in *syslogrpc.AddFileLogReq) *model.TFileLog {
	userID := in.UserId
	if userID == "" {
		userID = metadataUserID(ctx)
	}

	terminalID := in.TerminalId
	if terminalID == "" {
		terminalID = metadataTerminalID(ctx)
	}

	return &model.TFileLog{
		Id:         in.Id,
		UserId:     userID,
		TerminalId: terminalID,
		FilePath:   in.FilePath,
		FileName:   in.FileName,
		FileType:   in.FileType,
		FileSize:   in.FileSize,
		FileMd5:    in.FileMd5,
		FileUrl:    in.FileUrl,
	}
}

func convertFileLogOut(record *model.TFileLog) *syslogrpc.FileLog {
	if record == nil {
		return nil
	}

	return &syslogrpc.FileLog{
		Id:         record.Id,
		UserId:     record.UserId,
		TerminalId: record.TerminalId,
		FilePath:   record.FilePath,
		FileName:   record.FileName,
		FileType:   record.FileType,
		FileSize:   record.FileSize,
		FileMd5:    record.FileMd5,
		FileUrl:    record.FileUrl,
		CreatedAt:  unixMilli(record.CreatedAt),
		UpdatedAt:  unixMilli(record.UpdatedAt),
	}
}

func convertFileLogListOut(records []*model.TFileLog) []*syslogrpc.FileLog {
	list := make([]*syslogrpc.FileLog, 0, len(records))
	for _, record := range records {
		list = append(list, convertFileLogOut(record))
	}
	return list
}

func buildLoginLogQuery(in *syslogrpc.FindLoginLogListReq) (int, int, string, string, []any) {
	var opts []query.Option
	if in.Paginate != nil {
		opts = append(opts, query.WithPage(int(in.Paginate.Page)))
		opts = append(opts, query.WithSize(int(in.Paginate.PageSize)))
		opts = append(opts, query.WithSorts(in.Paginate.Sorts...))
	}
	if strings.TrimSpace(in.UserId) != "" {
		opts = append(opts, query.WithCondition("user_id = ?", strings.TrimSpace(in.UserId)))
	}
	return query.NewQueryBuilder(opts...).Build()
}

func buildVisitLogQuery(in *syslogrpc.FindVisitLogListReq) (int, int, string, string, []any) {
	var opts []query.Option
	if in.Paginate != nil {
		opts = append(opts, query.WithPage(int(in.Paginate.Page)))
		opts = append(opts, query.WithSize(int(in.Paginate.PageSize)))
		opts = append(opts, query.WithSorts(in.Paginate.Sorts...))
	}
	if strings.TrimSpace(in.UserId) != "" {
		opts = append(opts, query.WithCondition("user_id = ?", strings.TrimSpace(in.UserId)))
	}
	if strings.TrimSpace(in.TerminalId) != "" {
		opts = append(opts, query.WithCondition("terminal_id = ?", strings.TrimSpace(in.TerminalId)))
	}
	if strings.TrimSpace(in.PageName) != "" {
		opts = append(opts, query.WithCondition("page_name like ?", "%"+strings.TrimSpace(in.PageName)+"%"))
	}
	return query.NewQueryBuilder(opts...).Build()
}

func buildOperationLogQuery(in *syslogrpc.FindOperationLogListReq) (int, int, string, string, []any) {
	var opts []query.Option
	if in.Paginate != nil {
		opts = append(opts, query.WithPage(int(in.Paginate.Page)))
		opts = append(opts, query.WithSize(int(in.Paginate.PageSize)))
		opts = append(opts, query.WithSorts(in.Paginate.Sorts...))
	}
	if keywords := strings.TrimSpace(in.Keywords); keywords != "" {
		like := "%" + keywords + "%"
		opts = append(opts, query.WithCondition("(opt_module like ? or opt_desc like ? or request_uri like ? or request_method like ?)", like, like, like, like))
	}
	return query.NewQueryBuilder(opts...).Build()
}

func buildFileLogQuery(in *syslogrpc.FindFileLogListReq) (int, int, string, string, []any) {
	var opts []query.Option
	if in.Paginate != nil {
		opts = append(opts, query.WithPage(int(in.Paginate.Page)))
		opts = append(opts, query.WithSize(int(in.Paginate.PageSize)))
		opts = append(opts, query.WithSorts(in.Paginate.Sorts...))
	}
	if strings.TrimSpace(in.FilePath) != "" {
		opts = append(opts, query.WithCondition("file_path like ?", "%"+strings.TrimSpace(in.FilePath)+"%"))
	}
	if strings.TrimSpace(in.FileName) != "" {
		opts = append(opts, query.WithCondition("file_name like ?", "%"+strings.TrimSpace(in.FileName)+"%"))
	}
	if strings.TrimSpace(in.FileType) != "" {
		opts = append(opts, query.WithCondition("file_type = ?", strings.TrimSpace(in.FileType)))
	}
	return query.NewQueryBuilder(opts...).Build()
}

func buildPageResp(page int, size int, total int64) *syslogrpc.PageResp {
	return &syslogrpc.PageResp{
		Page:     int64(page),
		PageSize: int64(size),
		Total:    total,
	}
}

func metadataUserID(ctx context.Context) string {
	value, _ := rpcutils.GetUserIdFromCtx(ctx)
	return value
}

func metadataTerminalID(ctx context.Context) string {
	value, _ := rpcutils.GetTerminalIdFromCtx(ctx)
	return value
}

func metadataAppName(ctx context.Context) string {
	value, _ := rpcutils.GetAppNameFromCtx(ctx)
	return value
}

func logoutTime(in *syslogrpc.AddLogoutLogReq) time.Time {
	if in.LogoutAt > 0 {
		return time.UnixMilli(in.LogoutAt)
	}
	return time.Now()
}

func unixMilli(v time.Time) int64 {
	if v.IsZero() {
		return 0
	}
	return v.UnixMilli()
}

func nullTimeUnixMilli(v sql.NullTime) int64 {
	if !v.Valid {
		return 0
	}
	return unixMilli(v.Time)
}
