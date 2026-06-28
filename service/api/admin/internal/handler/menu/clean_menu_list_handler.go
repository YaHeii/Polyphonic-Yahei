// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package menu

import (
	"net/http"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/logic/menu"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 清空菜单列表
func CleanMenuListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmptyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewCleanMenuListLogic(r.Context(), svcCtx)
		resp, err := l.CleanMenuList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
