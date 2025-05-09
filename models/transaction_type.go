package models

import "time"

type TransactionType struct {
	ID        string     `json:"id"`
	Name      string     `json:"name" binding:"required"`
	Description string     `json:"description" binding:"required"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}