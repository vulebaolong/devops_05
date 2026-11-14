// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package middleware

import (
	"net/http"

	commonmiddleware "nodepad-be/internal/common/middleware"

	"github.com/zeromicro/go-zero/rest/handler"
)

type CookieAuthMiddleware struct {
	secret string
}

func NewCookieAuthMiddleware(secret string) *CookieAuthMiddleware {
	return &CookieAuthMiddleware{secret: secret}
}

func (m *CookieAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	authorize := handler.Authorize(m.secret)(next)
	return commonmiddleware.CookieAuth(authorize.ServeHTTP)
}
