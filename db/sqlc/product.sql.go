// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: product.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (
    name, description, price, stock, category_id
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, name, description, price, stock, category_id, created_at, updated_at
`

type CreateProductParams struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       pgtype.Numeric `json:"price"`
	Stock       int32          `json:"stock"`
	CategoryID  int64          `json:"category_id"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Stock,
		arg.CategoryID,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Stock,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE name = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, name string) error {
	_, err := q.db.Exec(ctx, deleteProduct, name)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT id, name, description, price, stock, category_id, created_at, updated_at FROM products
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, name string) (Product, error) {
	row := q.db.QueryRow(ctx, getProduct, name)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Stock,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listProduct = `-- name: ListProduct :many
SELECT id, name, description, price, stock, category_id, created_at, updated_at FROM products
WHERE category_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListProductParams struct {
	CategoryID int64 `json:"category_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) ListProduct(ctx context.Context, arg ListProductParams) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProduct, arg.CategoryID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Stock,
			&i.CategoryID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE products 
SET
    name = COALESCE($1, name),
    description = COALESCE($2, description),
    price = COALESCE($3, price),
    stock = COALESCE($4, stock)
WHERE
    id = $5
RETURNING id, name, description, price, stock, category_id, created_at, updated_at
`

type UpdateProductParams struct {
	Name        pgtype.Text    `json:"name"`
	Description pgtype.Text    `json:"description"`
	Price       pgtype.Numeric `json:"price"`
	Stock       pgtype.Int4    `json:"stock"`
	ID          int64          `json:"id"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, updateProduct,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Stock,
		arg.ID,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Stock,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
