package controllers

import (
	"database/sql"
	"madoo-pulsa-api/models"
	"madoo-pulsa-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Product
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := services.Create(db, input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Product created"})
	}
}

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
