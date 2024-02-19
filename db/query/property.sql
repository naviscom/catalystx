-- name: CreateProperty :one
INSERT INTO properties (
    property_name,
    lat,
    long,
    block_id
) VALUES (
 $1, $2, $3, $4
)
RETURNING *;

-- name: GetProperty0 :one
SELECT * FROM properties
WHERE id = $1 LIMIT 1;

-- name: GetProperty1 :one
SELECT * FROM properties
WHERE property_name = $1 LIMIT 1;

-- name: ListProperties :many
SELECT * FROM properties
WHERE block_id = $3
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateProperty :one
UPDATE properties
SET property_name = $2,
lat = $3,
long = $4,
block_id = $5
WHERE id = $1
RETURNING *;

-- name: DeleteProperty :exec
DELETE FROM properties
WHERE id = $1;
