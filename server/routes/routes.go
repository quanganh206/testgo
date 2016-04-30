package routes

import (
	"github.com/quanganh206/testgo/server/api/todo/routes"
	"github.com/quanganh206/testgo/server/common/static"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	static.Init(e)
	todoroutes.Init(e)
}
