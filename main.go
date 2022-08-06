package main

import (
	"myexample/go-gin/config"
	"myexample/go-gin/migrations"
)

func main() {
	db := config.PostgreSQLConfig()
	migrations.DBMigrator(db)
}
