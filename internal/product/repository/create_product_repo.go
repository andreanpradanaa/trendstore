package repository

import (
	"fmt"

	"github.com/andreanpradanaa/trendstore/internal/product/model"
)

func (r *productRepository) CreateProduct(args *model.Product) error {
	err := r.db.Create(&args).Error
	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}
	return nil
}
