package product

import (
	"net/http"

	"github.com/andreanpradanaa/trendstore/internal/application/product"
	"github.com/andreanpradanaa/trendstore/internal/application/product/dto"
	"github.com/andreanpradanaa/trendstore/pkg/response"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	productService product.Service
}

func NewHandler(productService product.Service) *Handler {
	return &Handler{
		productService: productService,
	}
}

func (h *Handler) Create(c echo.Context) error {
	form := &dto.ProductRequest{}

	err := c.Bind(form)
	if err != nil {
		return err
	}

	err = h.productService.Create(form)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"status":  "ok",
		"message": "successfully create product",
	})
}

func (h *Handler) List(c echo.Context) error {
	res, err := h.productService.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response[any]{
			Status:  false,
			Message: "Failed to fetch products",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response[[]dto.ProductListResponse]{
		Status:  true,
		Message: "Products fetched successfully",
		Data:    res,
	})
}
