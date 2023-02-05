package routes

import (
	"github.com/labstack/echo/v4"
	"goEcho/internal/controllers"
)

func AuthRoutes(e *echo.Group) {
	e.POST("/", controllers.Registration) // регистрация
	e.PATCH("/", controllers.Auth)        // аторизация  обычная
	e.PUT("/", controllers.RefreshToken)  // refresh token
	e.DELETE("/", controllers.AllLogout)  // all logout

}
