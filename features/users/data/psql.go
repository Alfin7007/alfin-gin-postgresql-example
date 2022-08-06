package data

import (
	"errors"
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
	return userCore, err
}
