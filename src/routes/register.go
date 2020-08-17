package routes

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pyrkamarcin/go-quiz/src/helpers"
	"github.com/pyrkamarcin/go-quiz/src/models"
	"github.com/pyrkamarcin/go-quiz/src/requests"
	"net/http"
)

func Register(db *gorm.DB) func(ctx echo.Context) error {
	return func(c echo.Context) error {

		r := &requests.UserRegisterRequest{}
		if err := c.Bind(r); err != nil {
			return err
		}

		u := &models.User{}
		u.Email = r.Email
		u.Password, _ = helpers.HashPassword(r.Password)

		if err := db.Create(u).Error; err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]string{
			"status": "ok",
		})
	}
}
