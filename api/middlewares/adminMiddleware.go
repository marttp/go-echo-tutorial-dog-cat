package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetAdminMiddleware(g *echo.Group) {
	//adminGroup := e.Group("/admin", middleware.Logger())
	//adminGroup.Use(middleware.Logger())
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "test" && password == "1234" {
			return true, nil
		}
		return false, nil
	}))
}
