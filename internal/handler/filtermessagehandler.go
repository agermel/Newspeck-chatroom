package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"newspeak-chat/internal/logic"
	"newspeak-chat/internal/svc"
	"newspeak-chat/internal/types"
)

func filterMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FilterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFilterMessageLogic(r.Context(), svcCtx)
		resp, err := l.FilterMessage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
