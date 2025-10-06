package category

import (
	"github.com/andreanpradanaa/trendstore/internal/application/category"
	"github.com/andreanpradanaa/trendstore/internal/application/category/dto"
	"github.com/andreanpradanaa/trendstore/pkg/response"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	serviceCategory category.Service
}

func NewHandler(serviceCategory category.Service) *Handler {
	return &Handler{
		serviceCategory: serviceCategory,
	}
}

func (h *Handler) Create(c echo.Context) error {
	var (
		form = &dto.CreateCategoryRequest{}
		err  error
	)

	if err = c.Bind(form); err != nil {
		return response.BadRequest(c, "invalid request payload", err)
	}

	if err = c.Validate(form); err != nil {
		return response.BadRequest(c, "validation failed", err)
	}

	err = h.serviceCategory.Create(form)
	if err != nil {
		return response.InternalServerError(c, "failed to create category", err)
	}

	return response.SuccessCreated(c, "successfully create new category", nil)
}
