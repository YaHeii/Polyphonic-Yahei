package accountrpclogic

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/YaHeii/Polyphonic-Yahei/common/enums"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizheader"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/cryptox"
	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/config"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"
	gzcache "github.com/zeromicro/go-zero/core/stores/cache"
	gzredis "github.com/zeromicro/go-zero/core/stores/redis"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

func getUserRoles(ctx context.Context, svcCtx *svc.ServiceContext, uid string) ([]*model.TRole, error) {
	return svcCtx.TRoleModel.FindRolesByUserID(ctx, uid)
}

func onLogin(ctx context.Context, svcCtx *svc.ServiceContext, user *model.TUser, loginType string) (*accountrpc.LoginResp, error) {
	if user.Status == enums.UserStatusDisabled {
		return nil, bizerr.NewBizError(bizcode.CodeUserDisabled, "用户已被禁用")
	}

	rList, err := getUserRoles(ctx, svcCtx, user.UserId)
	if err != nil {
		return nil, err
	}

	if err := svcCtx.OnlineUserService.Login(ctx, user.UserId); err != nil {
		return nil, err
	}

	return &accountrpc.LoginResp{
		User:      convertUserInfoOut(user, rList),
		LoginType: loginType,
	}, nil
}

func convertUserInfoOut(user *model.TUser, roles []*model.TRole) *accountrpc.UserInfo {
	if user == nil {
		return nil
	}

	return &accountrpc.UserInfo{
		UserId:       user.UserId,
		Username:     user.Username,
		Nickname:     user.Nickname,
		Avatar:       user.Avatar,
		Email:        user.Email,
		Phone:        user.Phone,
		Info:         user.Info,
		Status:       user.Status,
		RegisterType: user.RegisterType,
		IpAddress:    user.IpAddress,
		IpSource:     user.IpSource,
		CreatedAt:    user.CreatedAt.Unix(),
		UpdatedAt:    user.UpdatedAt.Unix(),
		Roles:        convertRoleLabels(roles),
	}
}

func convertRoleLabels(roles []*model.TRole) []*accountrpc.UserRoleLabel {
	labels := make([]*accountrpc.UserRoleLabel, 0, len(roles))
	for _, role := range roles {
		labels = append(labels, &accountrpc.UserRoleLabel{
			RoleId:      role.Id,
			RoleKey:     role.RoleKey,
			RoleComment: role.RoleComment,
		})
	}

	return labels
}

func convertUserOut(user *model.TUser) *accountrpc.User {
	if user == nil {
		return nil
	}

	return &accountrpc.User{
		UserId:       user.UserId,
		Username:     user.Username,
		Nickname:     user.Nickname,
		Avatar:       user.Avatar,
		Email:        user.Email,
		Phone:        user.Phone,
		Info:         user.Info,
		Status:       user.Status,
		RegisterType: user.RegisterType,
		IpAddress:    user.IpAddress,
		IpSource:     user.IpSource,
		CreatedAt:    user.CreatedAt.Unix(),
		UpdatedAt:    user.UpdatedAt.Unix(),
	}
}

func convertVisitorInfoOut(visitor *model.TVisitor) *accountrpc.VisitorInfo {
	if visitor == nil {
		return nil
	}

	return &accountrpc.VisitorInfo{
		Id:         visitor.Id,
		TerminalId: visitor.TerminalId,
		Os:         visitor.Os,
		Browser:    visitor.Browser,
		IpAddress:  visitor.IpAddress,
		IpSource:   visitor.IpSource,
		CreatedAt:  visitor.CreatedAt.Unix(),
		UpdatedAt:  visitor.UpdatedAt.Unix(),
	}
}

func convertUserOauthInfoOut(item *model.TUserOauth) *accountrpc.UserOauthInfo {
	if item == nil {
		return nil
	}

	return &accountrpc.UserOauthInfo{
		Platform:  item.Platform,
		OpenId:    item.OpenId,
		Nickname:  item.Nickname,
		Avatar:    item.Avatar,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
}

func buildRegisterModelCacheConf(c config.RedisConf) gzcache.CacheConf {
	return gzcache.CacheConf{
		{
			RedisConf: gzredis.RedisConf{
				Host: c.Host + ":" + c.Port,
				Type: "node",
				Pass: c.Password,
			},
			Weight: 100,
		},
	}
}

func getRemoteIPFromContext(ctx context.Context) string {
	if p, ok := peer.FromContext(ctx); ok && p.Addr != nil {
		if host, _, err := net.SplitHostPort(p.Addr.String()); err == nil {
			return host
		}
		return p.Addr.String()
	}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		for _, key := range []string{"x-forwarded-for", "x-real-ip"} {
			values := md.Get(key)
			if len(values) == 0 {
				continue
			}

			ip := strings.TrimSpace(strings.Split(values[0], ",")[0])
			if ip != "" {
				return ip
			}
		}
	}

	return "127.0.0.1"
}

func getIncomingMetadataValue(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	values := md.Get(key)
	if len(values) == 0 {
		return ""
	}

	return strings.TrimSpace(values[0])
}

