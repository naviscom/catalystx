-- name: CreateTech :one
INSERT INTO techs (
    tech_name,
    tech_desc
) VALUES (
 $1, $2
)
RETURNING *;

-- name: GetTech0 :one
SELECT * FROM techs
WHERE id = $1 LIMIT 1;

-- name: GetTech1 :one
SELECT * FROM techs
WHERE tech_name = $1 LIMIT 1;

-- name: ListTechs :many
SELECT * FROM techs
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTech :one
UPDATE techs
SET tech_name = $2,
tech_desc = $3
WHERE id = $1
RETURNING *;

-- name: DeleteTech :exec
DELETE FROM techs
WHERE id = $1;
