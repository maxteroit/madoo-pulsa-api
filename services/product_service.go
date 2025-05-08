package services

import (
	"database/sql"
	"madoo-pulsa-api/models"
	"madoo-pulsa-api/repositories"

	"github.com/google/uuid"
)

func Create(db *sql.DB, p models.Product) error {
	p.ID = uuid.New().String()
	return repositories.CreateProduct(db, p)
}

func GetAll(db *sql.DB) ([]models.Product, error) {
	return repositories.GetProducts(db)
}
