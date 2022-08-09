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

func (repo psqlUserRepo) InsertUser(userCore users.Core) (err error) {
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

func (repo psqlUserRepo) FindUser(email string) (userCore users.Core, err error) {
	userModel := User{}

	result := repo.db.Where("email = ?", email).Find(&userModel)
	if result.RowsAffected == 0 {
		return userCore, errors.New("user not found")
	}

	userCore = userModel.ToCore()
	return userCore, nil
}

func (repo psqlUserRepo) SelectUser(id int) (userCore users.Core, err error) {
	userModel := User{}

	result := repo.db.Where("id = ?", id).Find(&userModel)
	if result.RowsAffected == 0 {
		return userCore, errors.New("user not found")
	}

	userCore = userModel.ToCore()
	return userCore, nil
}
