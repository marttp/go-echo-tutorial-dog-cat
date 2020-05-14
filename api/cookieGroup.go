package api

import (
	"github.com/labstack/echo"
	"go-echo-sample/api/handlers"
)

func CookieGroup(e *echo.Group) {
	e.GET("/main", handlers.MainCookie)
}
