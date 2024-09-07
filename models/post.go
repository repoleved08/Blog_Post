package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	UserId   uint      `json:"user_id"`
	User     User      `json:"-"`
	Comments []Comment `json:"comments"`
}
