package repository

import (
	"github.com/andreanpradanaa/trendstore/internal/product/model"
	"gorm.io/gorm"
)

//go:generate mockery --name=ProductRepository --output=./mocks
type ProductRepository interface {
	CreateProduct(args *model.Product) error
	ListProduct() ([]model.ProductListItem, error)

	// TODO: need to implement
	// GetProductByID(id int64) (*model.Product, error)
	// GetProductsByCategoryID(categoryID int64) ([]*Product, error)
	// GetProductsByIDs(ids []int64) ([]*Product, error)
	// UpdateProduct(args *Product) error
	// DeleteProduct(id int64) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProduct(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
