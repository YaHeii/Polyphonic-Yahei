// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package tag

import (
	"net/http"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/logic/tag"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 分页获取标签列表
func FindTagListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryTagReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tag.NewFindTagListLogic(r.Context(), svcCtx)
		resp, err := l.FindTagList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
