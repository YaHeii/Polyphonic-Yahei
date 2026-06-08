// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package config

import (
	"github.com/YaHeii/Polyphonic-Yahei/pkg/oss"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	BlogRpcConf zrpc.RpcClientConf

	UploadConfig *oss.Config
	RedisConf    RedisConf
}

// redis缓存配置
type RedisConf struct {
	DB       int    `json:"db" yaml:"db"`     // redis的哪个数据库
	Host     string `json:"host" yaml:"host"` // 服务器地址:端口
	Port     string `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"` // 密码
}
