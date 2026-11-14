// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package main

import (
	"flag"
	"fmt"
	"net/http"

	"nodepad-be/internal/common/apperror"
	"nodepad-be/internal/common/middleware"
	"nodepad-be/internal/config"
	"nodepad-be/internal/handler"
	"nodepad-be/internal/svc"

	"github.com/joho/godotenv"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/nodepad-api.yaml", "the config file")

func main() {
	flag.Parse()

	_ = godotenv.Load()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	server := rest.MustNewServer(
		c.RestConf,
		rest.WithCors(
			"http://localhost:3000",
			"https://nodepad.vulebaolong.com",
		),
		rest.WithNotFoundHandler(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				httpx.ErrorCtx(r.Context(), w, apperror.NotFoundRoute(fmt.Sprintf("This API was not found. Route %s %s does not exist.", r.Method, r.URL.Path)))
			}),
		),
		rest.WithNotAllowedHandler(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				httpx.ErrorCtx(r.Context(), w, apperror.NotFoundMethod(fmt.Sprintf("Method %s is not allowed for route %s.", r.Method, r.URL.Path)))
			}),
		),
	)
	defer server.Stop()

	// Ghi log đồng thời ra file và console.
	// logx.AddWriter(logx.NewWriter(os.Stdout))
	server.Use(middleware.JSONRecover)
	server.Use(middleware.CookieAuth)

	ctx, cleanup, err := svc.NewServiceContext(c)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandlerCtx(apperror.AppErrorHandler())

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
