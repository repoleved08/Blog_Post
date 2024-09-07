package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/repoleved08/blog/config"
	"github.com/repoleved08/blog/models"
)

func Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}
	user.HashPassword(user.Password)
	user.Role = "user"
	config.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {

}
