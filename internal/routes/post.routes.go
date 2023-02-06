package routes

import (
	"github.com/labstack/echo/v4"
	"goEcho/internal/controllers"
)

func PostRoutes(e *echo.Group) {

	e.GET("/", controllers.GetPosts)
	e.POST("/", controllers.CreatePost)

	e.GET("/:id(\\d++)", controllers.GetPostById)
	e.PUT("/:id(\\d++)", controllers.UpdatePostById)
	e.DELETE("/:id(\\d++)", controllers.DeletePostById)
}
