package services

import (
	"database/sql"
	"madoo-pulsa-api/models"
	"madoo-pulsa-api/repositories"
	"madoo-pulsa-api/utils"

	"github.com/google/uuid"
)

func Register(db *sql.DB, username, password string) error {
	hash, _ := utils.HashPassword(password)
	user := models.User{
		ID:       uuid.New().String(),
		Username: username,
		Password: hash,
	}
	return repositories.CreateUser(db, user)
}

func Login(db *sql.DB, username, password string) (string, string, error) {
	user, err := repositories.GetUserByUsername(db, username)
	if err != nil || !utils.CheckPasswordHash(password, user.Password) {
		return "", "", err
	}
	token, _ := utils.GenerateToken(username)
	refresh, _ := utils.GenerateRefreshToken(username)
	return token, refresh, nil
}
