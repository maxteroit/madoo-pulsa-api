package repositories

import (
	"database/sql"
	"madoo-pulsa-api/models"
)

func CreateProduct(db *sql.DB, p models.Product) error {
	_, err := db.Exec("INSERT INTO products (id, name, price, qty, expired_date, image) VALUES ($1,$2,$3,$4,$5,$6)", p.ID, p.Name, p.Price, p.Qty, p.ExpiredDate, p.Image)
	return err
}

func GetProducts(db *sql.DB) ([]models.Product, error) {
	rows, err := db.Query("SELECT id, name, price, qty, expired_date, image, deleted_at, created_at, updated_at FROM products WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Qty, &p.ExpiredDate, &p.Image, &p.DeletedAt, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
