package dto

type CommentDTO struct {
	Content string `json:"content" validate:"required"` 
	PostID  uint   `json:"post_id" validate:"required"` 
}
