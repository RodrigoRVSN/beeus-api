package db

import (
	"log"
	"os"

	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase() (*Database, error) {
	databaseUrl := os.Getenv("DB_URL")

	database, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	log.Println("🔥 Database connected!\n", database)
	database.Logger = logger.Default.LogMode(logger.Info)

	log.Println("😎 Running migrations...")
	database.AutoMigrate(&entity.User{})

	return &Database{db: database}, nil
}

func (d *Database) GetDB() *gorm.DB {
	return d.db
}
