package controllers

import (
	"net/http"
	"stock-trades-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTrade(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := c.MustGet("user_id").(uint)

	var trade models.Trade
	if err := c.ShouldBindJSON(&trade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	trade.UserID = userID
	db.Create(&trade)
	c.JSON(http.StatusOK, trade)
}

func GetTrades(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := c.MustGet("user_id").(uint)

	var trades []models.Trade
	db.Where("user_id = ?", userID).Find(&trades)
	c.JSON(http.StatusOK, trades)
}