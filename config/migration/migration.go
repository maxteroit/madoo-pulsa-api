package migration

import "database/sql"

// Tipe alias untuk fungsi migrasi
type Migration func(*sql.DB) error

func RunMigrations(db *sql.DB) error {
	migrations := []Migration{
		// UserMigration,
		// ProductMigration,
		TransactionTypeMigration,
		TermsConditionUserMigration,
		UserTypeMigration,
	}

	for _, m := range migrations {
		if err := m(db); err != nil {
			return err
		}
	}
	return nil
}