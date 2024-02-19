-- name: CreateContinent :one
INSERT INTO continents (
    continent_name,
    continent_desc
) VALUES (
 $1, $2
)
RETURNING *;

-- name: GetContinent0 :one
SELECT * FROM continents
WHERE id = $1 LIMIT 1;

-- name: GetContinent1 :one
SELECT * FROM continents
WHERE continent_name = $1 LIMIT 1;

-- name: ListContinents :many
SELECT * FROM continents
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateContinent :one
UPDATE continents
SET continent_name = $2,
continent_desc = $3
WHERE id = $1
RETURNING *;

-- name: DeleteContinent :exec
DELETE FROM continents
WHERE id = $1;
