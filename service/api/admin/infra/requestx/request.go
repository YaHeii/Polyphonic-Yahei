package requestx

import (
	"context"
	"net/http"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizheader"
)

// 请求上下文,一般存放请求头参数
// TODO:如果作为整个应用的传递的话是不对的，缺少trace ID,并且不应该将token传递
type Context struct {
	context.Context `json:"-" header:"-"`
	Token           string `json:"token" header:"token" example:""`
	Uid             string `json:"uid" header:"-" example:""`
	IpAddress       string `json:"ip_address" header:"-" example:""`
	UserAgent       string `json:"user_agent" header:"-" example:""`
}

func (s *Context) GetContext() context.Context {
	return s.Context
}

// 获取请求上下文
func ParseRequestContext(r *http.Request) *Context {
	reqCtx := &Context{
		Context:   r.Context(),
		Token:     r.Header.Get(bizheader.HeaderToken),
		Uid:       r.Header.Get(bizheader.HeaderUid),
		IpAddress: r.RemoteAddr,
		UserAgent: r.UserAgent(),
	}

	return reqCtx
}
