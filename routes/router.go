package routes

import (
	"database/sql"
	apiv1 "madoo-pulsa-api/controllers/v1"
	// apiv2 "madoo-pulsa-api/controllers/v2" // contoh jika ada versi lain
	"madoo-pulsa-api/middleware"

	"github.com/gin-gonic/gin"
	// "os"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {

	v1 := r.Group("/api/v1")
	v1.Use(middleware.RequireHeader("X-Api-Key", "API_KEY"))
	auth := v1.Group("/auth")
	auth.POST("/register", apiv1.Register(db))
	auth.POST("/login", apiv1.Login(db))

	protected := v1.Group("/")
	protected.Use(middleware.JWTAuth())
	{
		products := protected.Group("/products")
		{
			products.POST("/", apiv1.CreateProduct(db))       // Create product
			products.GET("/", apiv1.GetProducts(db))          // Get all products
			products.GET("/:id", apiv1.GetProductByID(db))     // Get product by ID
			products.PUT("/:id", apiv1.UpdateProduct(db))      // Update product
			products.DELETE("/:id", apiv1.DeleteProduct(db))   // Soft delete product

			products.POST("/upload/:id", apiv1.UploadProductImage(db)) // Upload product image
		}

		transaction_types := protected.Group("/transaction-types")
		{
			transaction_types.POST("/", apiv1.CreateTransactionType(db)) // Create transaction type
			transaction_types.GET("/", apiv1.GetTransactionTypes(db))     // Get all transaction types
			// transaction_types.GET("/:id", apiv1.GetTransactionTypeByID(db)) // Get transaction type by ID
			// transaction_types.PUT("/:id", apiv1.UpdateTransactionType(db)) // Update transaction type		
			// transaction_types.DELETE("/:id", apiv1.DeleteTransactionType(db)) // Soft delete transaction type
		}
	}
}