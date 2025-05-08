package controllers

import (
	"database/sql"
	"fmt"
	"madoo-pulsa-api/models"
	"madoo-pulsa-api/services"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateProduct handles the creation of a new product.
// @Summary Buat produk baru
// @Description Tambahkan produk baru ke sistem
// @Tags products
// @Accept  json
// @Produce  json
// @Param   product  body   models.Product  true  "Product JSON"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/products [post]
func CreateProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Product
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Save product
		err := services.Create(db, input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Product created"})
	}
}

// UploadProductImage godoc
// @Summary Upload gambar produk
// @Description Mengupload gambar untuk produk berdasarkan ID
// @Tags products
// @Accept  multipart/form-data
// @Produce  json
// @Param id path string true "Product ID"
// @Param file formData file true "Product Image"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/products/{id}/upload [post]
func UploadProductImage(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// Get the uploaded file
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No image uploaded"})
			return
		}

		// Generate a unique filename for the uploaded image
		fileExtension := filepath.Ext(file.Filename)
		fileName := fmt.Sprintf("%s%s", uuid.New().String(), fileExtension)
		filePath := filepath.Join("uploads", fileName)

		// Create the uploads directory if it doesn't exist
		if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create uploads directory"})
			return
		}

		// Save the uploaded file to the uploads directory
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}

		// Retrieve the product by ID
		product, err := services.GetByID(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		// Update the product with the new image path
		product.Image = filePath
		err = services.Update(db, id, *product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
			return
		}

		// Return success response
		c.JSON(http.StatusOK, gin.H{"message": "Product image uploaded successfully"})
	}
}

// GetProducts retrieves all products
func GetProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := services.GetAll(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}

// GetProductByID retrieves a product by its ID
func GetProductByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		product, err := services.GetByID(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

// UpdateProduct updates an existing product.
func UpdateProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var input models.Product
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := services.Update(db, id, input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Product updated"})
	}
}

// DeleteProduct handles soft delete for a product
func DeleteProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := services.SoftDelete(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Product deleted (soft)"})
	}
}
