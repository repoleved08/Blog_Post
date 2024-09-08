package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	var input models.User
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	var user models.User
	config.DB.Where("username = ?", input.Username).First(&user)

	if user.ID == 0 || user.CheckPassword(input.Password) != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid credentials"})
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "role":    user.Role,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, echo.Map{"token": tokenString})
}
