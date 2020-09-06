package routes

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pyrkamarcin/go-quiz/src/helpers"
	"github.com/pyrkamarcin/go-quiz/src/models"
	"github.com/pyrkamarcin/go-quiz/src/requests"
)

func Login(db *gorm.DB) func(ctx echo.Context) error {
	return func(c echo.Context) error {

		r := &requests.UserLoginRequest{}
		if err := c.Bind(r); err != nil {
			return err
		}

		u := &models.User{}
		db.Model(&models.User{}).Where("email = ?", r.Email).Take(&u)

		if err := helpers.CheckPasswordHash(r.Password, u.Password); err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": err.Error(),
			})
		}

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = u.Email
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		refreshToken := jwt.New(jwt.SigningMethodHS256)
		rtClaims := refreshToken.Claims.(jwt.MapClaims)
		rtClaims["sub"] = 1
		rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		rt, err := refreshToken.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return c.JSON(http.StatusBadGateway, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token":         t,
			"refresh_token": rt,
		})
	}
}
