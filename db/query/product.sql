-- name: CreateProduct :one
INSERT INTO products (
    name, description, price, stock, category_id
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE name = $1 LIMIT 1;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE name = $1;

-- name: UpdateProduct :one
UPDATE products 
SET
    name = COALESCE(sqlc.narg(name), name),
    description = COALESCE(sqlc.narg(description), description),
    price = COALESCE(sqlc.narg(price), price),
    stock = COALESCE(sqlc.narg(stock), stock)
WHERE
    id = sqlc.arg(id)
RETURNING *;

-- name: ListProduct :many
SELECT * FROM products
WHERE category_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;