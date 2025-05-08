package main

import (
	"log"
	"madoo-pulsa-api/config"
	"madoo-pulsa-api/routes"
	"os"

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

	r.Run(":" + os.Getenv("PORT"))
}
