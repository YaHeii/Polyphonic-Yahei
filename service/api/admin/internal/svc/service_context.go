// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/config"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config       config.Config
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
