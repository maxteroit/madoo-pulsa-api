package tests

import (
	"bytes"
	"encoding/json"
	"madoo-pulsa-api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndGetProducts(t *testing.T) {
	r, _ := setupTestRouter()

	// Login dulu untuk ambil token
	loginPayload := AuthPayload{
		Username: "testuser",
		Password: "password123",
	}
	payloadBytes, _ := json.Marshal(loginPayload)
	req := httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var loginResp map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &loginResp)
	token := loginResp["token"]

	// Create product
	product := models.Product{
		Name:        "Test Product",
		Price:       99.99,
		Qty:         10,
		ExpiredDate: "2026-01-01",
		Image:       "image.png",
	}
	productBytes, _ := json.Marshal(product)
	req = httptest.NewRequest("POST", "/api/products", bytes.NewBuffer(productBytes))
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Get all products
	req = httptest.NewRequest("GET", "/api/products", nil)
	req.Header.Set("Authorization", token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var products []models.Product
	_ = json.Unmarshal(w.Body.Bytes(), &products)
	assert.GreaterOrEqual(t, len(products), 1)
}
