package auth

import (
	"net/http"
	"time"

	"nodepad-be/internal/common/middleware"
	"nodepad-be/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LogoutHandler(_ *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name: middleware.AccessTokenCookie, Value: "", Path: "/",
			HttpOnly: true, SameSite: http.SameSiteLaxMode,
			Secure: middleware.IsSecureRequest(r), MaxAge: -1,
			Expires: time.Unix(1, 0),
		})
		httpx.OkJsonCtx(r.Context(), w, map[string]string{"message": "Đăng xuất thành công"})
	}
}
