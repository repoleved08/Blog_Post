package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/repoleved08/blog/config"
	"github.com/repoleved08/blog/models"
)

//	@Summary		Create a new post
//	@Description	Create a new post
//	@Tags			posts
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			title	formData	string	true	"Post title"
//	@Param			content	formData	string	true	"Post content"
//	@Security		BearerAuth
//	@Success		201	{object}	models.Post
//	@Failure		400	{object}	map[string]string
//	@Router			/api/posts [post]
func CreatePost(c echo.Context) error {
	// Get form values
	title := c.FormValue("title")
	content := c.FormValue("content")
	
	// Log the values for debugging
	log.Printf("Title: '%s', Content: '%s'", title, content)
	if title == "" {
		return sendErrorResponse(c, http.StatusBadRequest, "title field cannot be empty")
	}
	
	if content == "" {
		return sendErrorResponse(c, http.StatusBadRequest, "content field cannot be empty")
	}

	// Get user_id from the context
	userID, ok := c.Get("user_id").(float64)
	if !ok {
		return sendErrorResponse(c, http.StatusInternalServerError, "user ID not found in context")
	}

	post := models.Post{
		Title:   title,
		Content: content,
		UserId:  uint(userID),
	}

	result := config.DB.Create(&post)
	if result.Error != nil {
		return sendErrorResponse(c, http.StatusInternalServerError, "failed to create post")
	}

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
	
	// Use Preload to load the associated Comments for each Post
	result := config.DB.Preload("Comments").Find(&posts)

	// Check for any errors in the DB operation
	if result.Error != nil {
		return sendErrorResponse(c, http.StatusInternalServerError, "failed to fetch posts")
	}
	// Check if no rows are found
	if result.RowsAffected == 0 {
		return sendErrorResponse(c, http.StatusNotFound, "no posts found")
	}
	// Return posts in the JSON response along with their comments
	return c.JSON(http.StatusOK, posts)
}


// swagger documenation
//	@Summary		Get a post by ID
//	@Description	Get a post by ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Post ID"
//	@Success		200	{object}	models.Post
//	@Failure		404	{object}	map[string]string
//	@Router			/api/posts/{id} [get]
func GetPostByID(c echo.Context) error {
	id := c.Param("id")
	var post models.Post
	result := config.DB.Preload("Comments").First(&post, id)

	// Checking if the post was not found
	if result.Error != nil {
		return sendErrorResponse(c, http.StatusNotFound, "post not found")
	}

	return c.JSON(http.StatusOK, post)
}


// swagger documenation
//	@Summary		Update a post
//	@Description	Update a post
//	@Tags			posts
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			id	path	string	true	"Post ID"
//	@Param			title	formData	string	true	"Post title"
//	@Param			content	formData	string	true	"Post content"
//	@Security		BearerAuth
//	@Success		200	{object}	models.Post
//	@Failure		400	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Router			/api/posts/{id} [put]
func UpdatePost(c echo.Context) error {
	id := c.Param("id")
	title := c.FormValue("title")
	content := c.FormValue("content")

	// Check if the title is empty
	if title == "" {
		return sendErrorResponse(c, http.StatusBadRequest, "title field cannot be empty")
	}

	// Check if the content is empty
	if content == "" {
		return sendErrorResponse(c, http.StatusBadRequest, "content field cannot be empty")
	}

	// Get the post from the database
	var post models.Post
	result := config.DB.First(&post, id)

	// Check if the post was not found
	if result.Error != nil {
		return sendErrorResponse(c, http.StatusNotFound, "post not found")
	}

	// Update the post
	post.Title = title
	post.Content = content

	// Save the updated post
	result = config.DB.Save(&post)
	if result.Error != nil {
		return sendErrorResponse(c, http.StatusInternalServerError, "failed to update post")
	}

	return c.JSON(http.StatusOK, post)
}

// swagger documenation
//	@Summary		Delete a post
//	@Description	Delete a post
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Post ID"
//	@Security		BearerAuth
//	@Success		204	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Router			/api/posts/{id} [delete]
func DeletePost(c echo.Context) error {
	id := c.Param("id")

	// Get the post from the database
	var post models.Post
	result := config.DB.First(&post, id)

	// Check if the post was not found
	if result.Error != nil {
		return sendErrorResponse(c, http.StatusNotFound, "post not found")
	}

	// Delete the post
	result = config.DB.Delete(&post)
	if result.Error != nil {
		return sendErrorResponse(c, http.StatusInternalServerError, "failed to delete post")
	}

	return c.JSON(http.StatusNoContent, nil)
}
