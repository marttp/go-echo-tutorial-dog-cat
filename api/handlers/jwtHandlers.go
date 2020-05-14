package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	log2 "log"
	"net/http"
)

func MainJwt(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	log2.Println("Username : ", claims["name"], "User ID: ", claims["jti"])
	return c.String(http.StatusOK, "You are on the top secret jwt page")
}