func getCurrentUserID(ctx context.Context) (string, error) {
	uid := getIncomingMetadataValue(ctx, bizheader.HeaderUid)
	if uid == "" {
		return "", bizerr.NewBizError(bizcode.CodeUserUnLogin, "用户未登录")
	}

	return uid, nil
}

func getCurrentUser(ctx context.Context, svcCtx *svc.ServiceContext) (*model.TUser, error) {
	uid, err := getCurrentUserID(ctx)
	if err != nil {
		return nil, err
	}

	user, err := svcCtx.TUserModel.FindOneByUserId(ctx, uid)
	if err != nil {
		return nil, bizerr.NewBizError(bizcode.CodeUserNotExist, "用户不存在")
	}

	return user, nil
}

func getPageParams(pageReq *accountrpc.PageReq) (page int64, pageSize int64, offset int64) {
	page = 1
	pageSize = 10
	if pageReq != nil {
		if pageReq.Page > 0 {
			page = pageReq.Page
		}
		if pageReq.PageSize > 0 {
			pageSize = pageReq.PageSize
		}
	}

	offset = (page - 1) * pageSize
	return
}

func buildPageResp(page, pageSize, total int64) *accountrpc.PageResp {
	return &accountrpc.PageResp{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}
}

func appendStringLikeCondition(conditions []string, args []any, field string, value string) ([]string, []any) {
	if strings.TrimSpace(value) == "" {
		return conditions, args
	}

	conditions = append(conditions, fmt.Sprintf("%s like $%d", field, len(args)+1))
	args = append(args, "%"+strings.TrimSpace(value)+"%")
	return conditions, args
}

func appendStringEqualCondition(conditions []string, args []any, field string, value string) ([]string, []any) {
	if strings.TrimSpace(value) == "" {
		return conditions, args
	}

	conditions = append(conditions, fmt.Sprintf("%s = $%d", field, len(args)+1))
	args = append(args, strings.TrimSpace(value))
	return conditions, args
}

func appendInt64EqualCondition(conditions []string, args []any, field string, value int64) ([]string, []any) {
	conditions = append(conditions, fmt.Sprintf("%s = $%d", field, len(args)+1))
	args = append(args, value)
	return conditions, args
}

func appendStringInCondition(conditions []string, args []any, field string, values []string) ([]string, []any) {
	if len(values) == 0 {
		return conditions, args
	}

	holders := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}

		holders = append(holders, "$"+strconv.Itoa(len(args)+1))
		args = append(args, value)
	}

	if len(holders) == 0 {
		return conditions, args
	}

	conditions = append(conditions, fmt.Sprintf("%s in (%s)", field, strings.Join(holders, ",")))
	return conditions, args
}

func joinWhereClause(conditions []string) string {
	if len(conditions) == 0 {
		return ""
	}

	return " where " + strings.Join(conditions, " and ")
}

func resolveTerminalID(ctx context.Context) string {
	if terminalID := getIncomingMetadataValue(ctx, bizheader.HeaderXTerminalId); terminalID != "" {
		return terminalID
	}

	userAgent := getIncomingMetadataValue(ctx, bizheader.HeaderUserAgent)
	ip := getIncomingMetadataValue(ctx, bizheader.HeaderRPCRemoteIP)
	if ip == "" {
		ip = getRemoteIPFromContext(ctx)
	}

	if userAgent == "" && ip == "" {
		return ""
	}

	return cryptox.Md5v(userAgent, ip)
}

func resolveClientOS(ctx context.Context) string {
	if osType := getIncomingMetadataValue(ctx, bizheader.HeaderOsType); osType != "" {
		return osType
	}
	return "unknown"
}

func resolveClientBrowser(ctx context.Context) string {
	if remoteAgent := getIncomingMetadataValue(ctx, bizheader.HeaderRPCRemoteAgent); remoteAgent != "" {
		return remoteAgent
	}
	if userAgent := getIncomingMetadataValue(ctx, bizheader.HeaderUserAgent); userAgent != "" {
		return userAgent
	}
	return "unknown"
}

func findArea(city string) string {
	if strings.Contains(city, "省") {
		for _, area := range provinces {
			if strings.HasPrefix(city, area) {
				return area
			}
		}
	} else if strings.Contains(city, "自治区") {
		for _, area := range autonomousRegions {
			if strings.HasPrefix(city, area) {
				return area
			}
		}
	} else if strings.Contains(city, "特别行政区") {
		for _, area := range specialAdministrativeRegions {
			if strings.HasPrefix(city, area) {
				return area
			}
		}
	} else {
		for _, area := range municipalities {
			if strings.HasPrefix(city, area) {
				return area
			}
		}
	}

	return "未知"
}

var provinces = []string{
	"河北", "山西", "辽宁", "吉林", "黑龙江",
	"江苏", "浙江", "安徽", "福建", "江西", "山东",
	"河南", "湖北", "湖南", "广东", "海南",
	"四川", "贵州", "云南", "陕西", "甘肃", "青海", "台湾",
}

var municipalities = []string{
	"北京市", "天津市", "上海市", "重庆市",
}

var autonomousRegions = []string{
	"内蒙古", "广西", "西藏", "宁夏", "新疆",
}

var specialAdministrativeRegions = []string{
	"香港", "澳门",
}
