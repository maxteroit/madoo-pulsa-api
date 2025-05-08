package models

type User struct {
	ID        string  `json:"id"`
	Username  string  `json:"username"`
	Password  string  `json:"-"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}
