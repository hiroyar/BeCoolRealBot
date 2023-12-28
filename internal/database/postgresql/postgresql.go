package postgresql

import (
	"BeCoolRealBot/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func Connect() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("Connected to db")
	db.Logger = logger.Default.LogMode(logger.Info)

	err = db.AutoMigrate(&models.TelegramNotification{}, &models.TelegramUser{})
	if err != nil {
		log.Fatal("Failed to connect to migrate. \n", err)
	}

	DB = DbInstance{
		Db: db,
	}
}
