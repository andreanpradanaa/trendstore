package usecase

import (
	"time"

	"github.com/andreanpradanaa/trendstore/internal/product/dto"
	"github.com/andreanpradanaa/trendstore/internal/product/model"
)

func (p *productUsecase) CreateProduct(args *dto.ProductRequest) error {
	product := model.Product{
		Name:        args.Name,
		Description: args.Description,
		Price:       args.Price,
		Stock:       args.Stock,
		CategoryID:  args.CategoryID,
		CreatedAt:   time.Now(),
	}

	err := p.ProductRepo.CreateProduct(&product)
	if err != nil {
		return err
	}

	return nil
}
