package migrations

import (
	articleData "myexample/go-gin/features/articles/data"
	userData "myexample/go-gin/features/users/data"

	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) {
	db.AutoMigrate(userData.User{})
	db.AutoMigrate(articleData.Article{})
}
