package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/repoleved08/blog/controllers"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/api/auth/register", controllers.Register)
	e.POST("/api/auth/login", controllers.Login)
}
