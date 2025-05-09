package seeder

import (
	"database/sql"
	"fmt"
	"strings"
)

type SeederFunc func(*sql.DB) error

func RunSeeders(db *sql.DB) error {
	seeders := []SeederFunc{
		TransactionTypeSeeder,
		UserTypeSeeder,
	}

	for _, seed := range seeders {
		if err := seed(db); err != nil {
			return err
		}
	}
	return nil
}

func InsertRows(db *sql.DB, table string, rows []map[string]interface{}, uniqueKeys []string) error {
	for _, row := range rows {
		whereParts := []string{}
		whereValues := []interface{}{}
		i := 1
		for _, key := range uniqueKeys {
			whereParts = append(whereParts, fmt.Sprintf("%s = $%d", key, i))
			whereValues = append(whereValues, row[key])
			i++
		}
		whereClause := strings.Join(whereParts, " AND ")
		checkQuery := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s);", table, whereClause)

		var exists bool
		err := db.QueryRow(checkQuery, whereValues...).Scan(&exists)
		if err != nil {
			return fmt.Errorf("failed to check existence in %s: %w", table, err)
		}

		if exists {
			fmt.Printf("Row already exists in %s, skipping insert\n", table)
			continue // skip insert
		}

		// Build insert query
		columns := []string{}
		values := []interface{}{}
		placeholders := []string{}
		j := 1
		for key, val := range row {
			columns = append(columns, key)
			values = append(values, val)
			placeholders = append(placeholders, fmt.Sprintf("$%d", j))
			j++
		}

		insertQuery := fmt.Sprintf(
			"INSERT INTO %s (%s) VALUES (%s);",
			table,
			strings.Join(columns, ", "),
			strings.Join(placeholders, ", "),
		)

		if _, err := db.Exec(insertQuery, values...); err != nil {
			return fmt.Errorf("failed to insert into %s: %w", table, err)
		}
	}

	return nil
}