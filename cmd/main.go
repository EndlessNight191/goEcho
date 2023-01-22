package main

import (
	"echoTest/configs"
	"echoTest/internal/database"
	"echoTest/internal/routes"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	config := configs.GetConfig()

	db := database.Init("postgres", "config.DbUrl")
	fmt.Print(config.DbUrl, config.Port)

	e := echo.New()

	g := e.Group("/admin")
	routes.PostRoute(g)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	s := ":"
	e.Logger.Fatal(e.Start(string(s) + config.Port))
}
