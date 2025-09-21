package repositories

import (
	"fmt"
	"time"

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

func (r *ProductRepository) Update(args *product.Product) error {
	// Buat map untuk field yang akan diupdate
	updates := make(map[string]interface{})

	// Tambahkan field yang ingin diupdate (bisa disesuaikan dengan kondisi)
	if args.Name != "" {
		updates["name"] = args.Name
	}
	if args.Description != "" {
		updates["description"] = args.Description
	}
	if args.Price != 0 {
		updates["price"] = args.Price
	}
	if args.Stock != 0 {
		updates["stock"] = args.Stock
	}
	// ... tambahkan field lainnya sesuai kebutuhan

	updates["updated_at"] = time.Now()

	if len(updates) > 0 {
		err := r.db.Model(&product.Product{}).
			Where("id = ?", args.ID).
			Updates(updates).Error

		if err != nil {
			return fmt.Errorf("failed to update product: %w", err)
		}
	}

	return nil
}

func (r *ProductRepository) Delete(id int64) error {
	err := r.db.Delete(&product.Product{}, id).Error
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	return nil
}
