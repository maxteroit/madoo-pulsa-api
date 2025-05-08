package repositories

import (
	"database/sql"
	"madoo-pulsa-api/models"
)

func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
	row := db.QueryRow("SELECT id, username, password, created_at, updated_at, deleted_at FROM users WHERE deleted_at IS NULL AND username=$1", username)
	var u models.User
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func CreateUser(db *sql.DB, user models.User) error {
	_, err := db.Exec("INSERT INTO users (id, username, password) VALUES ($1, $2, $3)", user.ID, user.Username, user.Password)
	return err
}
