package usecase

import "github.com/andreanpradanaa/trendstore/internal/product/model"

func (u *productUsecase) ListProduct() ([]model.ProductListItem, error) {
	res, err := u.ProductRepo.ListProduct()
	if err != nil {
		return nil, err
	}
	return res, nil
}
