package repositories

import (
	"database/sql"
	"log"
	"madoo-pulsa-api/models"
)

func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
	log.Println("[Repo] Mencari user dengan username:", username)

	row := db.QueryRow(`
	SELECT id, username, password, 
	COALESCE(token, '') as token, 
	created_at, updated_at, deleted_at 
	FROM users 
	WHERE deleted_at IS NULL AND username=$1`, username)

	var u models.User
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Token, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)

	if err != nil {
		log.Println("[Repo] Gagal scan row user:", err)
		return nil, err
	}

	log.Println("[Repo] User ditemukan:", u.Username)
	return &u, nil
}

func CreateUser(db *sql.DB, user models.User) error {
	_, err := db.Exec("INSERT INTO users (id, username, password, token, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", user.ID, user.Username, user.Password, user.Token, user.CreatedAt, user.UpdatedAt)
	return err
}
