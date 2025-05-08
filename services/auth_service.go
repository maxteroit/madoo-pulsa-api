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

	if !utils.CheckPasswordHash(password, user.Password) {
		log.Println("[Login] Password tidak cocok")
		return "", "", errors.New("invalid password")
	}

	token, _ := utils.GenerateToken(username)
	refresh, _ := utils.GenerateRefreshToken(username)
	log.Println("[Login] Token berhasil dibuat")

	return token, refresh, nil
}
