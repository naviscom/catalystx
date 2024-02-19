-- name: CreateArea :one
INSERT INTO areas (
    area_name,
    area_desc
) VALUES (
 $1, $2
)
RETURNING *;

-- name: GetArea0 :one
SELECT * FROM areas
WHERE id = $1 LIMIT 1;

-- name: GetArea1 :one
SELECT * FROM areas
WHERE area_name = $1 LIMIT 1;

-- name: ListAreas :many
SELECT * FROM areas
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateArea :one
UPDATE areas
SET area_name = $2,
area_desc = $3
WHERE id = $1
RETURNING *;

-- name: DeleteArea :exec
DELETE FROM areas
WHERE id = $1;
