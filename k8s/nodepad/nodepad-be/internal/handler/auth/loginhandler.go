// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package auth

import (
	"net/http"
	"time"

	"nodepad-be/internal/common/apperror"
	"nodepad-be/internal/common/middleware"
	"nodepad-be/internal/logic/auth"
	"nodepad-be/internal/svc"
	"nodepad-be/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, apperror.BadRequest(err.Error()))
			return
		}

		l := auth.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			http.SetCookie(w, &http.Cookie{
				Name: "accessToken", Value: resp.AccessToken, Path: "/",
				HttpOnly: true, SameSite: http.SameSiteLaxMode,
				Secure: middleware.IsSecureRequest(r), MaxAge: int(resp.ExpiresIn),
				Expires: time.Now().Add(time.Duration(resp.ExpiresIn) * time.Second),
			})
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
