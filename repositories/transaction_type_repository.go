package repositories

import (
	"database/sql"
	"log"
	"madoo-pulsa-api/models"
	"time"
)

func IsTransactionTypeNameExists(db *sql.DB, name string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM transaction_types WHERE name = $1 AND deleted_at IS NULL)", name).Scan(&exists)
	if err != nil {
		log.Println("Error checking if transaction type name exists:", err)
		return false, err
	}
	return exists, nil
}
func IsTransactionTypeNameExistsExcludeID(db *sql.DB, name string, id string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM transaction_types WHERE name = $1 AND id != $2 AND deleted_at IS NULL)", name, id).Scan(&exists)
	if err != nil {
		log.Println("Error checking if transaction type name exists excluding ID:", err)
		return false, err
	}
	return exists, nil
}
func CreateTransactionType(db *sql.DB, transactionType models.TransactionType) error {
	_, err := db.Exec(`INSERT INTO transaction_types (name, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4)`,
		transactionType.Name, transactionType.Description, time.Now(), time.Now())
	if err != nil {
		log.Println("Error creating transaction type:", err)
		return err
	}
	return nil
}
func FetchAllTransactionTypes(db *sql.DB) ([]models.TransactionType, error) {
	rows, err := db.Query("SELECT id, name, description, created_at, updated_at FROM transaction_types WHERE deleted_at IS NULL")
	if err != nil {
		log.Println("Error fetching transaction types:", err)
		return nil, err
	}
	defer rows.Close()

	var transactionTypes []models.TransactionType
	for rows.Next() {
		var t models.TransactionType
		err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			log.Println("Error scanning transaction type:", err)
			return nil, err
		}
		transactionTypes = append(transactionTypes, t)
	}
	return transactionTypes, nil
}
func FetchTransactionTypeByID(db *sql.DB, id string) (models.TransactionType, error) {
	var t models.TransactionType
	err := db.QueryRow("SELECT id, name, description, created_at, updated_at FROM transaction_types WHERE id = $1 AND deleted_at IS NULL", id).
		Scan(&t.ID, &t.Name, &t.Description, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		log.Println("Error fetching transaction type by ID:", err)
		return t, err
	}
	return t, nil
}
func UpdateTransactionType(db *sql.DB, id string, transactionType models.TransactionType) error {	
	_, err := db.Exec(`UPDATE transaction_types SET name=$1, description=$2, updated_at=$3 WHERE id=$4 AND deleted_at IS NULL`,
		transactionType.Name, transactionType.Description, time.Now(), id)
	if err != nil {
		log.Println("Error updating transaction type:", err)
		return err
	}
	return nil
}
func SoftDeleteTransactionType(db *sql.DB, id string) error {
	_, err := db.Exec(`UPDATE transaction_types SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL`, time.Now(), id)
	if err != nil {
		log.Println("Error soft deleting transaction type:", err)
		return err
	}
	return nil
}
func RestoreTransactionType(db *sql.DB, id string) error {
	_, err := db.Exec(`UPDATE transaction_types SET deleted_at = NULL WHERE id = $1`, id)
	if err != nil {
		log.Println("Error restoring transaction type:", err)
		return err
	}
	return nil
}
func HardDeleteTransactionType(db *sql.DB, id string) error {
	_, err := db.Exec(`DELETE FROM transaction_types WHERE id = $1`, id)
	if err != nil {
		log.Println("Error hard deleting transaction type:", err)
		return err
	}
	return nil
}
func FetchAllDeletedTransactionTypes(db *sql.DB) ([]models.TransactionType, error) {
	rows, err := db.Query("SELECT id, name, description, created_at, updated_at FROM transaction_types WHERE deleted_at IS NOT NULL")
	if err != nil {
		log.Println("Error fetching deleted transaction types:", err)
		return nil, err
	}
	defer rows.Close()

	var transactionTypes []models.TransactionType
	for rows.Next() {
		var t models.TransactionType
		err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			log.Println("Error scanning deleted transaction type:", err)
			return nil, err
		}
		transactionTypes = append(transactionTypes, t)
	}
	return transactionTypes, nil
}
func FetchTransactionTypeByName(db *sql.DB, name string) (models.TransactionType, error) {
	var t models.TransactionType
	err := db.QueryRow("SELECT id, name, description, created_at, updated_at FROM transaction_types WHERE name = $1 AND deleted_at IS NULL", name).
		Scan(&t.ID, &t.Name, &t.Description, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		log.Println("Error fetching transaction type by name:", err)
		return t, err
	}
	return t, nil
}