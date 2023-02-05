package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"goEcho/internal/controllers"
)

func PostRoutes(e *echo.Group) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, POST!")
	})
	e.POST("/", controllers.CreatePost)

	e.GET("/:id(\\d++)", controllers.GetPostById)
	e.PUT("/:id(\\d++)", controllers.UpdatePostById)
	e.DELETE("/:id(\\d++)", controllers.DeletePostById)

}
