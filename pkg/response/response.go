package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func Success(c echo.Context, statusCode int, msg string, data any) error {
	return c.JSON(statusCode, Response[any]{
		Status:  true,
		Message: msg,
		Data:    data,
	})
}

func Error(c echo.Context, statusCode int, msg string, err error) error {
	return c.JSON(statusCode, Response[any]{
		Status:  false,
		Message: msg,
		Error:   err.Error(),
	})
}

func SuccessOK(c echo.Context, msg string, data any) error {
	return Success(c, http.StatusOK, msg, data)
}

func SuccessCreated(c echo.Context, msg string, data any) error {
	return Success(c, http.StatusCreated, msg, data)
}

func NotFound(c echo.Context, msg string, err error) error {
	return Error(c, http.StatusNotFound, msg, err)
}

func BadRequest(c echo.Context, msg string, err error) error {
	return Error(c, http.StatusBadRequest, msg, err)
}

func InternalServerError(c echo.Context, msg string, err error) error {
	return Error(c, http.StatusInternalServerError, msg, err)
}
