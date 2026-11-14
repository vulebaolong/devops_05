// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package auth

import (
	"net/http"

	"nodepad-be/internal/logic/auth"
	"nodepad-be/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func MeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewMeLogic(r.Context(), svcCtx)
		resp, err := l.Me()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
