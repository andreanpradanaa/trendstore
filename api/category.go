package api

import (
	"net/http"

	"github.com/andreanpradanaa/trendstore/api/entity"
	db "github.com/andreanpradanaa/trendstore/db/sqlc"
	"github.com/gin-gonic/gin"
)

func (server *Server) createCategory(ctx *gin.Context) {
	var req entity.Category

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	category, err := server.store.CreateCategory(ctx, db.CreateCategoryParams{
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := db.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, res)
}

func (server *Server) updateCategory(ctx *gin.Context) {
	var 
}
