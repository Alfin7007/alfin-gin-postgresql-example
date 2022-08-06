package data

import (
	"myexample/go-gin/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func fromCore(userCore users.Core) User {
	userModel := User{
		Name:     userCore.Name,
		Email:    userCore.Email,
		Password: userCore.Password,
	}
	return userModel
}
