package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/andreanpradanaa/trendstore/util"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T) Product {
	category := createRandomCategory(t)

	price, _ := util.RandomPrice(10, 100)
	arg := CreateProductParams{
		Name:        util.RandomString(10),
		Description: util.RandomDescription(10),
		Price:       price,
		Stock:       util.RandomStock(),
		CategoryID:  category.ID,
	}

	product, err := testStore.CreateProduct(context.Background(), arg)

	fmt.Println("stock", product.Stock)

	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Description, product.Description)
	require.Equal(t, arg.Stock, product.Stock)
	require.Equal(t, arg.CategoryID, product.CategoryID)

	require.NotZero(t, product.ID)
	require.NotZero(t, product.CreatedAt)
	require.True(t, category.UpdatedAt.IsZero())

	return product
}

func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)
}

func TestGetProduct(t *testing.T) {
	product := createRandomProduct(t)

	product2, err := testStore.GetProduct(context.Background(), product.Name)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product.Name, product2.Name)
	require.Equal(t, product.Description, product2.Description)
	require.Equal(t, product.Price, product2.Price)
	require.Equal(t, product.Stock, product2.Stock)
}

func TestDeleteProduct(t *testing.T) {
	product := createRandomProduct(t)

	err := testStore.DeleteProduct(context.Background(), product.Name)
	require.NoError(t, err)

	product2, err := testStore.GetProduct(context.Background(), product.Name)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, product2)
}

func TestUpdateProduct(t *testing.T) {
	product := createRandomProduct(t)

	newName := util.RandomString(10)
	newDesc := util.RandomDescription(10)
	updateProduct, err := testStore.UpdateProduct(context.Background(), UpdateProductParams{
		Name: pgtype.Text{
			String: newName,
			Valid:  true,
		},
		Description: pgtype.Text{
			String: newDesc,
			Valid:  true,
		},
		// Price:      pgtype.Numeric{},
		// Stock:      pgtype.Int4{},
		// CategoryID: pgtype.Int8{},
		ID: product.ID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, updateProduct)

	require.Equal(t, newName, updateProduct.Name)
	require.NotEqual(t, product.Name, updateProduct.Name)
	require.Equal(t, product.ID, updateProduct.ID)
}

// func TestListProduct(t *testing.T) {
// 	// category := createRandomCategory(t)
// 	// product := createRandomProduct(t)

// 	var lastProduct Product
// 	for i := 0; i < 10; i++ {
// 		lastProduct = createRandomProduct(t)
// 	}

// 	listProduct, err := testStore.ListProduct(context.Background(), ListProductParams{
// 		CategoryID: lastProduct.CategoryID,
// 		Limit:      5,
// 		Offset:     0,
// 	})

// 	require.NoError(t, err)
// 	require.Len(t, listProduct, 5)

// 	// require.Equal(t, lastProduct., lastProduct.CategoryID)
// }
