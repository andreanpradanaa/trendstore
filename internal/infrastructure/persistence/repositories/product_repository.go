package repositories

import (
	"fmt"

	"github.com/andreanpradanaa/trendstore/internal/domain/product"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) product.Repository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Create(args *product.Product) error {
	err := r.db.Create(&args).Error
	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}
	return nil
}

func (r *ProductRepository) List() ([]product.Product, error) {
	res := []product.Product{}
	err := r.db.Find(&res).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find list products: %w", err)
	}

	return res, nil
}

func (r *ProductRepository) GetByID(id int64) (*product.Product, error) {
	res := &product.Product{}
	err := r.db.First(res, id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find get product by id: %w", err)
	}

	return res, nil
}
