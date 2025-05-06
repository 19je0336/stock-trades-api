package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"stock-trades-api/config"
	"stock-trades-api/routes"
)

func TestProtectedTradeRoute(t *testing.T) {
	db := config.SetupDatabase()
	r := gin.Default()
	routes.RegisterRoutes(r, db)

	req, _ := http.NewRequest("GET", "/trades", nil)
	// Missing token should return unauthorized
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401 for missing token, got %d", w.Code)
	}
}