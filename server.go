package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pyrkamarcin/go-quiz/src/models"
	"github.com/pyrkamarcin/go-quiz/src/routes"
)

func main() {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.LogMode(true)
	db.AutoMigrate(&models.User{})

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("login", routes.Login(db))
	e.POST("register", routes.Register(db))

	e.GET("open", routes.Accessible)

	r := e.Group("/close")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("secret"),
		AuthScheme:  "Bearer",
		TokenLookup: "header:Authorization",
	}))
	r.GET("", routes.Restricted)

	e.Logger.Fatal(e.Start(":5000"))
}
