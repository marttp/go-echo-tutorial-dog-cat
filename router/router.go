package router

import (
	"github.com/labstack/echo"
	"go-echo-sample/api"
	"go-echo-sample/api/middlewares"
)

func New() *echo.Echo {
	e := echo.New()

	// Create groups
	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	// Set all middleware
	middlewares.SetMainMiddleware(e)
	middlewares.SetAdminMiddleware(adminGroup)
	middlewares.SetCookieMiddleware(cookieGroup)
	middlewares.SetJwtMiddlewares(jwtGroup)

	// Set main routes
	api.MainGroup(e)

	// Set group routes
	api.AdminGroup(adminGroup)
	api.CookieGroup(cookieGroup)
	api.JwtGroup(jwtGroup)

	return e
}
