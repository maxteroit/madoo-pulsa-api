package v1

import (
	"database/sql"
	"madoo-pulsa-api/models"
	"madoo-pulsa-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransactionType(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.TransactionType
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Save transaction type
		err := services.CreateTransactionType(db, input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Transaction type created"})
	}
}

func GetTransactionTypes(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		transactionTypes, err := services.GetAllTransactionTypes(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, transactionTypes)
	}
}