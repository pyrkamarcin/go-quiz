package routes

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
)

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	return c.JSON(http.StatusOK, map[string]string{
		"email": email,
	})
}
