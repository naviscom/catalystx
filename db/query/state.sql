-- name: CreateState :one
INSERT INTO states (
    state_name,
    state_desc,
    country_id,
    area_id
) VALUES (
 $1, $2, $3, $4
)
RETURNING *;

-- name: GetState0 :one
SELECT * FROM states
WHERE id = $1 LIMIT 1;

-- name: GetState1 :one
SELECT * FROM states
WHERE state_name = $1 LIMIT 1;

-- name: ListStates :many
SELECT * FROM states
WHERE country_id = $3 OR area_id = $4
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateState :one
UPDATE states
SET state_name = $2,
state_desc = $3,
country_id = $4,
area_id = $5
WHERE id = $1
RETURNING *;

-- name: DeleteState :exec
DELETE FROM states
WHERE id = $1;
