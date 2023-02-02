package cmd

import (
	"configs"
	"goEcho/internal/database"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	if err := configs.GetConfig(); err != nil {
		log.Fatalf("error trying initialize config: %v", err)
	}

	if err := database.CreateConnection(); err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	// g := e.Group("/admin")
	// routes.PostRoute(g)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":" + viper.GetString("port")))
}
