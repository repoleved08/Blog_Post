package dto

type PostDTO struct {
	Title   string `json:"title" form:"title" query:"title"`
	Content string `json:"content" form:"content" query:"content"`
}
