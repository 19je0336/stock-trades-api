package config

import (
	"log"
	"stock-trades-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("trades.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	db.AutoMigrate(&models.User{}, &models.Trade{})
	DB = db
	return db
}
