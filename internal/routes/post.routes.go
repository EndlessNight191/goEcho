package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	"goEcho/internal/controllers"
)

func PostRoutes() {
	e := echo.New()

	g := e.Group("/post")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, 2!")
	})
	g.POST("/", controllers.CreatePost)

	g.GET("/:id(\\d++)", controllers.GetPostById)
	g.PUT("/:id(\\d++)", controllers.UpdatePostById)
	g.DELETE("/:id(\\d++)", controllers.DeletePostById)

	e.Logger.Fatal(e.Start(":" + viper.GetString("port")))
}
