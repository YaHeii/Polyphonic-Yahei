// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package talk

import (
	"net/http"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/logic/talk"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 分页获取说说列表
func FindTalkListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryTalkReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := talk.NewFindTalkListLogic(r.Context(), svcCtx)
		resp, err := l.FindTalkList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
