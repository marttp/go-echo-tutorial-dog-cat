package api

import (
	"github.com/labstack/echo"
	"go-echo-sample/api/handlers"
)

func JwtGroup(e *echo.Group) {
	e.GET("/main", handlers.MainJwt)
}
