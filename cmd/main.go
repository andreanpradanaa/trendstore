package main

import (
	"log"

	appProduct "github.com/andreanpradanaa/trendstore/internal/application/product"
	httpProduct "github.com/andreanpradanaa/trendstore/internal/infrastructure/http/product"
	repoProduct "github.com/andreanpradanaa/trendstore/internal/infrastructure/persistence/repositories"
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

	productRepo := repoProduct.NewProductRepository(db)
	productService := appProduct.NewService(productRepo)
	productHandler := httpProduct.NewHandler(productService)

	// Product Routes
	e.POST("api/v1/products", productHandler.Create)
	e.GET("api/v1/products", productHandler.List)
	e.GET("api/v1/products/:id", productHandler.GetByID)
	e.PUT("api/v1/products/:id", productHandler.Update)

	// TODO: add api product
	// GET /api/v1/products?categoryId={catId}
	// DELETE /api/v1/products/{productId}

	// Start Server
	e.Logger.Fatal(e.Start(":5000"))
}
