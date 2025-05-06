package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"stock-trades-api/config"
	"stock-trades-api/routes"
)

func TestLoginEndpoint(t *testing.T) {
	db := config.SetupDatabase()
	r := gin.Default()
	routes.RegisterRoutes(r, db)

	w := httptest.NewRecorder()
	body := `{"username":"testuser","password":"testpass"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 200 or 401, got %d", w.Code)
	}
}