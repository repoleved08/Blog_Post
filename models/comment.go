package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model `swaggerignore:"true"`
	Content    string `json:"content"`
	PostId     uint   `json:"post_id"`
	Post       Post   `json:"-"`
	UserId     uint   `json:"user_id"`
	User       User   `json:"-"`
}
