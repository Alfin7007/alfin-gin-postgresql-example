package data

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name   string `json:"name"`
	Stock  uint   `json:"stock"`
	Price  uint   `json:"price"`
	Detail string `json:"detail"`
}
