package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/repoleved08/blog/handlers"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/api/auth/register", handlers.Register)
	e.POST("/api/auth/login", handlers.Login)
}
