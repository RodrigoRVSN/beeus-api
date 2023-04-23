package config

import (
	"github.com/rodrigoRVSN/beeus-api/src/infra/db"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	database, error := db.NewDatabase()

	if error != nil {
		panic(error)
	}

	return database.GetDB()
}
