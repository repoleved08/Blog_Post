package main

import (
	"github.com/labstack/echo/v4"
	"github.com/repoleved08/blog/config"
	"github.com/repoleved08/blog/routes"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/repoleved08/blog/docs"
)

// @title        Blog API
// @description  This is a simple blog API
// @version      1.0
// @host         localhost:8080
// @schemes      http https
// @securityDefinitions.ApiKey BearerAuth
// @in header
// @name Authorization

func main() {
	// Initialize the database with error handling
	config.InitDB()

	// Create a new Echo instance
	e := echo.New()

	// Swagger route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Initialize the routes
	routes.InitRoutes(e)

	// Start the server with error handling
	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
