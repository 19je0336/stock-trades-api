package services

import (
	"stock-trades-api/models"
	"gorm.io/gorm"
)

func CreateUserTrade(db *gorm.DB, trade *models.Trade) error {
	return db.Create(trade).Error
}

func GetUserTrades(db *gorm.DB, userID uint) ([]models.Trade, error) {
	var trades []models.Trade
	err := db.Where("user_id = ?", userID).Find(&trades).Error
	return trades, err
}