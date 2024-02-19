-- name: CreateCountry :one
INSERT INTO countries (
    country_name,
    country_desc,
    continent_id
) VALUES (
 $1, $2, $3
)
RETURNING *;

-- name: GetCountry0 :one
SELECT * FROM countries
WHERE id = $1 LIMIT 1;

-- name: GetCountry1 :one
SELECT * FROM countries
WHERE country_name = $1 LIMIT 1;

-- name: ListCountries :many
SELECT * FROM countries
WHERE continent_id = $3
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateCountry :one
UPDATE countries
SET country_name = $2,
country_desc = $3,
continent_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteCountry :exec
DELETE FROM countries
WHERE id = $1;
