-- name: CreateCategory :one
INSERT INTO categories (
    name,
    description
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetCategory :one
SELECT * FROM categories
WHERE name = $1 LIMIT 1;

-- name: UpdateCategory :one
UPDATE categories
SET
    name = COALESCE(sqlc.narg(name), name),
    description = COALESCE(sqlc.narg(description), description)
WHERE
    id = sqlc.arg(id)
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE name = $1;

-- name: ListCategory :many
SELECT * FROM categories
ORDER BY id
LIMIT $1
OFFSET $2;