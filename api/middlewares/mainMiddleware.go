package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetMainMiddleware(e *echo.Echo) {
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "../static",
	}))
	e.Use(ServerHeader)
}

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("propose", "Testing1")
		return next(c)
	}
}
