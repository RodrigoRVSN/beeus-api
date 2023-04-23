package configs

import (
	"log"
	"os"

	"github.com/rodrigoRVSN/beeus-api/src/application/users/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

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

	DB = database
}
