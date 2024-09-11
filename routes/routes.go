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
	e.GET("/api/posts/:id", handlers.GetPostByID)
	e.PUT("/api/posts/:id", handlers.UpdatePost, middleware.JWTMiddleware)
	e.DELETE("/api/posts/:id", handlers.DeletePost, middleware.JWTMiddleware)
	// comments endpoints
	e.POST("/api/posts/comments", handlers.CreateComment, middleware.JWTMiddleware)
	e.GET("/api/posts/comment/:id", handlers.GetCommentById)
	e.GET("/api/posts/comments/:id", handlers.GetCommentsByPostId)
	e.DELETE("/api/posts/comments/:id", handlers.DeleteComment, middleware.JWTMiddleware)
}
