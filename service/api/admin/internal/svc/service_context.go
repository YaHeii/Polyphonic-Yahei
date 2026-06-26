// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"encoding/json"
	"fmt"

	goredis "github.com/redis/go-redis/v9"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/captcha"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/mail"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/oauth"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/oss"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/permissionx"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/tokenx"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/config"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/docs"
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
	"github.com/go-openapi/loads"
	"github.com/zeromicro/go-zero/core/logx"
	gzredis "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
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

	Redis            *gzredis.Redis
	RedisClient      *goredis.Client
	CaptchaHolder    *captcha.CaptchaHolder
	EmailDeliver     mail.IEmailDeliver
	OauthProviders   map[string]oauth.Oauth
	Uploader         oss.Uploader
	JwtTokenManager  *tokenx.JwtTokenManager
	PermissionHolder permissionx.PermissionHolder

	RequestMeta  rest.Middleware
	Permission   rest.Middleware
	OperationLog rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	var options []zrpc.ClientOption
	options = append(options)

	rds, err := ConnectRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}
	rdb := newRedisClient(c.RedisConf)

	uploader := oss.NewLocal(c.UploadConfig.RootDir(), c.UploadConfig.BaseURL())

	th := tokenx.NewJwtTokenManager(
		tokenx.NewRedisStore(rds, ""),
		c.Auth.AccessSecret,
		c.Name,
		c.Auth.AccessExpire,
		c.RefreshToken.Expire,
	)

	doc, err := loads.Analyzed(json.RawMessage(docs.Docs), "")
	if err != nil {
		panic(err)
	}

	accountRpc := accountrpc.NewAccountRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))
	permissionRpc := permissionrpc.NewPermissionRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))
	articleRpc := articlerpc.NewArticleRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))
	newsRpc := newsrpc.NewNewsRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))
	noticeRpc := noticerpc.NewNoticeRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))
	resourceRpc := resourcerpc.NewResourceRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))
	talkRpc := socialrpc.NewSocialRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))
	websiteRpc := websiterpc.NewWebsiteRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))
	configRpc := configrpc.NewConfigRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))
	syslogRpc := syslogrpc.NewSyslogRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))

	ph := permissionx.NewCasbinHolder(c.RedisConf.Host, permissionRpc)
	err = ph.LoadPolicy()
	if err != nil {
		logx.Infof("load permission policy fail: %v", err)
	}

	return &ServiceContext{
		Config:        c,
		AccountRpc:    accountRpc,
		PermissionRpc: permissionRpc,
		ArticleRpc:    articleRpc,
		NewsRpc:       newsRpc,
		NoticeRpc:     noticeRpc,
		ResourceRpc:   resourceRpc,
		SocialRpc:     talkRpc,
		WebsiteRpc:    websiteRpc,
		ConfigRpc:     configRpc,
		SyslogRpc:     syslogRpc,

		Redis:         rds,
		RedisClient:   rdb,
		CaptchaHolder: captcha.NewCaptchaHolder(captcha.WithRedisStore(rdb)),
		EmailDeliver: mail.NewEmailDeliver(&mail.EmailConfig{
			Host:     c.EmailConf.Host,
			Port:     c.EmailConf.Port,
			Username: c.EmailConf.Username,
			Password: c.EmailConf.Password,
			Nickname: c.EmailConf.Nickname,
			BCC:      c.EmailConf.BCC,
		}),
		Uploader:         uploader,
		JwtTokenManager:  th,
		PermissionHolder: ph,

		RequestMeta:  middleware.NewRequestMetaMiddleware().Handle,
		Permission:   middleware.NewPermissionMiddleware(ph).Handle,
		OperationLog: middleware.NewOperationLogMiddleware(doc.Spec(), syslogRpc, permissionRpc).Handle,
	}
}

func ConnectRedis(c config.RedisConf) (*gzredis.Redis, error) {
	address := c.Host + ":" + c.Port
	redisClient, err := gzredis.NewRedis(gzredis.RedisConf{
		Host: address,
		Type: gzredis.NodeType,
		Pass: c.Password,
		Tls:  false,
	})

	if err != nil {
		return nil, fmt.Errorf("redis 连接失败: %v", err)
	}

	return redisClient, nil
}

func newRedisClient(c config.RedisConf) *goredis.Client {
	return goredis.NewClient(&goredis.Options{
		Addr:     c.Host + ":" + c.Port,
		Password: c.Password,
		DB:       c.DB,
	})
}
