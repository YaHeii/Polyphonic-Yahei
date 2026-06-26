// Code scaffolded manually for aggregated blog RPC startup.

package main

import (
	"flag"
	"fmt"

	interceptorx "github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/infra/interceptor"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/config"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/configrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/newsrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/noticerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/resourcerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/socialrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/websiterpc"
	accountrpcserver "github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/server/accountrpc"
	articlerpcserver "github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/server/articlerpc"
	configrpcserver "github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/server/configrpc"
	newsrpcserver "github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/server/newsrpc"
	noticerpcserver "github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/server/noticerpc"
	permissionrpcserver "github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/server/permissionrpc"
	resourcerpcserver "github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/server/resourcerpc"
	socialrpcserver "github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/server/socialrpc"
	syslogrpcserver "github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/server/syslogrpc"
	websiterpcserver "github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/server/websiterpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/blog.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	ctx := svc.NewServiceContext(c)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		accountrpc.RegisterAccountRpcServer(grpcServer, accountrpcserver.NewAccountRpcServer(ctx))
		articlerpc.RegisterArticleRpcServer(grpcServer, articlerpcserver.NewArticleRpcServer(ctx))
		configrpc.RegisterConfigRpcServer(grpcServer, configrpcserver.NewConfigRpcServer(ctx))
		newsrpc.RegisterNewsRpcServer(grpcServer, newsrpcserver.NewNewsRpcServer(ctx))
		noticerpc.RegisterNoticeRpcServer(grpcServer, noticerpcserver.NewNoticeRpcServer(ctx))
		permissionrpc.RegisterPermissionRpcServer(grpcServer, permissionrpcserver.NewPermissionRpcServer(ctx))
		resourcerpc.RegisterResourceRpcServer(grpcServer, resourcerpcserver.NewResourceRpcServer(ctx))
		socialrpc.RegisterSocialRpcServer(grpcServer, socialrpcserver.NewSocialRpcServer(ctx))
		syslogrpc.RegisterSyslogRpcServer(grpcServer, syslogrpcserver.NewSyslogRpcServer(ctx))
		websiterpc.RegisterWebsiteRpcServer(grpcServer, websiterpcserver.NewWebsiteRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(
		interceptorx.ServerMetaInterceptor,
		interceptorx.ServerLogInterceptor,
		interceptorx.ServerErrorInterceptor,
	)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
