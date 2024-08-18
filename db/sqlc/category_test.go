package db

import (
	"context"
	"testing"
	"time"

	"github.com/andreanpradanaa/trendstore/util"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	arg := CreateCategoryParams{
		Name:        util.RandomString(5),
		Description: "Produk Elektronik",
	}

	category, err := testStore.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.Name, category.Name)
	require.Equal(t, arg.Description, category.Description)

	require.NotZero(t, category.ID)
	require.True(t, category.UpdatedAt.IsZero())
	require.NotZero(t, category.CreatedAt)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	category2, err := testStore.GetCategory(context.Background(), category1.Name)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.Name, category2.Name)
	require.Equal(t, category1.Description, category2.Description)
	require.WithinDuration(t, category1.CreatedAt, category2.CreatedAt, time.Second)
}

func TestUpdateCategory(t *testing.T) {
	category := createRandomCategory(t)

	newName := util.RandomString(5)
	NewDesc := util.RandomString(20)

	arg := UpdateCategoryParams{
		ID: category.ID,
		Name: pgtype.Text{
			String: newName,
			Valid:  true,
		},
		Description: pgtype.Text{
			String: NewDesc,
			Valid:  true,
		},
	}

	updatedCategory, err := testStore.UpdateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedCategory)

	require.NotEqual(t, category.Name, updatedCategory.Name)
	require.Equal(t, newName, updatedCategory.Name)
	require.NotEqual(t, category.Description, updatedCategory.Description)
	require.Equal(t, NewDesc, updatedCategory.Description)
}

func TestDeleteCategory(t *testing.T) {
	category := createRandomCategory(t)

	err := testStore.DeleteCategory(context.Background(), category.Name)
	require.NoError(t, err)

	category2, err := testStore.GetCategory(context.Background(), category.Name)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, category2)
}

func TestListCategory(t *testing.T) {
	for i := 0; i < 4; i++ {
		createRandomCategory(t)
	}

	arg := ListCategoryParams{
		Limit:  5,
		Offset: 0,
	}

	categories, err := testStore.ListCategory(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, categories, 5)

	for _, category := range categories {
		require.NotEmpty(t, category)
	}
}
