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
		api.POST("/products", controllers.CreateProduct(db))
		api.GET("/products", controllers.GetProducts(db))
	}
}
