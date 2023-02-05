package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func InitRoutes() {
	e := echo.New()

	auth := e.Group("/api/auth")
	users := e.Group("/api/users")
	posts := e.Group("/api/posts")

	AuthRoutes(auth)
	UserRoutes(users)
	PostRoutes(posts)

	e.Logger.Fatal(e.Start(":" + viper.GetString("port")))
}
