package services

import (
	"stock-trades-api/models"
	"gorm.io/gorm"
)

func AuthenticateUser(db *gorm.DB, username, password string) (*models.User, error) {
	var user models.User
	err := db.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
