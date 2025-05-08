package repositories

import (
	"database/sql"
	"log"
	"madoo-pulsa-api/models"
	"time"
)

func IsProductNameExists(db *sql.DB, name string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE name = $1 AND deleted_at IS NULL)", name).Scan(&exists)
	return exists, err
}

func IsProductNameExistsExcludeID(db *sql.DB, name string, id string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE name = $1 AND id != $2 AND deleted_at IS NULL)", name, id).Scan(&exists)
	return exists, err
}

func CreateProduct(db *sql.DB, product models.Product) error {
	_, err := db.Exec(`INSERT INTO products (name, price, qty, expired_date, image, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())`,
		product.Name, product.Price, product.Qty, product.ExpiredDate, product.Image)
	return err
}

func FetchAllProducts(db *sql.DB) ([]models.Product, error) {
	rows, err := db.Query("SELECT id, name, price, qty, expired_date, image, created_at, updated_at FROM products WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Qty, &p.ExpiredDate, &p.Image, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func FetchProductByID(db *sql.DB, id string) (models.Product, error) {
	var p models.Product
	err := db.QueryRow("SELECT id, name, price, qty, expired_date, image, created_at, updated_at FROM products WHERE id = $1 AND deleted_at IS NULL", id).
		Scan(&p.ID, &p.Name, &p.Price, &p.Qty, &p.ExpiredDate, &p.Image, &p.CreatedAt, &p.UpdatedAt)
	return p, err
}

func UpdateProduct(db *sql.DB, id string, product models.Product) error {
	_, err := db.Exec(`UPDATE products SET name=$1, price=$2, qty=$3, expired_date=$4, image=$5, updated_at=NOW() WHERE id=$6 AND deleted_at IS NULL`,
		product.Name, product.Price, product.Qty, product.ExpiredDate, product.Image, id)
	return err
}

func SoftDeleteProduct(db *sql.DB, id string) error {
	_, err := db.Exec(`UPDATE products SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL`, time.Now(), id)
	return err
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
