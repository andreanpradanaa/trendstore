package mapper

import (
	"github.com/andreanpradanaa/trendstore/internal/application/product/dto"
	"github.com/andreanpradanaa/trendstore/internal/domain/product"
)

func ToResponse(domainProduct *product.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:          domainProduct.ID,
		Name:        domainProduct.Name,
		Description: domainProduct.Description,
		Price:       domainProduct.Price,
		Stock:       domainProduct.Stock,
		CategoryID:  domainProduct.CategoryID,
		SKU:         domainProduct.SKU,
		IsActive:    domainProduct.IsActive,
		ImageURL:    domainProduct.ImageURL,
		Weight:      domainProduct.Weight,
		CreatedAt:   domainProduct.CreatedAt,
		UpdatedAt:   domainProduct.UpdatedAt,
	}
}

func ToListResponse(domainProduct *product.Product) dto.ProductListResponse {
	return dto.ProductListResponse{
		ID:          domainProduct.ID,
		Name:        domainProduct.Name,
		Description: domainProduct.Description,
		Price:       domainProduct.Price,
		Stock:       domainProduct.Stock,
		CategoryID:  domainProduct.CategoryID,
		ImageURL:    domainProduct.ImageURL,
	}
}
