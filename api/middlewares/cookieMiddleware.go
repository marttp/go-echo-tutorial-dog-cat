package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"strings"
)

func SetCookieMiddleware(g *echo.Group) {
	g.Use(checkCookie)
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionId")

		if err != nil {
			log.Print(err)
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "You don't have any cookie")
			}
			return err
		}

		if cookie.Value == "some_string" {
			return next(c)
		}
		return c.String(http.StatusUnauthorized, "You don't have the right cookie")
	}
}
