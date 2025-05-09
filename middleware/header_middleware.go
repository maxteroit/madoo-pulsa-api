package middleware

import (
	"net/http"
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RequireHeader(headerName string, envKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerValue := c.GetHeader(headerName)
		expectedValue := os.Getenv(envKey)

		if expectedValue == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%s not configured in environment", envKey)})
			c.Abort()
			return
		}

		if headerValue != expectedValue {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Invalid %s", headerName)})
			c.Abort()
			return
		}

		c.Next()
	}
}