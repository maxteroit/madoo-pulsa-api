package migration

import (
	"database/sql"
	"fmt"
)

func TermsConditionUserMigration(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS terms_conditions_users (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		title VARCHAR(255) NOT NULL,
		description VARCHAR(255),
		created_at TIMESTAMPTZ DEFAULT NOW(),
		updated_at TIMESTAMPTZ DEFAULT NOW(),
		deleted_at TIMESTAMPTZ
	);`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create terms_conditions_users table: %w", err)
	}
	return nil
}