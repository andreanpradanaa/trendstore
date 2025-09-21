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

func (s *Service) UpdateProduct(args *dto.ProductUpdateRequest) error {
	exisitingProduct, err := s.productRepo.GetByID(args.ID)
	if err != nil {
		return err
	}

	if args.Name != "" {
		if err := exisitingProduct.ChangeName(args.Name); err != nil {
			return err
		}
	}

	if args.Description != "" {
		if err := exisitingProduct.ChangeDescription(args.Description); err != nil {
			return err
		}
	}

	if args.Price != 0 {
		if err := exisitingProduct.ChangePrice(args.Price); err != nil {
			return err
		}
	}

	if args.Stock != 0 {
		if err := exisitingProduct.ChangeStock(args.Stock); err != nil {
			return err
		}
	}

	if args.CategoryID != 0 {
		if err := exisitingProduct.ChangeCategoryID(args.CategoryID); err != nil {
			return err
		}
	}

	product := *exisitingProduct
	product.UpdatedAt = time.Now()

	err = s.productRepo.Update(&product)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Delete(id int64) error {
	_, err := s.productRepo.GetByID(id)
	if err != nil {
		return err
	}

	err = s.productRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
