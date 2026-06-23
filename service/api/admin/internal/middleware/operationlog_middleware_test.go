package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/authctx"
	openapispec "github.com/go-openapi/spec"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/syslogrpc"
	"google.golang.org/grpc"
)

func TestNewOperationLogMiddlewareAcceptsGoOpenAPISwagger(t *testing.T) {
	m := NewOperationLogMiddleware(&openapispec.Swagger{}, nil, nil)
	if m == nil {
		t.Fatal("expected middleware instance")
	}
}

type stubOperationPermissionRPC struct {
	permissionrpc.PermissionRpc
}

func (s *stubOperationPermissionRPC) FindAllApi(context.Context, *permissionrpc.FindAllApiReq, ...grpc.CallOption) (*permissionrpc.FindAllApiResp, error) {
	return &permissionrpc.FindAllApiResp{
		List: []*permissionrpc.Api{
			{Path: "/admin-api/v1/article/add_article", Traceable: 1},
		},
	}, nil
}

type stubOperationSyslogRPC struct {
	syslogrpc.SyslogRpc
	req *syslogrpc.AddOperationLogReq
}

func (s *stubOperationSyslogRPC) AddOperationLog(_ context.Context, in *syslogrpc.AddOperationLogReq, _ ...grpc.CallOption) (*syslogrpc.AddOperationLogResp, error) {
	s.req = in
	return &syslogrpc.AddOperationLogResp{}, nil
}

func TestOperationLogMiddlewareReadsUserIDFromJWTContext(t *testing.T) {
	swagger := &openapispec.Swagger{
		SwaggerProps: openapispec.SwaggerProps{
			Paths: &openapispec.Paths{
				Paths: map[string]openapispec.PathItem{
					"/admin-api/v1/article/add_article": {
						PathItemProps: openapispec.PathItemProps{
							Post: &openapispec.Operation{
								OperationProps: openapispec.OperationProps{
									Summary: "add article",
									Tags:    []string{"文章管理"},
								},
							},
						},
					},
				},
			},
		},
	}
	syslogRPC := &stubOperationSyslogRPC{}
	permissionRPC := &stubOperationPermissionRPC{}
	m := NewOperationLogMiddleware(swagger, syslogRPC, permissionRPC)

	req := httptest.NewRequest(http.MethodPost, "/admin-api/v1/article/add_article", strings.NewReader(`{"title":"hello"}`))
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(context.WithValue(req.Context(), authctx.UserIDKey, "u-1"))
	rec := httptest.NewRecorder()

	m.Handle(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"ok":true}`))
	})(rec, req)

	if syslogRPC.req == nil {
		t.Fatal("expected operation log request")
	}
	if syslogRPC.req.UserId != "u-1" {
		t.Fatalf("expected jwt user id, got %#v", syslogRPC.req)
	}
	if syslogRPC.req.OptModule != "文章管理" || syslogRPC.req.OptDesc != "add article" {
		t.Fatalf("unexpected swagger metadata: %#v", syslogRPC.req)
	}
	if syslogRPC.req.ResponseStatus != http.StatusCreated {
		t.Fatalf("unexpected response status: %#v", syslogRPC.req)
	}
}
