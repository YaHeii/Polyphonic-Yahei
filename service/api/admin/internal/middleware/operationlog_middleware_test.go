package middleware

import (
	"testing"

	openapispec "github.com/go-openapi/spec"
)

func TestNewOperationLogMiddlewareAcceptsGoOpenAPISwagger(t *testing.T) {
	m := NewOperationLogMiddleware(&openapispec.Swagger{}, nil, nil)
	if m == nil {
		t.Fatal("expected middleware instance")
	}
}
