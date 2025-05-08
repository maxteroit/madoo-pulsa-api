package models

import "time"

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Price       float64   `json:"price" binding:"required"`
	Qty         int       `json:"qty" binding:"required"`
	ExpiredDate time.Time `json:"expired_date" binding:"required"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}