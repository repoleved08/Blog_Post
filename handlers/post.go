package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/repoleved08/blog/config"
	"github.com/repoleved08/blog/models"
)

func CreatePost(c echo.Context) (error) {
	title := c.FormValue("title")
	if title == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error":"title field cannot be empty"})
	}
	content := c.FormValue("content")
	if content == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error":"content field cannot be empty"})
	}
	// get user_id from the context
	UserID := uint(c.Get("user_id").(float64))
	post := models.Post {
		Title: title,
		Content: content,
		UserId: UserID,
	}
	config.DB.Create(&post)
	return c.JSON(http.StatusCreated, post)
}
