package main

import (
	"github.com/labstack/echo/v4"
	"github.com/repoleved08/blog/config"
	"github.com/repoleved08/blog/routes"
)

func main() {
	config.InitDB()

	e := echo.New()
	routes.InitRoutes(e)

	e.Start(":8080")
}
