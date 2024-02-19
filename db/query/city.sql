-- name: CreateCity :one
INSERT INTO cities (
    city_name,
    city_desc,
    state_id
) VALUES (
 $1, $2, $3
)
RETURNING *;

-- name: GetCity0 :one
SELECT * FROM cities
WHERE id = $1 LIMIT 1;

-- name: GetCity1 :one
SELECT * FROM cities
WHERE city_name = $1 LIMIT 1;

-- name: ListCities :many
SELECT * FROM cities
WHERE state_id = $3
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateCity :one
UPDATE cities
SET city_name = $2,
city_desc = $3,
state_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteCity :exec
DELETE FROM cities
WHERE id = $1;
