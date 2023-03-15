package database

import (
	"fmt"
	"log"
	"os"

	"github.com/rodrigorvsn/beeus-api/src/domain/entities/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	database, error := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if error != nil {
		log.Fatal("⚠️ Failed to connect to database!\n", error)
		os.Exit(1)
	}

	log.Println("🔥 Database connected!\n", database)
	database.Logger = logger.Default.LogMode(logger.Info)

	log.Println("😎 Running migrations...")
	database.AutoMigrate(&user.Fact{})
}
