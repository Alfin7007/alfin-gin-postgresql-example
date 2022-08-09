package data

import (
	"errors"
	"myexample/go-gin/features/articles"

	"gorm.io/gorm"
)

type psqlArticleRepo struct {
	db *gorm.DB
}

func NewArticleRepo(conn *gorm.DB) articles.Data {
	return &psqlArticleRepo{
		db: conn,
	}
}

func (repo psqlArticleRepo) InsertData(artCore articles.Core) error {
	artModel := fromCore(artCore)
	result := repo.db.Create(&artModel)
	if result.RowsAffected == 0 {
		return errors.New("failed insert")
	}
	return nil
}

func (repo psqlArticleRepo) SelectAll() ([]articles.Core, error) {
	var artModel []Article
	result := repo.db.Find(&artModel)
	if result.RowsAffected == 0 {
		return nil, errors.New("failed get data")
	}
	return toCoreList(artModel), nil
}

func (repo psqlArticleRepo) SelectData(id int) (articles.Core, error) {
	var artModel Article
	result := repo.db.Preload("User").Where("id = ?", id).First(&artModel)
	if result.RowsAffected == 0 {
		return articles.Core{}, errors.New("failed get data")
	}
	return artModel.toCore(), nil
}
