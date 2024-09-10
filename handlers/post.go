package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/repoleved08/blog/config"
	"github.com/repoleved08/blog/models"
)
// swagger documenation
//	@Summary		Create a new post
//	@Description	Create a new post
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			title	formData	string	true	"Post title"
//	@Param			content	formData	string	true	"Post content"
//	@Security		BearerAuth
//	@Success		201	{object}	models.Post
//	@Failure		400	{object}	map[string]string
//	@Router			/api/posts [post]

func CreatePost(c echo.Context) (error) {
	title := c.FormValue("title")
	if title == "" {
		return sendErrorResponse(c, http.StatusBadRequest, "title field cannot be empty")
	}
	content := c.FormValue("content")
	if content == "" {
		return sendErrorResponse(c, http.StatusBadRequest, "content field cannot be empty")
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

// swagger documenation
//	@Summary		Get all posts
//	@Description	Get all posts
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Post
//	@Failure		404	{object}	map[string]string
//	@Router			/api/posts [get]
func GetPosts(c echo.Context) error {
	var posts []models.Post
	result := config.DB.Find(&posts)

	// Check for any errors in the DB operation
	if result.Error != nil {
		return sendErrorResponse(c, http.StatusInternalServerError, "failed to fetch posts")
	}
	// Check if no rows are found
	if result.RowsAffected == 0 {
		return sendErrorResponse(c, http.StatusNotFound, "no posts found")
	}
	// Return posts in the JSON response
	return c.JSON(http.StatusOK, posts)
}

