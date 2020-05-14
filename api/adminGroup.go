package api

import (
	"github.com/labstack/echo"
	"go-echo-sample/api/handlers"
)

func AdminGroup(e *echo.Group) {
	e.GET("/main", handlers.MainAdmin)
	//adminGroup.GET("/main", mainAdmin, middleware.Logger())
}
