package handler

import (
	"net/http"

	"github.com/andreanpradanaa/trendstore/internal/product/dto"
	"github.com/andreanpradanaa/trendstore/internal/product/model"
	"github.com/andreanpradanaa/trendstore/internal/product/usecase"
	"github.com/andreanpradanaa/trendstore/internal/util"
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

func (h *ProductHandler) ListProduct(c echo.Context) error {
	res, err := h.productUsecase.ListProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.Response[any]{
			Status:  false,
			Message: "Failed to fetch products",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, util.Response[[]model.ProductListItem]{
		Status:  true,
		Message: "Products fetched successfully",
		Data:    res,
	})
}
