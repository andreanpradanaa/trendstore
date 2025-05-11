package handler

import (
	"net/http"

	"github.com/andreanpradanaa/trendstore/internal/product/dto"
	"github.com/andreanpradanaa/trendstore/internal/product/usecase"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUsecase: productUsecase,
	}
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	form := &dto.ProductRequest{}

	err := c.Bind(form)
	if err != nil {
		return err
	}

	err = h.productUsecase.CreateProduct(form)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"status":  "ok",
		"message": "successfully create product",
	})
}