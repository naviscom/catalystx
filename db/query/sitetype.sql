-- name: CreateSitetype :one
INSERT INTO sitetypes (
    type_name,
    type_desc
) VALUES (
 $1, $2
)
RETURNING *;

-- name: GetSitetype0 :one
SELECT * FROM sitetypes
WHERE id = $1 LIMIT 1;

-- name: GetSitetype1 :one
SELECT * FROM sitetypes
WHERE type_name = $1 LIMIT 1;

-- name: ListSitetypes :many
SELECT * FROM sitetypes
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateSitetype :one
UPDATE sitetypes
SET type_name = $2,
type_desc = $3
WHERE id = $1
RETURNING *;

-- name: DeleteSitetype :exec
DELETE FROM sitetypes
WHERE id = $1;
