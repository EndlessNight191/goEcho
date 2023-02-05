package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	"goEcho/internal/controllers"
)

func PostRoutes() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world")
	})

	g := e.Group("/api/post")

	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, POST!")
	})
	g.POST("/", controllers.CreatePost)

	g.GET("/:id(\\d++)", controllers.GetPostById)
	g.PUT("/:id(\\d++)", controllers.UpdatePostById)
	g.DELETE("/:id(\\d++)", controllers.DeletePostById)

	e.Logger.Fatal(e.Start(":" + viper.GetString("port")))
}
