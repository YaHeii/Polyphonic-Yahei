// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"github.com/YaHeii/Polyphonic-Yahei/pkg/oss"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/tokenx"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/permissionx"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/config"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/middleware"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/configrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/newsrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/noticerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/resourcerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/socialrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/syslogrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/websiterpc"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config

	AccountRpc    accountrpc.AccountRpc
	PermissionRpc permissionrpc.PermissionRpc
	ArticleRpc    articlerpc.ArticleRpc
	NewsRpc       newsrpc.NewsRpc
	NoticeRpc     noticerpc.NoticeRpc
	ResourceRpc   resourcerpc.ResourceRpc
	SocialRpc     socialrpc.SocialRpc
	WebsiteRpc    websiterpc.WebsiteRpc
	ConfigRpc     configrpc.ConfigRpc
	SyslogRpc     syslogrpc.SyslogRpc

	Redis            *redis.Redis
	Uploader         oss.Uploader //TODO
	TokenManager     tokenx.TokenManager
	PermissionHolder permissionx.PermissionHolder


	AdminToken   rest.Middleware
	Permission   rest.Middleware
	OperationLog rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		AdminToken:   middleware.NewAdminTokenMiddleware().Handle,
		Permission:   middleware.NewPermissionMiddleware().Handle,
		OperationLog: middleware.NewOperationLogMiddleware().Handle,
	}
}
