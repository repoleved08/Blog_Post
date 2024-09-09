package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/repoleved08/blog/config"
	"github.com/repoleved08/blog/models"
)


func CreateComment(c echo.Context) (error) {
	var comment models.Comment
	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error":"invalid input"})
	}
	config.DB.Create(comment)
	return c.JSON(http.StatusCreated, comment)
}

func GetCommentByD(c echo.Context) (error) {
	id := c.Param("id")
	var comment models.Comment
	if err := config.DB.Find(&comment, id); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error":"comment not found"})
	}
	return c.JSON(http.StatusOK, comment)
}

// func UpdateComent(c echo.Context) (error)
