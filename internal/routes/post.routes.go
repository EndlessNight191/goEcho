package routes

import (
	//"goEcho/internal/controllers"
	"github.com/labstack/echo/v4"
	"net/http"
)

func PostRoutes(e *echo.Group) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, 2!")
	})

	//e.GET("/:id", controllers.FindPost)
	//e.PUT("/:id", controllers.UpdatePost)
	//e.DELETE("/:id", controllers.DeletePost)
}
