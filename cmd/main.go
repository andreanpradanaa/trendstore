package main

import (
	"log"
	"net/http"

	appCategory "github.com/andreanpradanaa/trendstore/internal/application/category"
	appProduct "github.com/andreanpradanaa/trendstore/internal/application/product"
	httpCategory "github.com/andreanpradanaa/trendstore/internal/infrastructure/http/category"
	httpProduct "github.com/andreanpradanaa/trendstore/internal/infrastructure/http/product"
	"github.com/andreanpradanaa/trendstore/internal/infrastructure/persistence/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Validator struct {
	validate *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validate.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()

	// Set Validator
	e.Validator = &Validator{validate: validator.New()}

	// Middleware
	e.Use(middleware.CORS())

	// Init Connection
	dsn := "host=localhost user=root password=password dbname=trendstore port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database!")
	}

	// Register
	productRepo := repositories.NewProductRepository(db)
	productService := appProduct.NewService(productRepo)
	productHandler := httpProduct.NewHandler(productService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := appCategory.NewService(categoryRepo)
	categoryHandler := httpCategory.NewHandler(*categoryService)

	v1 := e.Group("api/v1/")

	// Product Routes
	v1.POST("products", productHandler.Create)
	v1.GET("products", productHandler.List)
	v1.GET("products/:id", productHandler.GetByID)
	v1.PUT("products/:id", productHandler.Update)
	v1.DELETE("products/:id", productHandler.Delete)

	// Category Routes
	v1.POST("category", categoryHandler.Create)

	// Start Server
	e.Logger.Fatal(e.Start(":5000"))
}
