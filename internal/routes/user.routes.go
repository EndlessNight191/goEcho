package routes

import (
	"github.com/labstack/echo/v4"
	"goEcho/internal/controllers"
)

func UserRoutes(e *echo.Group) {
	e.GET("/", controllers.GetCurrentUser)
	e.PATCH("/", controllers.UpdateCurrentUser)
}
