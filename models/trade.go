package models

type Trade struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `json:"user_id"`
	Stock     string `json:"stock"`
	Quantity  int    `json:"quantity"`
	TradeType string `json:"trade_type"` // BUY or SELL
}