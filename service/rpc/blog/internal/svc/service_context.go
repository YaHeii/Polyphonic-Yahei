package svc

import "github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
