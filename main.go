package main

import (
	"log"
	"os"

	"stock-trades-api/config"
	"stock-trades-api/routes"

	"github.com/gin-gonic/gin"
	// "gorm.io/driver/sqlite"
	// "gorm.io/gorm"
)

func main() {
	db := config.SetupDatabase()
	defer func() {
		dbSQL, err := db.DB()
		if err == nil {
			dbSQL.Close()
		}
	}()

	r := gin.Default()

	routes.RegisterRoutes(r, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
