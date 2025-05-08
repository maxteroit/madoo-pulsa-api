package routes

import (
	"database/sql"
	"madoo-pulsa-api/controllers"
	"madoo-pulsa-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	auth := r.Group("/auth")
	auth.POST("/register", controllers.Register(db))
	auth.POST("/login", controllers.Login(db))

	api := r.Group("/api")
	api.Use(middleware.JWTAuth())
	{
		// Endpoint untuk upload gambar produk
		api.POST("/products", controllers.CreateProduct(db))       // Create product
		api.GET("/products", controllers.GetProducts(db))          // Get all products
		api.GET("/products/:id", controllers.GetProductByID(db))   // Get product by ID
		api.PUT("/products/:id", controllers.UpdateProduct(db))    // Update product
		api.DELETE("/products/:id", controllers.DeleteProduct(db)) // Soft delete product

		// Endpoint untuk upload gambar produk
		api.POST("/products/upload/:id", controllers.UploadProductImage(db))
	}
}
