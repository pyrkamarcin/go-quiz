package routes

import (
	"github.com/labstack/echo"
	"net/http"
)

func Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}
