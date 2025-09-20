package product

import (
	"time"

	"github.com/andreanpradanaa/trendstore/internal/application/product/dto"
	"github.com/andreanpradanaa/trendstore/internal/application/product/mapper"
	"github.com/andreanpradanaa/trendstore/internal/domain/product"
)

type Service struct {
	productRepo product.Repository
}

func NewService(productRepo product.Repository) Service {
	return Service{
		productRepo: productRepo,
	}
}

func (s *Service) Create(args *dto.ProductRequest) error {
	product := product.Product{
		Name:        args.Name,
		Description: args.Description,
		Price:       args.Price,
		Stock:       args.Stock,
		CategoryID:  args.CategoryID,
		CreatedAt:   time.Now(),
	}

	err := s.productRepo.Create(&product)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) List() ([]dto.ProductListResponse, error) {
	products, err := s.productRepo.List()
	if err != nil {
		return nil, err
	}

	response := []dto.ProductListResponse{}
	for _, product := range products {
		response = append(response, mapper.ToListResponse(&product))
	}

	return response, nil
}

func (s *Service) GetByID(id int64) (*dto.ProductResponse, error) {
	products, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	response := mapper.ToResponse(products)
	return &response, nil
}
