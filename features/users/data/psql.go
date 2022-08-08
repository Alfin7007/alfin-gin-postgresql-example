package data

import (
	"errors"
	"fmt"
	"myexample/go-gin/features/users"

	"gorm.io/gorm"
)

type psqlUserRepo struct {
	db *gorm.DB
}

func NewUserRepo(conn *gorm.DB) users.Data {
	return &psqlUserRepo{
		db: conn,
	}
}

func (repo psqlUserRepo) InsertData(userCore users.Core) (err error) {
	userModel := fromCore(userCore)
	result := repo.db.Create(&userModel)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("register fail")
	}
	return nil
}

func (repo psqlUserRepo) Login(email string) (userCore users.Core, err error) {
	userModel := User{}

	result := repo.db.Where("email = ?", email).Find(&userModel)
	if result.RowsAffected == 0 {
		return userCore, errors.New("user not found")
	}
	fmt.Println(result.Error)
	userCore = userModel.ToCore()
	return userCore, err
}
