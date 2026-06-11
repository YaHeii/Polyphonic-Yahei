package plugins

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/plugins/knife4j"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/staticx"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/docs"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
)

func RegisterPluginHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	staticx.RegisterLocalStatic(server, serverCtx.Config.UploadConfig.BaseURL(), serverCtx.Config.UploadConfig.RootDir())

	// 注册knife4j服务
	var knife4jPrefix = "/admin-api/v1/swagger"
	server.AddRoutes(staticx.PrefixRoutes(knife4jPrefix, http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
		knife4j.NewKnife4jPlugin(docs.Docs).Handler(knife4jPrefix).ServeHTTP(w, r)
	}))
}
