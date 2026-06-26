package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestBlogServerRegistersCustomUnaryInterceptors(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "blog.go", nil, parser.ImportsOnly|parser.ParseComments)
	if err != nil {
		t.Fatalf("ParseFile imports failed: %v", err)
	}

	foundImport := false
	for _, imp := range file.Imports {
		if imp.Path != nil && imp.Path.Value == `"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/infra/interceptor"` {
			foundImport = true
			break
		}
	}
	if !foundImport {
		t.Fatal("expected blog.go to import blog rpc interceptors")
	}

	file, err = parser.ParseFile(fset, "blog.go", nil, 0)
	if err != nil {
		t.Fatalf("ParseFile full failed: %v", err)
	}

	var foundCall bool
	var foundArgs []string
	ast.Inspect(file, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok || sel.Sel == nil || sel.Sel.Name != "AddUnaryInterceptors" {
			return true
		}

		foundCall = true
		for _, arg := range call.Args {
			switch v := arg.(type) {
			case *ast.SelectorExpr:
				if pkg, ok := v.X.(*ast.Ident); ok {
					foundArgs = append(foundArgs, pkg.Name+"."+v.Sel.Name)
				}
			}
		}

		return true
	})

	if !foundCall {
		t.Fatal("expected blog.go to register custom unary interceptors")
	}

	want := []string{
		"interceptorx.ServerMetaInterceptor",
		"interceptorx.ServerLogInterceptor",
		"interceptorx.ServerErrorInterceptor",
	}
	if len(foundArgs) != len(want) {
		t.Fatalf("unexpected interceptor count: got=%v want=%v", foundArgs, want)
	}
	for i := range want {
		if foundArgs[i] != want[i] {
			t.Fatalf("unexpected interceptors: got=%v want=%v", foundArgs, want)
		}
	}
}
