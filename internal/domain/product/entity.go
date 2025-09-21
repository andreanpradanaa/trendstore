package product

import (
	"errors"
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

func (p *Product) ChangeName(name string) error {
	if len(name) < 3 {
		return errors.New("name must be at least 3 characters long")
	}
	p.Name = name
	return nil
}

func (p *Product) ChangeDescription(description string) error {
	if len(description) < 10 {
		return errors.New("description must be at least 10 characters long")
	}
	p.Description = description
	return nil
}

func (p *Product) ChangePrice(price float64) error {
	if price < 0 {
		return errors.New("price must be greater than or equal to 0")
	}
	p.Price = price
	return nil
}

func (p *Product) ChangeStock(stock int64) error {
	if stock < 0 {
		return errors.New("stock must be greater than or equal to 0")
	}
	p.Stock = stock
	return nil
}

func (p *Product) ChangeCategoryID(categoryID int64) error {
	if categoryID < 1 {
		return errors.New("category ID must be greater than or equal to 1")
	}
	p.CategoryID = categoryID
	return nil
}

func (p *Product) UpdateTimestamp() {
	p.UpdatedAt = time.Now()
}
