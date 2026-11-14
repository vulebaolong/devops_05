package middleware

import (
	"net/http"
	"nodepad-be/internal/common/apperror"
	"runtime/debug"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

const AccessTokenCookie = "accessToken"

func IsSecureRequest(r *http.Request) bool {
	return r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https"
}

func CookieAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			if cookie, err := r.Cookie(AccessTokenCookie); err == nil && cookie.Value != "" {
				r.Header.Set("Authorization", "Bearer "+cookie.Value)
			}
		}
		next(w, r)
	}
}

func JSONRecover(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			recovered := recover()
			if recovered == nil {
				return
			}

			logc.Errorf(
				r.Context(),
				"panic recovered: value=%v\n%s",
				recovered,
				debug.Stack(),
			)

			httpx.ErrorCtx(
				r.Context(),
				w,
				apperror.InternalServerError(),
			)
		}()

		next(w, r)
	}
}
