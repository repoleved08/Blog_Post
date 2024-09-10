package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `swaggerignore:"true"`
	Username   string `json:"username" gorm:"unique"`
	Email      string `json:"email"`
	Password   string `json:"-"`
	Role       string `json:"role"`
}

func (user *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
}
