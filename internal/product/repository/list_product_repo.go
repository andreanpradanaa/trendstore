package repository

import (
	"fmt"

	"github.com/andreanpradanaa/trendstore/internal/product/model"
)

func (r *productRepository) ListProduct() ([]model.ProductListItem, error) {
	res := []model.ProductListItem{}
	err := r.db.Find(&res).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find list products: %w", err)
	}

	return res, nil
}
