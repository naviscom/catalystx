-- name: CreateTown :one
INSERT INTO towns (
    town_name,
    town_desc,
    district_id
) VALUES (
 $1, $2, $3
)
RETURNING *;

-- name: GetTown0 :one
SELECT * FROM towns
WHERE id = $1 LIMIT 1;

-- name: GetTown1 :one
SELECT * FROM towns
WHERE town_name = $1 LIMIT 1;

-- name: ListTowns :many
SELECT * FROM towns
WHERE district_id = $3
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTown :one
UPDATE towns
SET town_name = $2,
town_desc = $3,
district_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteTown :exec
DELETE FROM towns
WHERE id = $1;
