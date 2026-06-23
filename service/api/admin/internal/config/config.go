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

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	BlogRpcConf zrpc.RpcClientConf

	UploadConfig   *oss.Config
	RedisConf      RedisConf
	EmailConf      EmailConf
	RefreshToken   RefreshTokenConf
	ThirdPartyConf map[string]map[string]ThirdPartyInfo
}

// redis缓存配置
type RedisConf struct {
	DB       int    `json:"db" yaml:"db"`     // redis的哪个数据库
	Host     string `json:"host" yaml:"host"` // 服务器地址:端口
	Port     string `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"` // 密码
}

type EmailConf struct {
	Host     string   `json:"host"`
	Port     int      `json:"port"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Nickname string   `json:"nickname"`
	BCC      []string `json:"bcc"`
}

type RefreshTokenConf struct {
	Expire int64 `json:"expire" yaml:"expire"`
}

type ThirdPartyInfo struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectUri  string `json:"redirect_uri"`
}
