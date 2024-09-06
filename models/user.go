package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
}
