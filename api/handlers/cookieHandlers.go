package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func MainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "You are on cookie page")
}
