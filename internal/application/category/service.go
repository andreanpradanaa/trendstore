package category

import (
	"github.com/andreanpradanaa/trendstore/internal/application/category/dto"
	"github.com/andreanpradanaa/trendstore/internal/domain/category"
)

type Service struct {
	categoryRepo category.Repository
}

func NewService(categoryRepo category.Repository) *Service {
	return &Service{
		categoryRepo: categoryRepo,
	}
}

func (s *Service) Create(args *dto.CreateCategoryRequest) error {
	category := &category.Category{
		Name:        args.Name,
		Description: args.Description,
	}

	err := s.categoryRepo.Create(category)
	if err != nil {
		return err
	}

	return nil
}
