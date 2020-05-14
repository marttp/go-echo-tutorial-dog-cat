package api

import (
	"github.com/labstack/echo"
	"go-echo-sample/api/handlers"
)

func MainGroup(e *echo.Echo) {
	e.GET("/", handlers.Yallo)
	e.GET("/login", handlers.Login)
	e.GET("/cats/:data", handlers.GetCats)
	e.POST("/cats", handlers.AddCat)
	e.POST("/dogs", handlers.AddDog)
	e.POST("/hamsters", handlers.AddHamster)
}
