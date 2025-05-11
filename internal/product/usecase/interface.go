package usecase

import (
	"github.com/andreanpradanaa/trendstore/internal/product/dto"
	"github.com/andreanpradanaa/trendstore/internal/product/repository"
)

type ProductUsecase interface {
	CreateProduct(args *dto.ProductRequest) error
}

type productUsecase struct {
	ProductRepo repository.ProductRepository
}

func NewProductUsecase(productRepo repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		ProductRepo: productRepo,
	}
}
