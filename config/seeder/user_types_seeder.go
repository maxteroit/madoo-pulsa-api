package seeder

import (
	"database/sql"
)

func UserTypeSeeder(db *sql.DB) error {
	user_types := []map[string]interface{}{
		{
			"name":  "Reguler",
		},
		{
			"name":  "Premium",
		},
		{
			"name":  "Gold",
		},
	}

	return InsertRows(db, "user_types", user_types, []string{"name"})
}