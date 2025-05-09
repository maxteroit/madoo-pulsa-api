package main

import (
	"log"
	"madoo-pulsa-api/config"
	"madoo-pulsa-api/routes"
	"madoo-pulsa-api/swagger"
	"os"

	_ "madoo-pulsa-api/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Madoo Pulsa API
// @version 1.0
// @description API Madoo Pulsa

// @host localhost:8080
// @BasePath /

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.InitDB()

	r := gin.Default()

	routes.SetupRoutes(r, db)

	swagger.SetupSwagger(r) // // Setup Swagger documentation

	r.Run(":" + os.Getenv("PORT"))

}
