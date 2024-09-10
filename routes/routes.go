package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/repoleved08/blog/handlers"
	"github.com/repoleved08/blog/middleware"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/api/auth/register", handlers.Register)
	e.POST("/api/auth/login", handlers.Login)
	// posts endpoints
	e.POST("/api/posts", handlers.CreatePost, middleware.JWTMiddleware)
	e.GET("/api/posts", handlers.GetPosts)
	// comments endpoints
	e.POST("/api/posts/:id/comments", handlers.CreateComment, middleware.JWTMiddleware)
}
