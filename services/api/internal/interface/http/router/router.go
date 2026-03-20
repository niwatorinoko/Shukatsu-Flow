package router

import (
	"github.com/labstack/echo/v4"

	gen "shukatsu-flow/api/internal/interface/http/gen"
)

func RegisterRoutes(
	echoServer *echo.Echo,
	serverHandler gen.ServerInterface,
) {
	gen.RegisterHandlers(echoServer, serverHandler)
}
