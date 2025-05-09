package main

import (
	"log"
	"madoo-pulsa-api/config"
	"madoo-pulsa-api/config/migration"
	"madoo-pulsa-api/config/seeder"
	"os"

	"github.com/joho/godotenv"
)

// @title cmd bash
// @version 1.0
// @description Skeleton Madoo Backend

// @host localhost:8080
// @BasePath /

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.InitDB()

	if err := migration.RunMigrations(db); err != nil {
		log.Fatalf("Migration error: %v", err)
	}

	if os.Getenv("SEED") == "true" {
		if err := seeder.RunSeeders(db); err != nil {
			log.Fatalf("Seeder failed: %v", err)
		}
		log.Println("Seeder done ✅")
	}

	log.Println("All tables created successfully.")

}
