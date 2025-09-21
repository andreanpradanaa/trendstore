package product

import (
	"strconv"
	"strings"

	"github.com/andreanpradanaa/trendstore/internal/application/product"
	"github.com/andreanpradanaa/trendstore/internal/application/product/dto"
	"github.com/andreanpradanaa/trendstore/pkg/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

	if err := c.Bind(form); err != nil {
		return response.BadRequest(c, "invalid request payload", err)
	}

	if err := c.Validate(form); err != nil {
		return response.BadRequest(c, "validation failed", err)
	}

	if err := h.productService.Create(form); err != nil {
		return response.InternalServerError(c, "failed to create product", err)
	}

	return response.SuccessCreated(c, "product created successfully", nil)
}

func (h *Handler) List(c echo.Context) error {
	res, err := h.productService.List()
	if err != nil {
		if strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error()) {
			return response.NotFound(c, "products not found", err)
		}
		return response.InternalServerError(c, "failed to fetch products by id", err)
	}

	return response.SuccessOK(c, "products fetched successfully", res)
}

func (h *Handler) GetByID(c echo.Context) error {
	idStr := c.Param("id")
	if idStr == "" {
		return response.BadRequest(c, "product id is required", nil)
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return response.BadRequest(c, "invalid product ID", err)
	}

	res, err := h.productService.GetByID(int64(id))
	if err != nil {
		if strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error()) {
			return response.NotFound(c, "product not found", err)
		}
		return response.InternalServerError(c, "failed to fetch product by id", err)
	}

	return response.SuccessOK(c, "product fetched successfully", res)
}

func (h *Handler) Update(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return response.BadRequest(c, "product id is required", nil)
	}

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return response.BadRequest(c, "invalid product ID", err)
	}

	form := &dto.ProductUpdateRequest{}
	form.ID = idInt

	if err := c.Bind(form); err != nil {
		return response.BadRequest(c, "invalid request payload", err)
	}

	// if err := c.Validate(form); err != nil {
	// 	return response.BadRequest(c, "validation failed", err)
	// }

	if err := h.productService.UpdateProduct(form); err != nil {
		if strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error()) {
			return response.NotFound(c, "product not found", err)
		}
		return response.InternalServerError(c, "failed to update product", err)
	}

	return response.SuccessOK(c, "product updated successfully", nil)
}
