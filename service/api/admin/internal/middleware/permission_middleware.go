// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package middleware

import (
	"net/http"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/authctx"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/permissionx"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/responsex"
	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionMiddleware struct {
	holder permissionx.PermissionHolder
}

func NewPermissionMiddleware(holder permissionx.PermissionHolder) *PermissionMiddleware {
	return &PermissionMiddleware{
		holder: holder,
	}
}

// 权限拦截
func (m *PermissionMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Debugf("PermissionMiddleware Handle path: %v", r.URL.Path)
		uid := authctx.CurrentUserID(r.Context())
		if uid == "" {
			responsex.Response(r, w, nil, bizerr.NewBizError(bizcode.CodeUserUnLogin, "user id missing"))
			return
		}

		// 验证用户是否有权限访问资源
		err := m.holder.Enforce(uid, r.URL.Path, r.Method)
		if err != nil {
			responsex.Response(r, w, nil, bizerr.NewBizError(bizcode.CodeUserNotPermission, err.Error()))
			return
		}

		// 调用下一层的处理
		next.ServeHTTP(w, r)
	}
}
