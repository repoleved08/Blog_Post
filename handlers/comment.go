package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/repoleved08/blog/config"
	"github.com/repoleved08/blog/models"
	"github.com/repoleved08/blog/models/dto"
	"gorm.io/gorm"
)

func CreateComment(c echo.Context) error {
	var commentDTO dto.CommentDTO
	if err := c.Bind(&commentDTO); err != nil {
		return sendErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	// Validate if the post exists
	var post models.Post
	if result := config.DB.First(&post, commentDTO.PostID); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return sendErrorResponse(c, http.StatusNotFound, "post not found")
		}
		return sendErrorResponse(c, http.StatusInternalServerError, "failed to validate post")
	}

	UserID := uint(c.Get("user_id").(float64))

	comment := models.Comment{
		Content: commentDTO.Content,
		PostId:  commentDTO.PostID,
		UserId:  UserID,
	}

	result := config.DB.Create(&comment)
	if result.Error != nil {
		return sendErrorResponse(c, http.StatusInternalServerError, "failed to create comment")
	}

	return c.JSON(http.StatusCreated, comment)
}



func GetCommentByD(c echo.Context) error {
	id := c.Param("id")
	var comment models.Comment
	if err := config.DB.Find(&comment, id); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "comment not found"})
	}
	return c.JSON(http.StatusOK, comment)
}

// func UpdateComent(c echo.Context) (error)
func DeleteComment(c echo.Context) error {
	id := c.Param("id")
	var comment models.Comment
	if err := config.DB.Find(&comment, id); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "comment not found"})
	}
	config.DB.Delete(&comment)
	return c.NoContent(http.StatusNoContent)
}
