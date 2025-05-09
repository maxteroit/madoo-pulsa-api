package repositories

import (
	"database/sql"
	"log"
	"madoo-pulsa-api/models"
)

func GetUserByPhoneNumber(db *sql.DB, phoneNumber string) (*models.User, error) {
	log.Println("[Repo] Mencari user dengan No.Telp", phoneNumber)

	row := db.QueryRow(`
	SELECT id, phone_number, password, 
	COALESCE(token, '') as token, 
	created_at, updated_at, deleted_at 
	FROM users 
	WHERE deleted_at IS NULL AND phone_number=$1`, phoneNumber)

	var u models.User
	err := row.Scan(&u.ID, &u.PhoneNumber, &u.Password, &u.Token, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)

	if err != nil {
		log.Println("[Repo] Gagal scan row user:", err)
		return nil, err
	}

	log.Println("[Repo] User ditemukan:", u.PhoneNumber)
	return &u, nil
}

func CreateUser(db *sql.DB, user models.User) error {
	_, err := db.Exec("INSERT INTO users (id, phone_number, password, token, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", user.ID, user.PhoneNumber, user.Password, user.Token, user.CreatedAt, user.UpdatedAt)
	return err
}

func UpdateUserToken(db *sql.DB, phoneNumber, token string) (int64, error) {
	result, err := db.Exec("UPDATE users SET token = $1 WHERE phone_number = $2", token, phoneNumber)

	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
