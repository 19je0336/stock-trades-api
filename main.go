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
	// Add middleware here if needed (e.g., logging, CORS, recovery)

	routes.RegisterRoutes(r, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
