package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"madoo-pulsa-api/models"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

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
	expiredDate, err := time.Parse("2006-01-02 15:04:05", "2026-01-01 00:00:00")
	if err != nil {
		t.Fatalf("Error parsing ExpiredDate: %v", err)
	}

	product := models.Product{
		Name:        "Test Product",
		Price:       99.99,
		Qty:         10,
		ExpiredDate: expiredDate,
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

func TestUploadProductImage(t *testing.T) {
	r, _ := setupTestRouter()

	// Login untuk ambil token
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

	// Upload gambar
	file, err := os.Open("test_image.png")
	if err != nil {
		t.Fatalf("Failed to open test image: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatalf("Failed to copy file contents: %v", err)
	}
	err = writer.Close()
	if err != nil {
		t.Fatalf("Failed to close writer: %v", err)
	}

	req = httptest.NewRequest("POST", "/api/products/1/upload", body)
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verify response
	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Image uploaded successfully", response["message"])

	// Verify the product image path was updated
	req = httptest.NewRequest("GET", "/api/products/1", nil)
	req.Header.Set("Authorization", token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var product models.Product
	_ = json.Unmarshal(w.Body.Bytes(), &product)
	assert.Equal(t, "test_image.png", product.Image)
}
