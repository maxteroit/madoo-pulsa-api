package migration

import (
	"database/sql"
	"fmt"
)

func UserTypeMigration(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS user_types (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name VARCHAR(255) NOT NULL,
		created_at TIMESTAMPTZ DEFAULT NOW(),
		updated_at TIMESTAMPTZ DEFAULT NOW(),
		deleted_at TIMESTAMPTZ
	);`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create user_types table: %w", err)
	}
	return nil
}