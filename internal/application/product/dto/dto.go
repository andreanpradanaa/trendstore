package dto

import "time"

type ProductRequest struct {
	Name        string  `json:"name" validate:"required,min=3"`
	Description string  `json:"description" validate:"required,min=10"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int64   `json:"stock" validate:"required,gte=0"`
	CategoryID  int64   `json:"category_id" validate:"required,gt=0"`
}

type ProductUpdateRequest struct {
	ID          int64   `json:"id" validate:"required"`
	Name        string  `json:"name,omitempty" validate:"omitempty,min=3"`
	Description string  `json:"description,omitempty" validate:"omitempty,min=10"`
	Price       float64 `json:"price,omitempty" validate:"omitempty,gt=0"`
	Stock       int64   `json:"stock,omitempty" validate:"omitempty,gte=0"`
	CategoryID  int64   `json:"category_id,omitempty" validate:"omitempty,gt=0"`
}

type ProductResponse struct {
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

type ProductListResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int64   `json:"stock"`
	CategoryID  int64   `json:"category_id"`
	ImageURL    string  `json:"image_url"`
}
