package main

import (
	"log"

	"github.com/andreanpradanaa/trendstore/internal/product/handler"
	"github.com/andreanpradanaa/trendstore/internal/product/repository"
	"github.com/andreanpradanaa/trendstore/internal/product/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	// Init Connection
	dsn := "host=localhost user=root password=password dbname=trendstore port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database!")
	}

	productRepo := repository.NewProduct(db)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProductHandler(productUsecase)

	// Product Routes
	e.POST("api/v1/products", productHandler.CreateProduct)
	e.GET("api/v1/products", productHandler.ListProduct)

	// TODO: add api product
	// GET /api/v1/products/{productId}
	// GET /api/v1/products?categoryId={catId}
	// GET /api/v1/products?ids=1,2,3
	// PUT /api/v1/products/{productId}
	// DELETE /api/v1/products/{productId}

	// Start Server
	e.Logger.Fatal(e.Start(":5000"))
}
