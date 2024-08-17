-- name: CreateCategory :one
INSERT INTO categories (
    id,
    name,
    description,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5
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
    name = sqlc.arg(name)
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE name = $1;