package product

import (
	"time"
)

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int64     `json:"stock"`
	CategoryID  int64     `json:"category_id"`
	SKU         string    `json:"sku"`
	IsActive    bool      `json:"is_active"`
	ImageURL    string    `json:"image_url"`
	Weight      float64   `json:"weight"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *Product) TableName() string {
	return "products"
}
