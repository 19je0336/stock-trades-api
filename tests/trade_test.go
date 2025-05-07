package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"stock-trades-api/config"
	"stock-trades-api/routes"

	"github.com/gin-gonic/gin"
)

func TestProtectedTradeRoute(t *testing.T) {
	db := config.SetupDatabase()
	r := gin.Default()
	routes.RegisterRoutes(r, db)

	req, _ := http.NewRequest("GET", "/trades", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401 for missing token, got %d", w.Code)
	}
}
