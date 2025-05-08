package services

import (
	"database/sql"
	"log"
	"madoo-pulsa-api/models"
)

// Create inserts a new product into the database.
func Create(db *sql.DB, product models.Product) error {
	query := `INSERT INTO products (name, price, qty, expired_date, image) 
			  VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(query, product.Name, product.Price, product.Qty, product.ExpiredDate, product.Image)
	if err != nil {
		log.Println("Error creating product:", err)
		return err
	}
	return nil
}

// GetAll retrieves all products from the database.
func GetAll(db *sql.DB) ([]models.Product, error) {
	rows, err := db.Query(`SELECT id, name, price, qty, expired_date, image FROM products WHERE deleted_at IS NULL`)
	if err != nil {
		log.Println("Error fetching products:", err)
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Qty, &product.ExpiredDate, &product.Image); err != nil {
			log.Println("Error scanning product:", err)
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// GetByID retrieves a product by its ID from the database.
func GetByID(db *sql.DB, id string) (*models.Product, error) {
	var product models.Product
	query := `SELECT id, name, price, qty, expired_date, image FROM products WHERE id = $1 AND deleted_at IS NULL`
	err := db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Qty, &product.ExpiredDate, &product.Image)
	if err != nil {
		log.Println("Error fetching product by ID:", err)
		return nil, err
	}
	return &product, nil
}

// Update updates the details of a product in the database.
func Update(db *sql.DB, id string, product models.Product) error {
	query := `UPDATE products SET name = $1, price = $2, qty = $3, expired_date = $4, image = $5 WHERE id = $6`
	_, err := db.Exec(query, product.Name, product.Price, product.Qty, product.ExpiredDate, product.Image, id)
	if err != nil {
		log.Println("Error updating product:", err)
		return err
	}
	return nil
}

// SoftDelete marks a product as deleted without actually removing it from the database.
func SoftDelete(db *sql.DB, id string) error {
	query := `UPDATE products SET deleted_at = NOW() WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Println("Error soft deleting product:", err)
		return err
	}
	return nil
}

// UpdateProductImage updates the image path for a product in the database.
func UpdateProductImage(db *sql.DB, id string, imagePath string) error {
	query := `UPDATE products SET image = $1 WHERE id = $2`
	_, err := db.Exec(query, imagePath, id)
	if err != nil {
		log.Println("Error updating product image:", err)
		return err
	}
	return nil
}
