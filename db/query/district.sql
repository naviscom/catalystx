-- name: CreateDistrict :one
INSERT INTO districts (
    district_name,
    district_desc,
    city_id
) VALUES (
 $1, $2, $3
)
RETURNING *;

-- name: GetDistrict0 :one
SELECT * FROM districts
WHERE id = $1 LIMIT 1;

-- name: GetDistrict1 :one
SELECT * FROM districts
WHERE district_name = $1 LIMIT 1;

-- name: ListDistricts :many
SELECT * FROM districts
WHERE city_id = $3
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateDistrict :one
UPDATE districts
SET district_name = $2,
district_desc = $3,
city_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteDistrict :exec
DELETE FROM districts
WHERE id = $1;
