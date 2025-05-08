package models

import "time"

type User struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	Token     *string    `json:"token"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
