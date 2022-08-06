package migrations

import (
	productData "myexample/go-gin/features/products/data"
	userData "myexample/go-gin/features/users/data"

	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) {
	db.AutoMigrate(userData.User{})
	db.AutoMigrate(productData.Product{})
}
