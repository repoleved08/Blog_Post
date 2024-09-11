package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/repoleved08/blog/config"
	"github.com/repoleved08/blog/models"
	"github.com/repoleved08/blog/models/dto"
	"gorm.io/gorm"
)

// @Summary Create a new comment
// @Description Create a new comment
// @Tags comments
// @Accept json
// @Produce json
// @Param commentDTO body dto.CommentDTO true "Comment DTO"
// @Security BearerAuth
// @Success 201 {object} models.Comment
// @Failure 400 {object} map[string]string
// @Router /api/posts/comments [post]
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


// @Summary Get comment by id
// @Description Get comment by id
// @Tags comments
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} models.Comment
// @Failure 404 {object} map[string]string
// @Router /api/posts/comment/{id} [get]
func GetCommentById(c echo.Context) error {
	id := c.Param("id")
	var comment models.Comment
	if err := config.DB.Find(&comment, id); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "comment not found"})
	}
	return c.JSON(http.StatusOK, comment)
}
// ENDPOINT TO GET COMMENT BY POST ID
// @Summary Get comments by post ID
// @Description Get comments by post ID
// @Tags comments
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {array} models.Comment
// @Failure 404 {object} map[string]string
// @Router /api/posts/comments/{id} [get]
func GetCommentsByPostId(c echo.Context) error {
	postId := c.Param("id")
	var comments []models.Comment
	if err := config.DB.Where("post_id = ?", postId).Find(&comments).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "no comments found"})
	}
	return c.JSON(http.StatusOK, comments)
}

// delete comment
// @Summary Delete comment
// @Description Delete comment
// @Tags comments
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Security BearerAuth
// @Success 204 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /api/posts/comments/{id} [delete]
func DeleteComment(c echo.Context) error {
	id := c.Param("id")
	var comment models.Comment

	// Find the comment by ID
	if err := config.DB.First(&comment, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "comment not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "could not retrieve comment"})
	}

	// Get the user ID from the JWT token in the context
	UserID := uint(c.Get("user_id").(float64))
	// Get the user role from the JWT token in the context (assuming it's included)
	userRole := c.Get("role").(string)

	// Check if the user is the owner of the comment or has the admin role
	if comment.UserId != UserID && userRole != "admin" {
		return c.JSON(http.StatusForbidden, echo.Map{"error": "you are not authorized to delete this comment"})
	}

	// Proceed with deletion if authorized
	config.DB.Delete(&comment)
	return c.NoContent(http.StatusNoContent)
}

