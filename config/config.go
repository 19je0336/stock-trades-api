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

	var count int64
	db.Model(&models.User{}).Where("username = ?", "testuser").Count(&count)
	if count == 0 {
		db.Create(&models.User{
			Username: "testuser",
			Password: "testpass",
		})
	}

	DB = db
	return db
}
