-- name: CreateBlock :one
INSERT INTO blocks (
    block_name,
    block_desc,
    total_population,
    town_id,
    clutter_id
) VALUES (
 $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetBlock0 :one
SELECT * FROM blocks
WHERE id = $1 LIMIT 1;

-- name: GetBlock1 :one
SELECT * FROM blocks
WHERE block_name = $1 LIMIT 1;

-- name: ListBlocks :many
SELECT * FROM blocks
WHERE town_id = $3 OR clutter_id = $4
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateBlock :one
UPDATE blocks
SET block_name = $2,
block_desc = $3,
total_population = $4,
town_id = $5,
clutter_id = $6
WHERE id = $1
RETURNING *;

-- name: DeleteBlock :exec
DELETE FROM blocks
WHERE id = $1;
