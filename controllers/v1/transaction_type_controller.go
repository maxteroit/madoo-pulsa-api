package v1

import (
	"database/sql"
	"madoo-pulsa-api/models"
	"madoo-pulsa-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"madoo-pulsa-api/utils"
)

func CreateTransactionType(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.TransactionType
		if err := c.ShouldBindJSON(&input); err != nil {
			utils.CreateResponse(c, http.StatusBadRequest, "Invalid input", nil)
			return
		}

		// Save transaction type
		err := services.CreateTransactionType(db, input)
		if err != nil {
			utils.CreateResponse(c, http.StatusInternalServerError, "Failed to create transaction type", nil)
			return
		}
		utils.CreateResponse(c, http.StatusCreated, "Transaction type created successfully", nil)
	}
}

func GetTransactionTypes(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		transactionTypes, err := services.GetAllTransactionTypes(db)
		if err != nil {
			utils.CreateResponse(c, http.StatusInternalServerError, "Failed to fetch transaction types", nil)
			return
		}
		utils.CreateResponse(c, http.StatusOK, "Transaction types fetched successfully", transactionTypes)
	}
}