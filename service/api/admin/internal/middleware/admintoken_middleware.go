// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizcode"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizerr"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizheader"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/responsex"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/infra/tokenx"
	"github.com/zeromicro/go-zero/core/logx"
)

type AdminTokenMiddleware struct {
	verifier *tokenx.JwtTokenManager
}

func NewAdminTokenMiddleware(verifier *tokenx.JwtTokenManager) *AdminTokenMiddleware {
	return &AdminTokenMiddleware{
		verifier: verifier,
	}
}

func (m *AdminTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Debugf("AdminTokenMiddleware Handle")
		var token string
		var uid string

		token = r.Header.Get(bizheader.HeaderAuthorization)
		uid = r.Header.Get(bizheader.HeaderUid)

		// 请求头缺少参数
		if uid == "" {
			responsex.Response(r, w, nil, bizerr.NewBizError(bizcode.CodeInvalidParam, fmt.Sprintf("request header field '%v' is missing", bizheader.HeaderUid)))
			return
		}

		if token == "" {
			responsex.Response(r, w, nil, bizerr.NewBizError(bizcode.CodeInvalidParam, fmt.Sprintf("request header field '%v' is missing", bizheader.HeaderAuthorization)))
			return
		}

		err := m.verifier.ValidateToken(uid, token)
		if err != nil {
			if errors.Is(err, tokenx.ErrTokenExpired) {
				responsex.Response(r, w, nil, bizerr.NewBizError(bizcode.CodeUserLoginExpired, err.Error()))
				return
			}
			responsex.Response(r, w, nil, bizerr.NewBizError(bizcode.CodeUserUnLogin, err.Error()))
			return
		}

		next.ServeHTTP(w, r)
	}
}
