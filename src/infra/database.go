package database

import (
	"log"
	"os"

	"github.com/rodrigoRVSN/beeus-api/src/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDb() {
	database, error := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if error != nil {
		log.Fatal("⚠️ Failed to connect to database!\n", error)
		os.Exit(1)
	}

	log.Println("🔥 Database connected!\n", database)
	database.Logger = logger.Default.LogMode(logger.Info)

	log.Println("😎 Running migrations...")
	database.AutoMigrate(&entities.User{})

	DB = DbInstance{
		Db: database,
	}
}
