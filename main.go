package main

import (
	"myexample/go-gin/config"
	"myexample/go-gin/factory"
	"myexample/go-gin/migrations"
	"myexample/go-gin/routes"
)

func main() {
	db := config.PostgreSQLConfig()
	migrations.DBMigrator(db)

	presenter := factory.InitFactory(db)
	router := routes.New(presenter)
	router.SetTrustedProxies(nil)

	router.Run("localhost:8000")
}
