package seeder

import (
	"database/sql"
)

func TransactionTypeSeeder(db *sql.DB) error {
	transaction_types := []map[string]interface{}{
		{
			"name":  "Pulsa",
			"description": "Transaksi Pulsa",
		},
		// {
		// 	"name":  "Top Up",
		// 	"description": "Topup balance atau deposit user",
		// },
		// {
		// 	"name":  "PPOB",
		// 	"description": "Transaksi Payment PPOB",
		// },
	}

	return InsertRows(db, "transaction_types", transaction_types, []string{"name", "description"})
}