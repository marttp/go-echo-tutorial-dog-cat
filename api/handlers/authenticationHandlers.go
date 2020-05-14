package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	log2 "log"
	"net/http"
	"time"
)

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func Login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	// Check username and password against DB after hashing the password
	if username == "test" && password == "1234" {
		cookie := &http.Cookie{}
		// the same as
		// cookie := new(http.Cookie)
		cookie.Name = "sessionId"
		cookie.Value = "some_string"
		cookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(cookie)

		token, err := createJwtToken()
		if err != nil {
			log2.Println("Error creation jwt token", err)
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		jwtCookie := &http.Cookie{}
		jwtCookie.Name = "JWTCookie"
		jwtCookie.Value = token
		jwtCookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(jwtCookie)

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You were logged in",
			"token":   token,
		})
	}
	return c.String(http.StatusUnauthorized, "Username or password is incorrect")
}

func createJwtToken() (string, error) {
	claims := JwtClaims{
		"test",
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err
	}
	return token, nil
}
