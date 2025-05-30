package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"madoo-pulsa-api/config"
	"madoo-pulsa-api/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() (*gin.Engine, *sql.DB) {
	gin.SetMode(gin.TestMode)
	db := config.InitDB()
	r := gin.Default()
	routes.SetupRoutes(r, db)
	return r, db
}

type AuthPayload struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func TestRegisterLoginFlow(t *testing.T) {
	r, _ := setupTestRouter()

	// Register
	registerPayload := AuthPayload{
		PhoneNumber: "081234567890",
		Password:    "password123",
	}
	payloadBytes, _ := json.Marshal(registerPayload)
	req := httptest.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Login
	loginPayload := AuthPayload{
		PhoneNumber: "081234567890",
		Password:    "password123",
	}
	payloadBytes, _ = json.Marshal(loginPayload)
	req = httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NotEmpty(t, resp["token"])
	assert.NotEmpty(t, resp["refresh_token"])
}
