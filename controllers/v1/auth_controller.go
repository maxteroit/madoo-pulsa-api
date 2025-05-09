package v1

import (
	"database/sql"
	"madoo-pulsa-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"madoo-pulsa-api/utils"
)

type AuthInput struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func Register(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input AuthInput
		if err := c.ShouldBindJSON(&input); err != nil {
			utils.CreateResponse(c, http.StatusBadRequest, "Invalid input", nil)
			return
		}
		err := services.Register(db, input.PhoneNumber, input.Password)
		if err != nil {
			utils.CreateResponse(c, http.StatusInternalServerError, "Failed to register user", nil)
			return
		}
		utils.CreateResponse(c, http.StatusCreated, "User registered successfully", nil)
	}
}

func Login(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input AuthInput
		if err := c.ShouldBindJSON(&input); err != nil {
			utils.CreateResponse(c, http.StatusBadRequest, "Invalid input", nil)
			return
		}
		token, refresh, err := services.Login(db, input.PhoneNumber, input.Password)
		if err != nil {
			utils.CreateResponse(c, http.StatusUnauthorized, "Invalid credentials", nil)
			return
		}
		utils.CreateResponse(c, http.StatusOK, "Login successful", gin.H{
			"token":         token,
			"refresh_token": refresh,
		})
	}
}
