package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"madoo-pulsa-api/models"
	"madoo-pulsa-api/repositories"
	"madoo-pulsa-api/utils"
	"time"

	"github.com/google/uuid"
)

func Register(db *sql.DB, username, password string) error {
	// Hash password
	hash, err := utils.HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Generate token
	token, err := utils.GenerateToken(username)
	if err != nil {
		return fmt.Errorf("failed to generate token: %w", err)
	}

	// Get current time and convert to pointer
	now := time.Now()

	// Create user model
	user := models.User{
		ID:        uuid.New().String(),
		Username:  username,
		Password:  hash,
		Token:     utils.StringPtr(token),
		CreatedAt: utils.TimePtr(now),
		UpdatedAt: utils.TimePtr(now),
	}

	// Save to database
	return repositories.CreateUser(db, user)
}

func Login(db *sql.DB, username, password string) (string, string, error) {
	user, err := repositories.GetUserByUsername(db, username)
	if err != nil {
		log.Println("[Login] Error ambil user dari DB:", err)
		return "", "", err
	}

	log.Println("[Login] Username ditemukan:", user.Username)
	log.Println("[Login] Password hash tersimpan:", user.Password)

	token, err := utils.GenerateToken(username)
	if err != nil {
		log.Println("[Login] Gagal generate token:", err)
		return "", "", err
	}

	refresh, err := utils.GenerateRefreshToken(username)
	if err != nil {
		log.Println("[Login] Gagal generate refresh token:", err)
		return "", "", err
	}

	// Update token to database
	rowsAffected, err := repositories.UpdateUserToken(db, username, token)
	if err != nil {
		log.Println("[Login] Gagal update token ke database:", err)
		return "", "", err
	}
	if rowsAffected == 0 {
		log.Println("[Login] Tidak ada baris yang terupdate (username mungkin tidak ditemukan)")
		return "", "", errors.New("failed to update token")
	}

	log.Println("[Login] Token berhasil dibuat dan disimpan ke DB")
	return token, refresh, nil
}
