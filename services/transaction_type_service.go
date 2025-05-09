package services

import (
	"database/sql"
	"log"
	"madoo-pulsa-api/models"
	"madoo-pulsa-api/repositories"
)

func CreateTransactionType(db *sql.DB, transactionType models.TransactionType) error {
	err := repositories.CreateTransactionType(db, transactionType)
	if err != nil {
		log.Println("Error creating transaction type:", err)
		return err
	}
	return nil
}

func GetAllTransactionTypes(db *sql.DB) ([]models.TransactionType, error) {
	transactionTypes, err := repositories.FetchAllTransactionTypes(db)
	if err != nil {
		log.Println("Error fetching transaction types:", err)
		return nil, err
	}
	return transactionTypes, nil
}