package models

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Qty         int     `json:"qty"`
	ExpiredDate string  `json:"expired_date"`
	Image       string  `json:"image"`
	DeletedAt   *string `json:"deleted_at,omitempty"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
