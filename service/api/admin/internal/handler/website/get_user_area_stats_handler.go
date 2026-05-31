// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package website

import (
	"net/http"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/logic/website"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取用户分布地区
func GetUserAreaStatsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserAreaStatsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := website.NewGetUserAreaStatsLogic(r.Context(), svcCtx)
		resp, err := l.GetUserAreaStats(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
