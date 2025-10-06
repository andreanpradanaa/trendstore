package repositories

import (
	"github.com/andreanpradanaa/trendstore/internal/domain/category"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) Create(category *category.Category) error {
	return r.db.Create(category).Error
}
